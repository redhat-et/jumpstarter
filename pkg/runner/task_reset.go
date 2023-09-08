package runner

import (
	"fmt"
	"time"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *ResetTask) run(device harness.Device) TaskResult {

	fmt.Println("Resetting device...")
	err := device.SetControl("r", "l")
	if err != nil {
		return TaskResult{
			status: Fatal,
			err:    err,
		}
	}

	ms := t.TimeMs
	if ms < 1000 {
		ms = 1000
	}

	time.Sleep(time.Duration(ms) * time.Millisecond)

	err = device.SetControl("r", "z")
	if err != nil {
		return TaskResult{
			status: Fatal,
			err:    err,
		}
	}

	return TaskResult{
		status: Changed,
		err:    nil,
	}
}
