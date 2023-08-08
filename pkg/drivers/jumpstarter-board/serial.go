package jumpstarter_board

import (
	"fmt"
	"time"

	"go.bug.st/serial"
)

const PROMPT = "#> "
const ESCAPE_SEQUENCE = "\x02\x02\x02\x02\x02" // CTRL+B 5 times to exit from console to prompt (just in case)

func (d *JumpstarterDevice) ensureSerial() error {
	d.singletonMutex.Lock()
	defer d.singletonMutex.Unlock()
	if d.serialPort == nil {
		return d.openSerial()
	}
	return nil
}

func (d *JumpstarterDevice) closeSerial() error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if d.serialPort != nil {
		err := d.serialPort.Close()
		d.serialPort = nil
		return err
	}
	return nil
}

func (d *JumpstarterDevice) openSerial() error {

	mode := &serial.Mode{
		BaudRate: 115200,
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()

	if d.serialPort != nil {
		d.serialPort.Close()

	}
	port, err := serial.Open(d.devicePath, mode)
	if err != nil {
		return fmt.Errorf("openSerial: %w", err)
	}

	d.serialPort = port
	return nil
}

func (d *JumpstarterDevice) exitConsole() error {

	// make sure we are not in console mode
	if err := d.sendAndExpect(ESCAPE_SEQUENCE, "\n"); err != nil {
		return fmt.Errorf("exitConsole: %w", err)
	}

	// and disable monitor if enabled, as monitor will print any data received in the middle of the output
	if err := d.sendAndExpect("monitor off", "Monitor disabled"); err != nil {
		return fmt.Errorf("exitConsole: %w", err)
	}
	return nil
}

func (d *JumpstarterDevice) sendAndExpect(cmd, expected string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if err := d.send(cmd); err != nil {
		return fmt.Errorf("sendAndExpect(%q, %q) sending: %w", cmd, expected, err)
	}

	if err := d.expect(expected); err != nil {
		return fmt.Errorf("sendAndExpect(%q, %q) expecting response: %w", cmd, expected, err)
	}

	if err := d.expect(PROMPT); err != nil {
		return fmt.Errorf("sendAndExpect(%q, %q) waiting for prompt: %w", cmd, expected, err)
	}
	return nil
}

func (d *JumpstarterDevice) sendAndExpectNoPrompt(cmd, expected string) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	if err := d.send(cmd); err != nil {
		return fmt.Errorf("sendAndExpect(%q, %q) sending: %w", cmd, expected, err)
	}

	if err := d.expect(expected); err != nil {
		return fmt.Errorf("sendAndExpect(%q, %q) expecting response: %w", cmd, expected, err)
	}

	return nil
}

func (d *JumpstarterDevice) send(cmd string) error {
	_, err := d.serialPort.Write([]byte(cmd + "\r\n"))
	if err != nil {
		return fmt.Errorf("sendCommand: %w", err)
	}

	return nil
}
func (d *JumpstarterDevice) expect(expected string) error {
	d.serialPort.SetReadTimeout(1 * time.Second)
	p := 0
	received := ""
	buf := make([]byte, 1)
	for p < len(expected) {
		n, err := d.serialPort.Read(buf)
		if err != nil {
			return fmt.Errorf("expect(%q): %w", expected, err)
		}
		if n == 0 {
			return fmt.Errorf("expect(%q): timeout, received=%q", expected, received)
		}
		c := buf[0]
		received += string(c)
		if c == expected[p] {
			p++
		} else {
			if c != expected[0] {
				p = 0
			} else {
				p = 1
			}
		}
	}
	return nil
}
