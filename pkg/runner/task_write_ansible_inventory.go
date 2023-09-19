package runner

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/redhat-et/jumpstarter/pkg/harness"
	"github.com/redhat-et/jumpstarter/pkg/tools"
)

func (t *WriteAnsibleInventoryTask) run(device harness.Device) TaskResult {

	// Open the inventory file t.Filename for writing
	// If the file doesn't exist, create it or append to the file
	inventoryFile, err := os.OpenFile(t.Filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return TaskResult{
			status: Fatal,
			err:    fmt.Errorf("WriteAnsibleInventory:run(%q) opening file for write: %w", t.Filename, err),
		}
	}
	defer inventoryFile.Close()
	if t.User == "" {
		t.User = "root"
	}

	if err := tools.CreateAnsibleInventory(device, inventoryFile, t.User, t.SshKey); err != nil {
		return TaskResult{
			status: Fatal,
			err:    fmt.Errorf("WriteAnsibleInventory:run(%q) writing inventory file: %w", t.Filename, err),
		}
	}

	color.Set(color.FgYellow)
	fmt.Println("\n\rwritten :", t.Filename)
	color.Unset()

	fmt.Println("")

	return TaskResult{
		status: Ok,
		err:    nil,
	}
}
