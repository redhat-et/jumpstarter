package runner

import (
	"log"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestPlaybookParsing(t *testing.T) {
	var playbook = `
- name: "Test Jetson kmods"
  tags:
    - nvidia-xavier-nx
    - 8gb
  expect-timeout: 60
  tasks:
    - name: "Load image"
      set-disk-image:
        image: "my-rhel9.3-aarch64-with-kmods.iso"
        attach_storage: true
        offset-gb: 0


    - name: "Power on"
      power:
        action: cycle

    - expect:
        this: "Press ESCAPE for boot options"
        send: "\e"

    - uefi-go-to:
        option: "Boot Manager"
    - uefi-go-to:
        option: "UEFI {{ StorageName }}" # find ansible variable insertion

    - name: "Interact with GRUB to setup special install parameters, just an example"
      expect:
        this: "GRUB version"
        send: "\e"
    - send:
        delay_ms: 100 # delay before and between each send sequence
        this:
          - "<UP>"
          - "e"
          - "<DOWN>"
          - "<DOWN>"
          - "<CTRL-E>"
          - " inst.text console=ttyS0,115200"
          - "<F10>"

    - expect:
        timeout: 1800 # 30 minutes
        fatal: "Unrecoverable error happened" # if this string is found, we will fail
        this: "Install finished"
        send: "\n"

    - name: "Detach storage"
      storage:
        attached: false

    - power:
        action: cycle

    - login-and-get-inventory:    # This will wait for the login sequency,
        user: "root"              # then gather the inventory for network connection
        password: "{{ env.IMAGE_PASSWORD }}"
        inventory: "inventory.json"

    - ansible_playbook:
        playbook: test-kmods.yaml
        inventory: "inventory.json"
        extra_args:
  cleanup:
    - name: "Power off"
      power:
        action: off

    - name: "Detach storage"
      storage:
        attached: false
`
	// parse yaml file into a JumpstarterPlaybook struct
	playbooks := []JumpstarterPlaybook{}

	err := yaml.Unmarshal([]byte(playbook), &playbooks)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if err != nil {
		t.Errorf("Expected no parsing error, got %v", err)
	}

	if len(playbooks) != 1 {
		t.Errorf("Expected 1 playbook, got %d", len(playbooks))
	}

	if len(playbooks[0].Tasks) != 12 {
		t.Errorf("Expected 12 tasks, got %d", len(playbooks[0].Tasks))
	}
}
