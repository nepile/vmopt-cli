// go:build linux

package linux

// if your OS is linux, you can uncomment the code below

/** import (
	"os"
	"golang.org/x/sys/unix"
)

type LinuxProcessManager struct{}

func NewLinuxProcessManager() *LinuxProcessManager {
	return &LinuxProcessManager{}
}

func (pm *LinuxProcessManager) SetPriority(pid int32, level int) error {
	return unix.Setpriority(unix.PRIO_PROCESS, int(pid), level)
}

func (pm *LinuxProcessManager) KillProcess(pid int32) error {
	proc, err := os.FindProcess(int(pid))

	if err != nil {
		return err
	}

	return proc.Kill()
}
**/
