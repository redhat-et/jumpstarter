package runner

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
