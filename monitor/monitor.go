package monitor

import "github.com/shirou/gopsutil/v3/process"

type VMProcess struct {
	PID      int32
	Name     string
	CPUUsage float64
	MemUsage float32
}

func GetVMProcess() ([]VMProcess, error) {
	procs, err := process.Processes()
	if err != nil {
		return nil, err
	}

	var result []VMProcess

	for _, p := range procs {
		name, _ := p.Name()
		if name == "vmware-vmx" || name == "vmware-vmx.exe" {
			cpu, _ := p.CPUPercent()
			mem, _ := p.MemoryPercent()

			result = append(result, VMProcess{
				PID:      p.Pid,
				Name:     name,
				CPUUsage: cpu,
				MemUsage: mem,
			})
		}
	}

	return result, nil
}
