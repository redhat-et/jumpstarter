package runner

import (
	"fmt"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

// yaml parser

type JumpstarterPlaybook struct {
	Name          string            `yaml:"name"`
	Tags          []string          `yaml:"tags"`
	Drivers       []string          `yaml:"drivers"`
	ExpectTimeout int               `yaml:"expect-timeout"`
	Tasks         []JumpstarterTask `yaml:"tasks"`
	Cleanup       []JumpstarterTask `yaml:"cleanup"`
}

type JumpstarterTask struct {
	// name of the task
	Name                 string                    `yaml:"name"`
	SetDiskImage         *SetDiskImageTask         `yaml:"set-disk-image,omitempty"`
	Expect               *ExpectTask               `yaml:"expect,omitempty"`
	Send                 *SendTask                 `yaml:"send,omitempty"`
	Storage              *StorageTask              `yaml:"storage,omitempty"`
	UefiGoTo             *UefiGoToTask             `yaml:"uefi-go-to,omitempty"`
	Power                *PowerTask                `yaml:"power,omitempty"`
	LoginAndGetInventory *LoginAndGetInventoryTask `yaml:"login-and-get-inventory,omitempty"`
	AnsiblePlaybook      *AnsiblePlaybookTask      `yaml:"ansible-playbook,omitempty"`
}

type SetDiskImageTask struct {
	Image         string `yaml:"image"`
	AttachStorage bool   `yaml:"attach_storage"`
}

type ExpectTask struct {
	This    string `yaml:"string"`
	Fatal   string `yaml:"fatal"`
	Send    string `yaml:"send"`
	Timeout uint   `yaml:"timeout"`
	DelayMs uint   `yaml:"delay_ms"`
}

type SendTask struct {
	This    []string `yaml:"string"`
	DelayMs uint     `yaml:"delay_ms"`
}

type StorageTask struct {
	Attached bool `yaml:"attached"`
}

type UefiGoToTask struct {
	Option string `yaml:"option"`
}

type PowerTask struct {
	Action string `yaml:"action"`
}

type LoginAndGetInventoryTask struct {
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Inventory string `yaml:"inventory"`
}

type AnsiblePlaybookTask struct {
	Playbook  string `yaml:"playbook"`
	Inventory string `yaml:"inventory"`
	ExtraArgs string `yaml:"extra-args"`
}

// a type enum with changed, ok, error
type TaskStatus int

const (
	Changed TaskStatus = iota
	Ok
	Fatal
)

type TaskResult struct {
	status TaskStatus
	err    error
}

func (p *JumpstarterTask) getName() string {
	if p.Name != "" {
		return p.Name
	}

	switch {
	case p.SetDiskImage != nil:
		return "set-disk-image"
	case p.Expect != nil:
		return "expect"
	case p.Send != nil:
		return "send"
	case p.Storage != nil:
		return "storage"
	case p.UefiGoTo != nil:
		return "uefi-go-to"
	case p.Power != nil:
		return "power"
	case p.LoginAndGetInventory != nil:
		return "login-and-get-inventory"
	case p.AnsiblePlaybook != nil:
		return "ansible-playbook"
	default:
		return "unknown"
	}
}

func (p *JumpstarterTask) run(device harness.Device) TaskResult {
	printHeader("TASK", p.getName())
	switch {
	case p.SetDiskImage != nil:
		return p.SetDiskImage.run(device)
		/*
			case p.Expect != nil:
				return p.Expect.run(device)
			case p.Send != nil:
				return p.Send.run(device)
			case p.Storage != nil:
				return p.Storage.run(device)
			case p.UefiGoTo != nil:
				return p.UefiGoTo.run(device)
			case p.Power != nil:
				return p.Power.run(device)
			case p.LoginAndGetInventory != nil:
				return p.LoginAndGetInventory.run(device)
			case p.AnsiblePlaybook != nil:
				return p.AnsiblePlaybook.run(device)
		*/
	}
	return TaskResult{
		status: Fatal,
		err:    fmt.Errorf("Invalid task: %s", p.getName()),
	}
}
