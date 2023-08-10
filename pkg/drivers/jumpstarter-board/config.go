package jumpstarter_board

import (
	"fmt"
	"strings"
	"sync"
	"time"
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
		consoleMode:    true, // let's asume it's in console mode so we will force exit at start
	}
	jp.readConfig()
	return jp
}

func (d *JumpstarterDevice) IsBusy() (bool, error) {
	return d.busy, nil
}

func (d *JumpstarterDevice) readConfig() error {
	if err := d.ensureSerial(); err != nil {
		d.name = "**BUSY**"
		d.busy = true
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
	// keep reading until we reach the prompt or the read times out
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
			d.name = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(line, "name:")))
		}
		if strings.HasPrefix(line, "tags:") {
			d.tags = strings.Split(strings.ToLower(strings.TrimSpace(strings.TrimPrefix(line, "tags:"))), ",")
		}
		if strings.HasPrefix(line, "storage:") {
			d.storage = strings.ToLower(strings.TrimSpace(strings.TrimPrefix(line, "storage:")))
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

func (d *JumpstarterDevice) Name() string {
	return d.name
}

func (d *JumpstarterDevice) SetTags(tags []string) error {
	joinTags := strings.Join(tags, ",")
	if err := d.ensureSerial(); err != nil {
		return fmt.Errorf("SetTags(%s): %w", joinTags, err)
	}

	if err := d.exitConsole(); err != nil {
		return fmt.Errorf("SetTags(%s): %w", joinTags, err)
	}

	if err := d.sendAndExpect("set-config tags "+strings.Join(tags, ","), "Set tags to "+joinTags); err != nil {
		return fmt.Errorf("SetTags(%s) %w", joinTags, err)
	}

	return nil
}

func (d *JumpstarterDevice) Tags() []string {
	return d.tags
}
