package jumpstarter_board

import (
	"io"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

type JumpstarterDevice struct {
	driver       *JumpstarterDriver
	devicePath   string
	version      string
	serialNumber string
}

func (d JumpstarterDevice) PowerOn() error {
	return nil
}

func (d JumpstarterDevice) PowerOff() error {
	return nil
}

func (d JumpstarterDevice) Console() (io.ReadWriteCloser, error) {
	return nil, nil
}

func (d JumpstarterDevice) SetConsoleSpeed(bps int) error {
	return harness.ErrNotImplemented
}

func (d JumpstarterDevice) Driver() harness.HarnessDriver {
	return d.driver
}

func (d JumpstarterDevice) Name() (string, error) {
	return "jp-" + d.serialNumber, nil
}

func (d JumpstarterDevice) Version() (string, error) {
	return d.version, nil
}

func (d JumpstarterDevice) Serial() (string, error) {
	return d.serialNumber, nil
}

func (d JumpstarterDevice) SetName(name string) error {
	return harness.ErrNotImplemented
}

func (d JumpstarterDevice) SetDiskImage(path string) error {
	return harness.ErrNotImplemented
}

func (d JumpstarterDevice) SetControl(signal string, value string) error {
	return harness.ErrNotImplemented
}
