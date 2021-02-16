package gwl

import (
    "fmt"
    "strings"
    "os/exec"
)

func getCompositor() (string, error) {
	cmd := exec.Command("loginctl", "show-session", "2", "-p", "Type")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("unable to find system compositor")
	}
	return strings.Trim(strings.ReplaceAll(string(output), "Type=", ""), " \n"), nil
}
