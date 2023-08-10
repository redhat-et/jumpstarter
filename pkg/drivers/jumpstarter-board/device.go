package jumpstarter_board

import (
	"fmt"
	"strings"
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
	name           string
	storage        string
	tags           []string
	busy           bool
	consoleMode    bool
}

func (d *JumpstarterDevice) Lock() error {
	return d.ensureSerial()
}

func (d *JumpstarterDevice) Unlock() error {
	return d.closeSerial()
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

type JumpstarterConsoleWrapper struct {
	serialPort        serial.Port
	jumpstarterDevice *JumpstarterDevice
}

func (c *JumpstarterDevice) getConsoleWrapper() harness.ConsoleInterface {
	return &JumpstarterConsoleWrapper{
		serialPort:        c.serialPort,
		jumpstarterDevice: c,
	}
}

func (c *JumpstarterConsoleWrapper) Write(p []byte) (n int, err error) {
	if c.serialPort == nil {
		return 0, fmt.Errorf("JumpstarterConsoleWrapper: console has been closed")
	}
	return c.serialPort.Write(p)
}

func (c *JumpstarterConsoleWrapper) Read(p []byte) (n int, err error) {
	if c.serialPort == nil {
		return 0, fmt.Errorf("JumpstarterConsoleWrapper: console has been closed")
	}
	return c.serialPort.Read(p)
}

func (c *JumpstarterConsoleWrapper) Close() error {
	err := c.jumpstarterDevice.exitConsole()
	c.serialPort = nil
	return err
}

func (c *JumpstarterConsoleWrapper) SetReadTimeout(t time.Duration) error {
	if c.serialPort == nil {
		return fmt.Errorf("JumpstarterConsoleWrapper: console has been closed")
	}
	return c.serialPort.SetReadTimeout(t)
}

func (d *JumpstarterDevice) Console() (harness.ConsoleInterface, error) {
	if err := d.ensureSerial(); err != nil {
		return nil, fmt.Errorf("Console: %w", err)
	}

	if d.consoleMode {
		return d.getConsoleWrapper(), nil
	}

	if err := d.exitConsole(); err != nil {
		return nil, fmt.Errorf("Console: %w", err)
	}

	if err := d.sendAndExpectNoPrompt("console", "Entering console mode, type CTRL+B 5 times to exit\r\n"); err != nil {
		return nil, fmt.Errorf("Console: %w", err)
	}

	d.consoleMode = true

	return d.getConsoleWrapper(), nil
}

func (d *JumpstarterDevice) SetConsoleSpeed(bps int) error {
	return harness.ErrNotImplemented
}

func (d *JumpstarterDevice) Driver() harness.HarnessDriver {
	return d.driver
}

func (d *JumpstarterDevice) Version() (string, error) {
	return d.version, nil
}

func (d *JumpstarterDevice) Serial() (string, error) {
	return d.serialNumber, nil
}

func (d *JumpstarterDevice) SetControl(signal string, value string) error {

	signal = strings.ToLower(signal)
	value = strings.ToLower(value)

	switch value {
	case "low":
		value = "l"
	case "high":
		value = "h"
	case "hiz":
		value = "z"
	}

	if signal == "reset" {
		signal = "r"
	}

	// check if is valid (r, a, b, c, or d)
	if !strings.Contains("rabcd", signal) {
		return fmt.Errorf("SetControl(%q,%q): invalid signal, must be any of reset|r|a|b|c|d", signal, value)
	}

	// check if value is valid (h, l or z)
	if !strings.Contains("hlz", value) {
		return fmt.Errorf("SetControl(%q,%q): invalid value, must be any of h|l|z|high|low|hiz", signal, value)
	}

	if err := d.ensureSerial(); err != nil {
		return fmt.Errorf("SetControl(%q,%q): %w", signal, value, err)
	}

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("SetControl(%q,%q): %w", signal, value, err)
	}

	setCmd := fmt.Sprintf("set %s %s", signal, value)

	if err := d.sendAndExpect(setCmd, setCmd+"\r\nSet "); err != nil {
		return fmt.Errorf("SetControl(%q,%q): %w", signal, value, err)
	}

	return nil
}

func (d *JumpstarterDevice) Device() (string, error) {
	return d.devicePath, nil
}
