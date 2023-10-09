package sysinfo_test

import (
	"testing"

	"github.com/mwiater/golangdocker/sysinfo"
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
	memInfo, err := sysinfo.GetMemInfo(false)
	if err != nil {
		t.Fatalf("GetMemInfo() returned an error: %v", err)
	}

	if memInfo == nil {
		t.Fatal("GetMemInfo() returned a nil *mem.VirtualMemoryStat object")
	}

	if memInfo.Total == 0 {
		t.Error("Expected Total memory to be non-zero")
	}
}

func TestGetCPUInfo(t *testing.T) {
	cpuInfo, err := sysinfo.GetCPUInfo(false)
	if err != nil {
		t.Fatalf("GetCPUInfo() returned an error: %v", err)
	}

	if len(cpuInfo) == 0 {
		t.Fatal("GetCPUInfo() returned an empty []cpu.InfoStat slice")
	}

	if cpuInfo[0].ModelName == "" {
		t.Error("Expected the first CPU's ModelName to be non-empty")
	}
}

func TestGetHostInfo(t *testing.T) {
	hostInfo, err := sysinfo.GetHostInfo(false)
	if err != nil {
		t.Fatalf("GetHostInfo() returned an error: %v", err)
	}

	if hostInfo == nil {
		t.Fatal("GetHostInfo() returned a nil *host.InfoStat object")
	}

	if hostInfo.Hostname == "" {
		t.Error("Expected Hostname to be non-empty")
	}
}

func TestGetNetInfo(t *testing.T) {
	netInfo, err := sysinfo.GetNetInfo(false)
	if err != nil {
		t.Fatalf("GetNetInfo() returned an error: %v", err)
	}

	if len(netInfo) == 0 {
		t.Fatal("GetNetInfo() returned an empty []net.InterfaceStat slice")
	}

	if netInfo[0].Name == "" {
		t.Error("Expected the first network interface's Name to be non-empty")
	}
}

func TestGetLoadInfo(t *testing.T) {
	loadInfo, err := sysinfo.GetLoadInfo(false)
	if err != nil {
		t.Fatalf("GetLoadInfo() returned an error: %v", err)
	}

	if loadInfo == nil {
		t.Fatal("GetLoadInfo() returned a nil *load.AvgStat object")
	}

	if loadInfo.Load1 < 0 {
		t.Error("Expected Load1 (1-minute load average) to be non-negative")
	}
}
