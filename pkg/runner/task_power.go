package runner

import (
	"time"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *PowerTask) run(device harness.Device) TaskResult {

	switch t.Action {
	case "on":
		err := device.Power(true)
		if err != nil {
			return TaskResult{
				status: Fatal,
				err:    err,
			}
		}
	case "off":
		err := device.Power(false)
		if err != nil {
			return TaskResult{
				status: Fatal,
				err:    err,
			}
		}
		time.Sleep(2 * time.Second)
	case "cycle":
		err := device.Power(false)
		if err != nil {
			return TaskResult{
				status: Fatal,
				err:    err,
			}
		}
		time.Sleep(2 * time.Second)
		err = device.Power(true)
		if err != nil {
			return TaskResult{
				status: Fatal,
				err:    err,
			}
		}

	}

	return TaskResult{
		status: Changed,
		err:    nil,
	}
}
