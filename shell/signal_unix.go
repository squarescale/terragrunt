// +build !windows

package shell

import (
	"os"
	"syscall"
)

var forwardSignals []os.Signal = []os.Signal{syscall.SIGTERM}
