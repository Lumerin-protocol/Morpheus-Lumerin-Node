package system

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"syscall"
)

type DarwinConfigurator struct {
	sysctl SysctlCaller
}

func NewOSConfigurator() *DarwinConfigurator {
	return &DarwinConfigurator{
		sysctl: &sysctl{},
	}
}

func (c *DarwinConfigurator) GetConfig() (*Config, error) {
	cfg := &Config{}
	portRangeFirst, err := c.sysctl.Get("net.inet.ip.portrange.first")
	if err != nil {
		return nil, err
	}
	portRangeLast, err := c.sysctl.Get("net.inet.ip.portrange.last")
	if err != nil {
		return nil, err
	}
	localPortRange := portRangeFirst + " " + portRangeLast
	cfg.LocalPortRange = localPortRange

	// net.ipv4.tcp_max_syn_backlog is not available on Darwin

	somaxconn, err := c.sysctl.Get("kern.ipc.somaxconn")
	if err != nil {
		return nil, err
	}
	cfg.Somaxconn = somaxconn

	// net.core.netdev_max_backlog is not available on Darwin

	var rlimit syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		return nil, err
	}

	cfg.RlimitSoft = rlimit.Cur
	cfg.RlimitHard = rlimit.Max

	return cfg, nil
}

func (c *DarwinConfigurator) ApplyConfig(cfg *Config) error {
	rng := strings.Split(cfg.LocalPortRange, " ")
	err := c.sysctl.Set("net.inet.ip.portrange.first", rng[0])
	if err != nil {
		return err
	}
	err = c.sysctl.Set("net.inet.ip.portrange.last", rng[1])
	if err != nil {
		return err
	}

	// net.ipv4.tcp_max_syn_backlog is not available on Darwin

	err = c.sysctl.Set("kern.ipc.somaxconn", cfg.Somaxconn)
	if err != nil {
		return err
	}

	// net.core.netdev_max_backlog is not available on Darwin

	// TODO: ensure these limits are actually applied
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &syscall.Rlimit{
		Cur: cfg.RlimitSoft,
		Max: cfg.RlimitHard,
	})
	return err
}

func (*DarwinConfigurator) GetFileDescriptors(ctx context.Context, pid int) ([]FD, error) {
	cmd := exec.CommandContext(ctx, "lsof", "-Fn", "+p", fmt.Sprint(pid))
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	items := make([]FD, 0)
	reader := bufio.NewReader(bytes.NewReader(out))
	// read header
	line, err := readPrefixedLine('p', reader)
	if err != nil {
		return nil, err
	}
	if line != fmt.Sprint(pid) {
		return nil, fmt.Errorf("unexpected lsof output: %s", line)
	}
	for {
		item := FD{}

		// fd id
		item.ID, err = readPrefixedLine('f', reader)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		// name
		item.Path, err = readPrefixedLine('n', reader)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

func readPrefixedLine(prefix rune, input *bufio.Reader) (string, error) {
	line, _, err := input.ReadLine()
	if err != nil {
		return "", err
	}

	if rune(line[0]) != prefix {
		return "", fmt.Errorf("unexpected lsof output: %s", line)
	}
	return string(line[1:]), nil
}
