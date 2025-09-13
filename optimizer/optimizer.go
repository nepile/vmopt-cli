package optimizer

import (
	"errors"
	"fmt"
	"runtime"
	"vmopt-cli/monitor"
	"vmopt-cli/platform/linux"
	"vmopt-cli/platform/windows"
)

type ProcessManager interface {
	SetPriority(pid int32, level int) error
	KillProcess(pid int32) error
}

func OptimizeVMs(maxCPU, maxRAM float64, action string) error {
	vms, err := monitor.GetVMProcess()
	if err != nil {
		return err
	}

	if len(vms) == 0 {
		return errors.New("no vms are running")
	}

	var pm ProcessManager
	if runtime.GOOS == "windows" {
		pm = windows.NewWindowsProcessManager()
	} else {
		pm = linux.NewLinuxProcessManager()
	}

	for _, vm := range vms {
		if vm.CPUUsage > maxCPU || float64(vm.MemUsage) > maxRAM {
			switch action {
			case "lower-priority":
				err := pm.SetPriority(vm.PID, 10)
				if err != nil {
					fmt.Printf("Fail to decrease VM priority %d: %v\n", vm.PID, err)
				} else {
					fmt.Printf("VM priority is %d decreased.\n", vm.PID)
				}
			case "kill":
				err := pm.KillProcess(vm.PID)
				if err != nil {
					fmt.Printf("Fail to kill the VM %d: %v\n", vm.PID, err)
				} else {
					fmt.Printf("VM %d has been killed.\n", vm.PID)
				}
			}
		}
	}
	return nil
}
