package harness

import (
	"fmt"
	"strings"
)

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

// FindDevices iterates over the available drivers and gets a list of devices.
// If a driver is specified, only devices for that driver are returned.

func FindDevices(driverName string, tag string) ([]Device, error) {
	var devices []Device
	for _, driver := range drivers {
		if driverName != "" && driverName != driver.Name() {
			continue // skip this driver
		}

		d, err := driver.FindDevices()
		if err != nil {
			return nil, fmt.Errorf("(%q).FindDevices: %w", driver.Name(), err)
		}
		for _, device := range d {
			if tag != "" {
				deviceTags, err := device.Tags()
				if err != nil {
					return nil, fmt.Errorf("(%q).Tags: %w", driver.Name(), err)
				}
				fmt.Println(deviceTags, tag)
				if contains_tag(deviceTags, tag) {
					devices = append(devices, device)
				}
			} else {
				devices = append(devices, device)
			}
		}

	}
	return devices, nil
}

func contains_tag(slice []string, str string) bool {
	for _, s := range slice {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}

// FindDevice iterates over the available drivers and return a specific Device.
// If a driver is specified, only devices for that driver are returned.
func FindDevice(driverName string, deviceId string) (Device, error) {
	devices, err := FindDevices(driverName, "")
	if err != nil {
		return nil, fmt.Errorf("FindDevices: %w", err)
	}

	for _, device := range devices {
		name, err := device.Name()
		if err != nil {
			return nil, fmt.Errorf("FindDevice (%q).Name: %w", device.Driver().Name(), err)
		}

		serialNumber, err := device.Serial()
		if err != nil {
			return nil, fmt.Errorf("FindDevice (%q).SerialNumber: %w", device.Driver().Name(), err)
		}
		if name == deviceId || serialNumber == deviceId {
			return device, nil
		}
	}
	return nil, fmt.Errorf("FindDevice: %q not found", deviceId)
}
