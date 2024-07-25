package main

import (
	"fmt"
	"os/exec"
	"testing"
)

// Build the main package and return the path to the executable
func buildMain(t *testing.T) string {
	exePath := "./testapp" // Path to the compiled executable

	cmd := exec.Command("go", "build", "-o", exePath, ".")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Failed to build executable: %v", err)
	}

	return exePath
}

// Run the CLI command and return the output and error
func runCommand(t *testing.T, exePath string, args ...string) (string, error) {
	cmd := exec.Command(exePath, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Fatalf("Commnd failed: %v; \nError: %v\nOutput: %s", args[0], err, output)
	}

	fmt.Println(args[0], ": ", string(output))
	
	return string(output), err
}

func TestCLICommands(t *testing.T) {
	// Build the main package
	exePath := buildMain(t)

	// Test a CLI command
	_, err := runCommand(t, exePath, "healthcheck")
	_, err = runCommand(t, exePath, "proxyRouterConfig")
	_, err = runCommand(t, exePath, "proxyRouterFiles")
	_, err = runCommand(t, exePath, "blockchainProviders", "create")
	// _, err = runCommand(t, exePath, "blockchainProviders")
	// _, err = runCommand(t, exePath, "createChatCompletions")
	// _, err = runCommand(t, exePath, "initiateProxySession")
	// _, err = runCommand(t, exePath, "blockchainProvidersBids")
	// _, err = runCommand(t, exePath, "blockchainModels")
	// _, err = runCommand(t, exePath, "openBlockchainSession")
	// _, err = runCommand(t, exePath, "closeBlockchainSession")

	// _, err = runCommand(t, exePath, "createChatCompletions")

	if err != nil {
		t.Fatal(err)
	}
	// Check the output
	// expectedOutput := "expected result"
	// if output != expectedOutput {
	// 	t.Errorf("Unexpected output: got %v, want %v", output, expectedOutput)
	// }

	// Clean up the executable
	err = exec.Command("rm", exePath).Run()
	if err != nil {
		t.Fatalf("Failed to remove executable: %v", err)
	}
}
