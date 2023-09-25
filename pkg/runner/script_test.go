package runner

import (
	"log"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestScriptParsing(t *testing.T) {
	var script_data = `
  name: "Test Jetson kmods"
  selector:
    - orin

  expect-timeout: 60
  
  steps:
    - comment: "Powering off and writing the image to disk"
    - power: off
  
    - set-disk-image:
        image: "rhel-guest-image.raw"
  
    - storage: attach
  
    - power: on
  
    - expect:
        this: "login: "
        debug_escapes: false
        timeout: 300
    - send:
        this:
          - "root\n"
          - "redhat\n"
  
    - pause: 3
  
    - expect:
        this: "[root@localhost ~]#"
        debug_escapes: false
  
    - send:
        echo: true
        this:
          - "sudo dnf update -y\n"
      
    - expect:
        timeout: 120
        echo: true
        debug_escapes: false
        this: "Complete"
  
    - expect:
        debug_escapes: false
        this: "[root@localhost ~]#"
  
    - send:
        echo: true
        this:
          - "sudo dnf install -y nvidia-jetpack-kmod\n"
  
    - expect:
        timeout: 120
        echo: true
        debug_escapes: false
        this: "Complete"
  
    - send:
        debug_escapes: false
        this:
          - "reboot\n"
  
    - expect:
        this: "login: "
        debug_escapes: false
        timeout: 500 # the kmod boot takes very long because of some issues with the crypto modules from nvidia
  
    - send:
        this:
          - "root\n"
          - "redhat\n"
  
    - send:
        echo: false # we dont want to capture any of the output so expect will catch it later
        this:
          - "\n"
          - "\n"
      
    - expect:
        debug_escapes: false
        echo: true
        this: "[root@localhost ~]#"
  
    - send:
        echo: false # we dont want to capture any of the output so expect will catch it
        this:
          - "lsmod | grep --color=never nv\n"
  
    - expect:
        echo: true
        this: "nvgpu"
  
    - write-ansible-inventory:
        filename: "inventory.yaml"
        ssh_key: ~/.ssh/id_rsa
    
    - local-shell:
        script: |
          ansible -m ping -i inventory.yaml all
  
  cleanup:
    - send:
        debug_escapes: false
        this:
          - "poweroff\n"
    - pause: 5
  
    - power: off
  
    - storage: detach
  
  
`
	// parse yaml file into a JumpstarterScript struct
	script := JumpstarterScript{}

	err := yaml.Unmarshal([]byte(script_data), &script)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	if err != nil {
		t.Errorf("Expected no parsing error, got %v", err)
	}

	if len(script.Steps) != 23 {
		t.Errorf("Expected 12 steps, got %d", len(script.Steps))
	}
}
