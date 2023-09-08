package runner

import (
	"fmt"
	"time"

	"github.com/redhat-et/jumpstarter/pkg/harness"
)

func (t *PauseTask) run(device harness.Device) TaskResult {
	s := t.Seconds
	if s < 1 {
		s = 1
	}
	fmt.Println("Pausing for", s, "seconds...")
	time.Sleep(time.Duration(s) * time.Second)

	return TaskResult{
		status: Changed,
		err:    nil,
	}
}
