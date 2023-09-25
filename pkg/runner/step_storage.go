package runner

import (
	"strings"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *StorageStep) run(device harness.Device) StepResult {

	attach := strings.ToLower(string(*t)) == "attach"
	err := device.AttachStorage(attach)
	if err != nil {
		return StepResult{
			status: Fatal,
			err:    err,
		}
	}

	return StepResult{
		status: Done,
		err:    nil,
	}
}
