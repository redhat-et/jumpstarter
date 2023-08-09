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
    
    - name: "Attach storage"
      storage:
        attached: true
    
    - name: "Power on"
      power:
        action: cycle
    
    - expect: 
        this: "Press ESCAPE for boot options"
        send: "<esc>"
    
    - uefi-go-to:
        option: "Boot Manager"
    - uefi-go-to:
        option: "UEFI {{ StorageName }}" # find ansible variable insertion
      
    - expect: 
        this: "GRUB version"
        send: "<up-arrow>e<down-arrow><down-arrow><CTRL-E> inst.ks=...."

    - expect: 
        this: "Install finished"

    - name: "Detach storage"
      storage:
        connected: false

    - login-and-get-inventory:
        user: "root"
        password: "{{ env.password }}"
        inventory: "inventory.json"

    - ansible_playbook:
        playbook: test-kmods.yaml
        inventory: "inventory.json"
        extra_args: 
      
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

	if len(playbooks[0].Tasks) != 11 {
		t.Errorf("Expected 12 tasks, got %d", len(playbooks[0].Tasks))
	}
}
