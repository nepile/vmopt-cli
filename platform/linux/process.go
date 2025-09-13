// go:build linux

package linux

// if your OS is linux, you can uncomment the code below

type LinuxProcessManager struct{}

func NewLinuxProcessManager() *LinuxProcessManager {
	return &LinuxProcessManager{}
}

func (pm *LinuxProcessManager) SetPriority(pid int32, level int) error {
	// if your os is linux, uncomment the code below
	// return unix.Setpriority(unix.PRIO_PROCESS, int(pid), level)

	// if your os is linux, comment the code below
	return nil
}

func (pm *LinuxProcessManager) KillProcess(pid int32) error {
	// if your os is linux, uncomment the code below
	/**
	proc, err := os.FindProcess(int(pid))

	if err != nil {
		return err
	}

	return proc.Kill()
	**/

	// if your os is linux, comment the code below
	return nil
}
