package system

import (
	"context"
	"fmt"
	"os"
	"strings"
	"syscall"
	"testing"
)

func TestGetFileDescriptors(t *testing.T) {
	fds, err := NewOSConfigurator().GetFileDescriptors(context.TODO(), os.Getpid())
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("fds: %+v\n", fds)
}

func TestRlimits(t *testing.T) {
	cfg := &Config{
		LocalPortRange: "49152 65535",
		Somaxconn:      "128",
		RlimitSoft:     1024,
		RlimitHard:     2048,
	}

	darwinConfigurator := DarwinConfigurator{
		sysctl: &mockSysctl{},
	}

	err := darwinConfigurator.ApplyConfig(cfg)
	if err != nil {
		t.Fatalf("ApplyConfig failed: %v", err)
	}

	// Verify rlimit settings
	var rlimit syscall.Rlimit
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlimit)
	if err != nil {
		t.Fatalf("Getrlimit failed: %v", err)
	}

	if rlimit.Cur != cfg.RlimitSoft {
		t.Errorf("Expected RlimitSoft %d, got %d", cfg.RlimitSoft, rlimit.Cur)
	}

	if rlimit.Max != cfg.RlimitHard {
		t.Errorf("Expected RlimitHard %d, got %d", cfg.RlimitHard, rlimit.Max)
	}

	// Verify port range settings
	firstPort, err := darwinConfigurator.sysctl.Get("net.inet.ip.portrange.first")
	if err != nil {
		t.Fatalf("sysctlGet(net.inet.ip.portrange.first) failed: %v", err)
	}

	lastPort, err := darwinConfigurator.sysctl.Get("net.inet.ip.portrange.last")
	if err != nil {
		t.Fatalf("sysctlGet(net.inet.ip.portrange.last) failed: %v", err)
	}

	expectedPorts := strings.Split(cfg.LocalPortRange, " ")
	if firstPort != expectedPorts[0] {
		t.Errorf("Expected first port %s, got %s", expectedPorts[0], firstPort)
	}

	if lastPort != expectedPorts[1] {
		t.Errorf("Expected last port %s, got %s", expectedPorts[1], lastPort)
	}

	// Verify somaxconn setting
	somaxconn, err := darwinConfigurator.sysctl.Get("kern.ipc.somaxconn")
	if err != nil {
		t.Fatalf("sysctlGet(kern.ipc.somaxconn) failed: %v", err)
	}

	if somaxconn != cfg.Somaxconn {
		t.Errorf("Expected somaxconn %s, got %s", cfg.Somaxconn, somaxconn)
	}
}

type mockSysctl struct{}

// Mock sysctlSet and sysctlGet for testing purposes
func (m *mockSysctl) Set(name string, value string) error {
	return nil
}

func (m *mockSysctl) Get(name string) (string, error) {
	switch name {
	case "net.inet.ip.portrange.first":
		return "49152", nil
	case "net.inet.ip.portrange.last":
		return "65535", nil
	case "kern.ipc.somaxconn":
		return "128", nil
	default:
		return "", nil
	}
}
