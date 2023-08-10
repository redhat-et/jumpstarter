package runner

import (
	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *StorageTask) run(device harness.Device) TaskResult {

	err := device.AttachStorage(t.Attached)
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
