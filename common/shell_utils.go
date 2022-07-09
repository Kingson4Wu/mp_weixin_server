package common

import (
	"log"
	"os/exec"
)

func ExecShellCmd(cmdStr string) string {
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	return string(output)
}
