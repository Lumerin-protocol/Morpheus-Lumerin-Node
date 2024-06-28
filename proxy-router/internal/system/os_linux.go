package system

import (
	"context"
	"fmt"
	"syscall"

	"github.com/shirou/gopsutil/v3/process"
)

type LinuxConfigurator struct {
	sysctl SysctlCaller
}

func NewOSConfigurator() *LinuxConfigurator {
	return &LinuxConfigurator{
		sysctl: &sysctl{},
	}
}

func (c *LinuxConfigurator) GetConfig() (*Config, error) {
	cfg := &Config{}
	localPortRange, err := c.sysctl.Get("net.ipv4.ip_local_port_range")
	if err != nil {
		return nil, err
	}
	cfg.LocalPortRange = localPortRange

	tcpMaxSynBacklog, err := c.sysctl.Get("net.ipv4.tcp_max_syn_backlog")
	if err != nil {
		return nil, err
	}
	cfg.TcpMaxSynBacklog = tcpMaxSynBacklog

	somaxconn, err := c.sysctl.Get("net.core.somaxconn")
	if err != nil {
		return nil, err
	}
	cfg.Somaxconn = somaxconn

	netdevMaxBacklog, err := c.sysctl.Get("net.core.netdev_max_backlog")
	if err != nil {
		return nil, err
	}
	cfg.NetdevMaxBacklog = netdevMaxBacklog

	var rlimit syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		return nil, err
	}

	cfg.RlimitSoft = rlimit.Cur
	cfg.RlimitHard = rlimit.Max

	return cfg, nil
}

func (c *LinuxConfigurator) ApplyConfig(cfg *Config) error {
	err := c.sysctl.Set("net.ipv4.ip_local_port_range", cfg.LocalPortRange)
	if err != nil {
		return err
	}
	err = c.sysctl.Set("net.ipv4.tcp_max_syn_backlog", cfg.TcpMaxSynBacklog)
	if err != nil {
		return err
	}
	err = c.sysctl.Set("net.core.somaxconn", cfg.Somaxconn)
	if err != nil {
		return err
	}
	err = c.sysctl.Set("net.core.netdev_max_backlog", cfg.NetdevMaxBacklog)
	if err != nil {
		return err
	}
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{
		Cur: cfg.RlimitSoft,
		Max: cfg.RlimitHard,
	})
	return err
}

func (*LinuxConfigurator) GetFileDescriptors(ctx context.Context, pid int) ([]FD, error) {
	process, err := process.NewProcessWithContext(ctx, int32(pid))
	if err != nil {
		return nil, err
	}
	files, err := process.OpenFilesWithContext(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]FD, len(files))
	for i, f := range files {
		items[i] = FD{
			ID:   fmt.Sprint(f.Fd),
			Path: f.Path,
		}
	}
	return items, nil
}
