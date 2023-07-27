package harness

import (
	"errors"
	"io"
)

type Device interface {
	Driver() HarnessDriver
	PowerOn() error
	PowerOff() error
	Console() (io.ReadWriteCloser, error)
	SetConsoleSpeed(bps int) error
	Version() (string, error)
	Name() (string, error)     // name of the device, can be assigned by the user
	SetName(name string) error // set the name of the device, should be stored in config or flashed to device
	Serial() (string, error)
	SetDiskImage(path string) error
	SetControl(signal string, value string) error
}

// basic errors
var ErrCannotOpenDevice = errors.New("unable to open device tty")
var ErrNotImplemented = errors.New("not implemented")
var ErrCannotSetBaudRate = errors.New("unable to set baud rate")
var ErrCannotSetDiskImage = errors.New("unable to set disk image")
var ErrDeviceNotResponding = errors.New("device not responding")
