package monitor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVMProcesses(t *testing.T) {
	procs, err := GetVMProcess()

	assert.Nil(t, err, "GetVMProcesses must running without error")
	assert.NotNil(t, procs, "The result can't be nil")

	if len(procs) > 0 {
		assert.NotZero(t, procs[0].PID)
		assert.NotEmpty(t, procs[0].Name)
	}
}

func BenchmarkGetVMProcesses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := GetVMProcess()
		if err != nil {
			b.Fatalf("Error: %v", err)
		}
	}
}
