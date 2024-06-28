package system

import (
	"fmt"
	"os/exec"
	"strings"
)

type SysctlCaller interface {
	Get(name string) (string, error)
	Set(name string, value string) error
}

type sysctl struct{}

func (s *sysctl) Set(name string, value string) error {
	_, err := run("sysctl", "-w", name+"="+value)
	return err
}

func (s *sysctl) Get(name string) (string, error) {
	return run("sysctl", "-n", name)
}

// run executes a command and returns its output. If the command fails, the
// error will contain full output of the command (stdout + stderr)
func run(name string, arg ...string) (out string, err error) {
	outBytes, err := exec.Command(name, arg...).CombinedOutput()
	output := strings.TrimSpace(string(outBytes))
	if err != nil {
		return "", fmt.Errorf("%s: %s", err, output)
	}
	return output, nil
}
