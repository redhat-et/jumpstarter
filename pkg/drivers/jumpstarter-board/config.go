package jumpstarter_board

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func newJumpstarter(ttyname string, version string, serial string) JumpstarterDevice {
	jp := JumpstarterDevice{
		devicePath:     "/dev/" + ttyname,
		version:        version,
		serialNumber:   serial,
		mutex:          &sync.Mutex{},
		singletonMutex: &sync.Mutex{},
		busy:           false,
		name:           "",
		storage:        "",
		tags:           []string{},
	}
	jp.readConfig()
	return jp
}

func (d *JumpstarterDevice) readConfig() error {
	if err := d.ensureSerial(); err != nil {
		d.name = "**BUSY**"
		return nil
	}

	defer d.closeSerial()

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("readConfig: %w", err)
	}

	if err := d.sendAndExpectNoPrompt("get-config", "get-config\r\n"); err != nil {
		return fmt.Errorf("readConfig: %w", err)
	}
	d.serialPort.SetReadTimeout(100 * time.Millisecond)

	buf := make([]byte, 1024)
	c := make([]byte, 1)
	p := 0
	for p < len(buf) {
		n, err := d.serialPort.Read(c)
		if err != nil {
			return fmt.Errorf("readConfig: %w", err)
		}
		if n == 0 || c[0] == '#' {
			break
		}
		buf[p] = c[0]
		p += 1
	}

	buf = buf[:p]

	lines := strings.Split(string(buf), "\r\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "name:") {
			d.name = strings.TrimSpace(strings.TrimPrefix(line, "name:"))
		}
		if strings.HasPrefix(line, "tags:") {
			d.tags = strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "tags:")), ",")
		}
		if strings.HasPrefix(line, "storage:") {
			d.storage = strings.TrimSpace(strings.TrimPrefix(line, "storage:"))
		}
	}
	if d.name == "" {
		d.name = "jp-" + d.serialNumber
	}

	return nil
}

func (d *JumpstarterDevice) SetName(name string) error {
	if err := d.ensureSerial(); err != nil {
		return fmt.Errorf("SetName(%v): %w", name, err)
	}

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("SetName(%v): %w", name, err)
	}

	if err := d.sendAndExpect("set-config name "+name, "Set name to "+name); err != nil {
		return fmt.Errorf("SetName(%v) %w", name, err)
	}

	return nil
}

func (d *JumpstarterDevice) Name() (string, error) {
	return d.name, nil
}

func (d *JumpstarterDevice) SetTags(tags []string) error {
	return harness.ErrNotImplemented
}

func (d *JumpstarterDevice) Tags() ([]string, error) {
	return d.tags, nil
}
