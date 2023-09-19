package runner

import (
	"os"
	"os/exec"
	"strings"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *LocalShell) run(device harness.Device) TaskResult {
	/*
		// create a temporary file in /tmp
		file, err := os.CreateTemp("/tmp", "jumpstarter-*-localshell.sh")
		if err != nil {
			return TaskResult{
				status: Fatal,
				err:    err,
			}
		}
		file.WriteString(t.Script)
		file.Close()

		defer os.Remove(file.Name())

		// run the script
		cmd := "bash " + file.Name()
		// execute in system

	*/

	cmd := exec.Command("bash")
	cmd.Stdin = strings.NewReader("set -x\n" + t.Script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	if err != nil {
		return TaskResult{
			status: Fatal,
			err:    err,
		}
	}

	return TaskResult{
		status: Ok,
		err:    nil,
	}

}
