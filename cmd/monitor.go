package cmd

import (
	"fmt"
	"time"
	"vmopt-cli/logger"
	"vmopt-cli/monitor"

	"github.com/spf13/cobra"
)

var interval int

var monitorCmd *cobra.Command = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor CPU & RAM usage of VMware VMs",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			vms, err := monitor.GetVMProcess()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if len(vms) == 0 {
				fmt.Println("No VMs are running.")
			} else {
				fmt.Printf("\n=== Monitoring VM (interval %ds) ===\n", interval)
				fmt.Printf("%-6s %-15s %-10s %-10s\n", "PID", "Name", "CPU(%)", "RAM(%)")
				for _, vm := range vms {
					fmt.Printf("%-6d %-15s %-10.2f %10.2f\n", vm.PID, vm.Name, vm.CPUUsage, vm.MemUsage)
					logger.LogEvent(fmt.Sprintf("Monitoring %s (CPU: %.2f, RAM: %.2f)", vm.Name, vm.CPUUsage, vm.MemUsage))
				}
			}

			time.Sleep(time.Duration(interval) * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().IntVarP(&interval, "interval", "i", 2, "Monitoring interval in seconds")
}
