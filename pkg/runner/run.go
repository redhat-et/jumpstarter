package runner

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/fatih/color"
	"github.com/redhat-et/jumpstarter/pkg/harness"
	"gopkg.in/yaml.v3"
)

func RunPlaybook(device_id, driver, yaml_file string) error {

	// parse yaml file into a JumpstarterPlaybook struct
	playbooks := []JumpstarterPlaybook{}

	// read yaml file
	if err := readPlaybook(yaml_file, &playbooks); err != nil {
		return fmt.Errorf("RunPlaybook: %w", err)
	}
	// TODO: check if the yaml contents are consistent

	if len(playbooks) != 1 {
		return fmt.Errorf("RunPlaybook: %q should only have one entry", yaml_file)
	}

	// iterate over each playbook entry
	playbook := playbooks[0]

	var device harness.Device

	// TODO implement retry/wait
	//      sometimes devices are busy or can happen fail due to a race condition
	device, err := playbook.getDevice(device_id, device, driver)
	if err != nil {
		return fmt.Errorf("RunPlaybook: %w", err)
	}
	color.Set(color.FgHiYellow)
	fmt.Printf("⚙ Using device %q with tags %v\n", device.Name(), device.Tags())
	color.Unset()

	return nil
}

func (p *JumpstarterPlaybook) getDevice(device_id string, device harness.Device, driver string) (harness.Device, error) {
	if device_id != "" {

		var err error
		device, err = harness.FindDevice(driver, device_id)
		if err != nil {
			return nil, fmt.Errorf("getDevice: %w", err)
		}
	} else {

		devices, err := harness.FindDevices(driver, p.Tags)
		if err != nil {
			return nil, fmt.Errorf("getDevice: %w", err)
		}

		nonBusy := filterOutBusy(devices)

		if len(devices) == 0 {
			return nil, fmt.Errorf("getDevice: no devices found")
		}

		if len(nonBusy) == 0 {

			return nil, fmt.Errorf("getDevice: all devices are busy")
		}

		device = nonBusy[rand.Intn(len(nonBusy))]
		if err := device.Lock(); err != nil {

			return nil, fmt.Errorf("getDevice: tried to open a device: %w", err)
		}
	}
	return device, nil
}

func filterOutBusy(devices []harness.Device) []harness.Device {
	var freeDevices []harness.Device
	for _, device := range devices {
		if busy, _ := device.IsBusy(); !busy {
			freeDevices = append(freeDevices, device)
		}
	}
	return freeDevices
}

func readPlaybook(yaml_file string, playbook *[]JumpstarterPlaybook) error {
	playbook_data, err := os.ReadFile(yaml_file)
	if err != nil {
		return fmt.Errorf("readPlaybook(%q): Error reading yaml file: %w\n", yaml_file, err)
	}

	if err := yaml.Unmarshal([]byte(playbook_data), &playbook); err != nil {
		return fmt.Errorf("readPlaybook(%q): %w", yaml_file, err)
	}
	return nil
}
