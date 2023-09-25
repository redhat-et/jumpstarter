package runner

import (
	"strings"
	"time"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *PowerStep) run(device harness.Device) StepResult {

	switch strings.ToLower(string(*t)) {
	case "on":
		err := device.Power(true)
		if err != nil {
			return StepResult{
				status: Fatal,
				err:    err,
			}
		}
	case "off":
		err := device.Power(false)
		if err != nil {
			return StepResult{
				status: Fatal,
				err:    err,
			}
		}
		time.Sleep(2 * time.Second)
	case "cycle":
		err := device.Power(false)
		if err != nil {
			return StepResult{
				status: Fatal,
				err:    err,
			}
		}
		time.Sleep(2 * time.Second)
		err = device.Power(true)
		if err != nil {
			return StepResult{
				status: Fatal,
				err:    err,
			}
		}

	}

	return StepResult{
		status: Done,
		err:    nil,
	}
}
