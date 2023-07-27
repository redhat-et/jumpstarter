package jumpstarter_board

import (
	"io"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

type JumpstarterDevice struct {
	driver *JumpstarterDriver
}

func (d *JumpstarterDevice) PowerOn() error {
	return nil
}

func (d *JumpstarterDevice) PowerOff() error {
	return nil
}

func (d *JumpstarterDevice) Console() (io.ReadWriteCloser, error) {
	return nil, nil
}

func (d *JumpstarterDevice) SetConsoleSpeed(bps int) error {
	return harness.ErrNotImplemented
}

func (d *JumpstarterDevice) Driver() harness.HarnessDriver {
	return d.driver
}
