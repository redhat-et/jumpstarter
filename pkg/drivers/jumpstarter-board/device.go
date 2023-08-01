package jumpstarter_board

import (
	"fmt"
	"io"
	"sync"
	"time"

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

func (d *JumpstarterDevice) Power(on bool) error {
	if err := d.ensureSerial(); err != nil {
		return fmt.Errorf("Power(%v): %w", on, err)
	}

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("Power(%v): %w", on, err)
	}
	if on {
		if err := d.sendAndExpect("power on", "Device powered on"); err != nil {
			return fmt.Errorf("PowerOn: %w", err)
		}
	} else {
		if err := d.sendAndExpect("power off", "Device powered off"); err != nil {
			return fmt.Errorf("Power(%v): %w", on, err)
		}
	}
	return nil
}

func (d *JumpstarterDevice) Console() (io.ReadWriteCloser, error) {
	if err := d.ensureSerial(); err != nil {
		return nil, fmt.Errorf("Console: %w", err)
	}

	if err := d.exitConsole(); err != nil {
		return nil, fmt.Errorf("Console: %w", err)
	}

	if err := d.sendAndExpectNoPrompt("console", "Entering console mode, type CTRL+B 5 times to exit\r\n"); err != nil {
		return nil, fmt.Errorf("Console: %w", err)
	}

	return d.serialPort, nil
}

func (d *JumpstarterDevice) SetConsoleSpeed(bps int) error {
	return harness.ErrNotImplemented
}

func (d *JumpstarterDevice) Driver() harness.HarnessDriver {
	return d.driver
}

func (d *JumpstarterDevice) Name() (string, error) {
	return "jp-" + d.serialNumber, nil
}

func (d *JumpstarterDevice) Version() (string, error) {
	return d.version, nil
}

func (d *JumpstarterDevice) Serial() (string, error) {
	return d.serialNumber, nil
}

func (d *JumpstarterDevice) SetName(name string) error {
	return harness.ErrNotImplemented
}

func (d *JumpstarterDevice) SetDiskImage(path string) error {

	fmt.Print("Detecting USB storage device and connecting to host: ")
	diskPath, err := d.detectStorageDevice()
	if err != nil {
		return fmt.Errorf("SetDiskImage: %w", err)
	}
	fmt.Println("done")

	fmt.Printf("%s -> %s: \n", path, diskPath)

	if err := writeImageToDisk(path, diskPath); err != nil {
		return fmt.Errorf("SetDiskImage: %w", err)
	}

	fmt.Print("Connecting storage to device under test.. ")
	if err := d.connectStorageTo(OFF); err != nil {
		return fmt.Errorf("SetDiskImage: %w", err)
	}

	time.Sleep(1 * time.Second) // enough time to power cycle the USB disk
	if err := d.connectStorageTo(DUT); err != nil {
		return fmt.Errorf("SetDiskImage: %w", err)
	}

	fmt.Println("done")

	return nil
}

func (d *JumpstarterDevice) SetControl(signal string, value string) error {
	return harness.ErrNotImplemented
}

func (d *JumpstarterDevice) Device() (string, error) {
	return d.devicePath, nil
}
