package jumpstarter_board

import (
	"fmt"
	"io"
	"sync"

	"github.com/redhat-et/jumpstarter/pkg/harness"
	"go.bug.st/serial"
)

type JumpstarterDevice struct {
	driver         *JumpstarterDriver
	devicePath     string
	version        string
	serialNumber   string
	serialPort     serial.Port
	mutex          *sync.Mutex
	singletonMutex *sync.Mutex
}

func (d JumpstarterDevice) PowerOn() error {

	if err := d.ensureSerial(); err != nil {
		return fmt.Errorf("PowerOn: %w", err)
	}

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("PowerOn: %w", err)
	}

	if err := d.sendAndExpect("power on", "Device powered on"); err != nil {
		return fmt.Errorf("PowerOn: %w", err)
	}
	return nil
}

func (d JumpstarterDevice) PowerOff() error {
	if err := d.ensureSerial(); err != nil {
		return fmt.Errorf("PowerOff: %w", err)
	}

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("PowerOff: %w", err)
	}

	if err := d.sendAndExpect("power off", "Device powered off"); err != nil {
		return fmt.Errorf("PowerOff: %w", err)
	}
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

func (d JumpstarterDevice) Device() (string, error) {
	return d.devicePath, nil
}
