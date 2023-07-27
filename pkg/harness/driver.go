package harness

import "fmt"

type HarnessDriver interface {
	Name() string
	Description() string
	FindDevices() ([]Device, error)
}

// The initial implementation has only one type of driver, the jumpstarter-board driver
// but this could be extended to support other types of drivers for HIL devices.

var drivers = []HarnessDriver{}

func RegisterDriver(driver HarnessDriver) {
	drivers = append(drivers, driver)
}

func GetDrivers() []HarnessDriver {
	return drivers
}

func FindDevices() ([]Device, error) {
	var devices []Device
	for _, driver := range drivers {
		d, err := driver.FindDevices()
		if err != nil {
			return nil, fmt.Errorf("(%q).FindDevices: %w", driver.Name(), err)
		}
		devices = append(devices, d...)
	}
	return devices, nil
}
