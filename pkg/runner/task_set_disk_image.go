package runner

import (
	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *SetDiskImageTask) run(device harness.Device) TaskResult {

	err := device.SetDiskImage(t.Image)
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
