package runner

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *ExpectTask) run(device harness.Device) TaskResult {

	console, err := device.Console()
	if err != nil {
		return TaskResult{
			status: Fatal,
			err:    fmt.Errorf("Expect:run(%q) opening console: %w", t.This, err),
		}
	}

	console.SetReadTimeout(time.Second)

	startTime := time.Now()
	expected := t.This
	timeout := float64(t.Timeout)
	p := 0
	received := ""
	buf := make([]byte, 1)

	for p < len(expected) {
		n, err := console.Read(buf)
		if err != nil {
			return TaskResult{
				status: Fatal,
				err:    fmt.Errorf("Expect:run(%q) reading %w", expected, err),
			}
		}
		if n == 0 {
			if time.Since(startTime).Seconds() > timeout {
				return TaskResult{
					status: Fatal,
					err:    fmt.Errorf("Expect:run(%q) timeout", expected),
				}
			}
			continue
		}
		c := buf[0]
		if t.Echo {
			if c != '\x1b' || !t.DebugEscapes {
				os.Stdout.Write(buf)
				// flush stdout
				os.Stdout.Sync()
			} else {
				os.Stdout.Write([]byte("<ESC>"))
			}
			os.Stdout.Sync()
		}
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

	if t.Send != "" {
		console.Write([]byte(sendStringParser(t.Send)))
		if t.Echo {
			color.Set(color.FgYellow)
			fmt.Println("\n\nsent:", t.Send)
			color.Unset()
		}
	}

	fmt.Println("")

	return TaskResult{
		status: Ok,
		err:    nil,
	}
}

func sendStringParser(send string) string {
	switch strings.ToUpper(send) {
	case "<ESC>":
		return "\x1b"
	case "<F1>":
		return "\x1bOP"
	case "<F2>":
		return "\x1bOQ"
	case "<F3>":
		return "\x1bOR"
	case "<F4>":
		return "\x1bOS"
	case "<F5>":
		return "\x1b[15~"
	case "<F6>":
		return "\x1b[17~"
	case "<F7>":
		return "\x1b[18~"
	case "<F8>":
		return "\x1b[19~"
	case "<F9>":
		return "\x1b[20~"
	case "<F10>":
		return "\x1b[21~"
	case "<F11>":
		return "\x1b[23~"
	case "<UP>":
		return "\x1b[A"
	case "<DOWN>":
		return "\x1b[B"
	case "<LEFT>":
		return "\x1b[D"
	case "<RIGHT>":
		return "\x1b[C"
	case "<ENTER>":
		return "\r\n"
	case "<TAB>":
		return "\t"
	case "<BACKSPACE>":
		return "\x7f"
	case "<DELETE>":
		return "\x1b[3~"
	case "<CTRL+E>":
		return "\x05"
	default:
		return send
	}
}
