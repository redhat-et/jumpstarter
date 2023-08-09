package runner

// yaml parser

type JumpstarterPlaybook struct {
	// name of the playbook
	Name string `yaml:"name"`
	// list of tags to match
	Tags []string `yaml:"tags"`
	// list of drivers to match
	Drivers []string `yaml:"drivers"`
	// expect-timeout
	ExpectTimeout int `yaml:"expect-timeout"`
	// list of tasks to execute
	Tasks []JumpstarterTask `yaml:"tasks"`
}

type JumpstarterTask struct {
	// name of the task
	Name                 string                    `yaml:"name"`
	SetDiskImage         *SetDiskImageTask         `yaml:"set-disk-image,omitempty"`
	Expect               *ExpectTask               `yaml:"expect,omitempty"`
	Storage              *StorageTask              `yaml:"storage,omitempty"`
	UefiGoTo             *UefiGoToTask             `yaml:"uefi-go-to,omitempty"`
	Power                *PowerTask                `yaml:"power,omitempty"`
	LoginAndGetInventory *LoginAndGetInventoryTask `yaml:"login-and-get-inventory,omitempty"`
	AnsiblePlaybook      *AnsiblePlaybookTask      `yaml:"ansible-playbook,omitempty"`
}

type SetDiskImageTask struct {
	Image string `yaml:"image"`
}

type ExpectTask struct {
	This string `yaml:"string"`
	Send string `yaml:"send"`
}

type StorageTask struct {
	Attach bool `yaml:"attach"`
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
