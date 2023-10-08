package sysinfo_test

import (
	"testing"

	"github.com/mwiater/golangdocker/sysinfo"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/stretchr/testify/assert"
)

func ExampleTestTZ() {
	sysinfo.TestTZ()
	// Output:
	// [ ✓ SUCCESS ] Successfully loaded: America/Los_Angeles
}

func ExampleTestTLS() {
	sysinfo.TestTLS()
	// Output:
	// [ ✓ SUCCESS ] Successfully established https connection to: https://google.com
}

func TestGetMemInfo(t *testing.T) {
	memStat, _ := sysinfo.GetMemInfo(false)
	assert.IsType(t, &mem.VirtualMemoryStat{}, memStat)
}

func TestGetCPUInfo(t *testing.T) {
	cpuStat, _ := sysinfo.GetCPUInfo(false)
	assert.IsType(t, []cpu.InfoStat{}, cpuStat)
}

func TestGetHostInfo(t *testing.T) {
	hostStat, _ := sysinfo.GetHostInfo(false)
	assert.IsType(t, &host.InfoStat{}, hostStat)
}

func TestGetNetInfo(t *testing.T) {
	netStat, _ := sysinfo.GetNetInfo(false)
	assert.IsType(t, []net.InterfaceStat{}, netStat)
}

func TestGetLoadInfo(t *testing.T) {
	loadStat, _ := sysinfo.GetLoadInfo(false)
	assert.IsType(t, &load.AvgStat{}, loadStat)
}
