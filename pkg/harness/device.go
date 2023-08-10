package harness

import (
	"errors"
	"io"
	"time"
)

type ConsoleInterface interface {
	io.ReadWriteCloser
	SetReadTimeout(t time.Duration) error
}

type Device interface {
	Driver() HarnessDriver
	Power(on bool) error
	Console() (ConsoleInterface, error)
	SetConsoleSpeed(bps int) error
	Version() (string, error)
	Name() string              // name of the device, can be assigned by the user
	Tags() []string            // tags assigned to the device, can be assigned by the user
	SetName(name string) error // set the name of the device, should be stored in config or flashed to device
	SetTags(tags []string) error
	Serial() (string, error)
	SetDiskImage(path string) error
	AttachStorage(connect bool) error
	SetControl(key string, value string) error
	Device() (string, error)
	IsBusy() (bool, error)
	Lock() error   // open/lock the device so other instances cannot use it
	Unlock() error // close the locked device so other instances can use it
}

// basic errors
var ErrCannotOpenDevice = errors.New("unable to open device tty")
var ErrNotImplemented = errors.New("not implemented")
var ErrCannotSetBaudRate = errors.New("unable to set baud rate")
var ErrCannotSetDiskImage = errors.New("unable to set disk image")
var ErrDeviceNotResponding = errors.New("device not responding")
