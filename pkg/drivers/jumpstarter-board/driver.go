package jumpstarter_board

import (
	"github.com/redhat-et/jumpstarter/pkg/harness"
)

type JumpstarterDriver struct{}

func (d *JumpstarterDriver) Name() string {
	return "jumpstarter-board"
}

func (d *JumpstarterDriver) Description() string {
	return `OpenSource HIL USB harness (https://github.com/redhat-et/jumpstarter-board)
	enables the control of Edge and Embedded devices via USB.
	It has the following capabilities: power metering, power cycling, and serial console
	access, and USB storage switching.
	`
}

func (d *JumpstarterDriver) FindDevices() ([]harness.Device, error) {
	return nil, nil
}

func init() {
	harness.RegisterDriver(&JumpstarterDriver{})
}
