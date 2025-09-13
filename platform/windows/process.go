// go:build windows

package windows

import (
	"os"

	"golang.org/x/sys/windows"
)

type WindowsProcessManager struct{}

func NewWindowsProcessManager() *WindowsProcessManager {
	return &WindowsProcessManager{}
}

func (pm *WindowsProcessManager) Setpriority(pid int32, level int) error {
	h, err := windows.OpenProcess(windows.PROCESS_SET_INFORMATION, false, uint32(pid))
	if err != nil {
		return err
	}
	defer windows.CloseHandle(h)

	return windows.SetPriorityClass(h, windows.BELOW_NORMAL_PRIORITY_CLASS)
}

func (pm *WindowsProcessManager) KillProcess(pid int32) error {
	proc, err := os.FindProcess(int(pid))

	if err != nil {
		return err
	}

	return proc.Kill()
}
