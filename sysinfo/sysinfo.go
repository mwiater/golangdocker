// Package sysinfo implements functions for gathering system information
// via the https://github.com/shirou/gopsutil/ package.
package sysinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mattwiater/golangdocker/common"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// TestTZ checks to see if the timezone data was properly included in the docker multi-stage build.
//
// This is only executed when 'debug' is set to 'true' in 'config/appConfig.yml'.
func TestTZ() (errors int) {
	pacific, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Printf("     %s Unable to load timezones: %s\n", common.ConsoleFailure("[ ✗ FAILURE ]"), common.ConsoleBold(err))
		errors++
	} else {
		fmt.Printf("     %s Successfully loaded: %s\n", common.ConsoleSuccess("[ ✓ SUCCESS ]"), common.ConsoleBold(pacific))
	}
	return
}

// TestTLS checks to see if the certificates were properly included in the docker multi-stage build.
//
// This is only executed when 'debug' is set to 'true' in 'config/appConfig.yml'.
func TestTLS() (errors int) {
	url := "https://google.com"
	rsp, err := http.Get(url)
	if err != nil {
		fmt.Printf("     %s Unable to establish https connection to: %s\n", common.ConsoleFailure("[ ✗ FAILURE ]"), common.ConsoleBold(err))
		errors++
	} else {
		rsp.Body.Close()
		fmt.Printf("     %s Successfully established https connection to: %s\n", common.ConsoleSuccess("[ ✓ SUCCESS ]"), common.ConsoleBold(url))
	}
	return
}

// GetAPIRoutes collects active Fiber routes, parses them, and returns the data or an error.
//
// Returned data example:
//
//	[
//		"/",
//		"/api/v1",
//		"/api/v1/cpu",
//		"/api/v1/docs/*",
//		"/api/v1/host",
//		"/api/v1/load",
//		"/api/v1/mem",
//		"/api/v1/metrics",
//		"/api/v1/net"
//	]
func GetAPIRoutes(c *fiber.Ctx) []string {
	app := c.App()
	routes := app.GetRoutes()
	routePaths := []string{}
	for _, route := range routes {
		routePaths = append(routePaths, route.Path)
	}
	routePaths = common.UniqueSlice(routePaths)
	sort.Sort(sort.StringSlice(routePaths))
	routePathsJSON, _ := json.Marshal(routePaths)

	if c.Locals("debug") == true {
		fmt.Printf("\n\n%s API Routes:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(routePathsJSON)
	}

	return routePaths
}

// GetMemInfo collects local system memory info, parses it, and returns the data or an error.
//
// Returned data example (truncated ...):
//
//	{
//		memInfo: {
//			total: 8247103488,
//			available: 5860028416,
//			used: 2075901952,
//			usedPercent: 25.171285349099282,
//			...
//		}
//	}
func GetMemInfo(debug bool) (*mem.VirtualMemoryStat, error) {
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return nil, fmt.Errorf("[mem.VirtualMemory() Error] %v", err.Error())
	}
	memInfoBytes, err := json.Marshal(memInfo)
	if err != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", err.Error())
	}
	if debug {
		fmt.Printf("\n\n%s Memory Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(memInfoBytes)
	}

	return memInfo, nil
}

// GetCPUInfo collects local system cpu info, parses it, and returns the data or an error.
//
// Returned data example (truncated ...):
//
//	{
//		memInfo: {
//			total: 8247103488,
//			available: 5860028416,
//			used: 2075901952,
//			usedPercent: 25.171285349099282,
//			...
//		}
//	}

//	{
//		cpuInfo: {
//			[
//				{
//					"cpu": 0,
//					"vendorId": "GenuineIntel",
//					"family": "6",
//					"model": "76",
//					"stepping": 4,
//					"physicalId": "0",
//					 "coreId": "0",
//					"cores": 1,
//					"modelName": "Intel(R) Pentium(R) CPU  N3710  @ 1.60GHz",
//					"mhz": 2560,
//					"cacheSize": 1024,
//					"flags": [
//						"fpu",
//						"vme",
//						"de",
//						...
//					],
//				}
//			]
//		}
//	}
func GetCPUInfo(debug bool) ([]cpu.InfoStat, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, fmt.Errorf("[mem.VirtualMemory() Error] %v", err.Error())
	}
	cpuInfoBytes, err := json.Marshal(cpuInfo)
	if err != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", err.Error())
	}

	if debug {
		fmt.Printf("\n\n%s Memory Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(cpuInfoBytes)
	}

	return cpuInfo, nil
}

// GetHostInfo collects local system host info, parses it, and returns the data or an error.
//
// Returned data example:
//
//	{
//		hostInfo: {
//			uptime: 1386790,
//			bootTime: 1669484114,
//			procs: 193,
//			os: "linux",
//			platform: "ubuntu",
//			platformFamily: "debian",
//			platformVersion: "20.04",
//			kernelVersion: "5.4.0-110-generic",
//			kernelArch: "x86_64",
//			virtualizationSystem: "kvm",
//			virtualizationRole: "host",
//			hostId: "3a114467-105a-48a5-9419-32654a9b2076"
//		}
//	}
func GetHostInfo(debug bool) (*host.InfoStat, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("[host.Info() Error] %v", err.Error())
	}
	hostInfoBytes, err := json.Marshal(hostInfo)
	if err != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", err.Error())
	}

	if debug {
		fmt.Printf("\n\n%s Memory Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(hostInfoBytes)
	}

	return hostInfo, nil
}

// GetNetInfo collects local system network info, parses it, and returns the data or an error.
//
// Returned data example (truncated ...):
//
//	{
//		netInfo: [
//			{
//				index: 1,
//				mtu: 65536,
//				name: "lo",
//				hardwareAddr: "",
//				flags: [
//					"up",
//					"loopback"
//				],
//				addrs: [
//					{
//						addr: "127.0.0.1/8"
//					},
//					{
//						addr: "::1/128"
//					}
//				]
//			},
//			...
//		]
//	}
func GetNetInfo(debug bool) ([]net.InterfaceStat, error) {
	netInfo, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("[net.Interfaces() Error] %v", err.Error())
	}
	netInfoBytes, err := json.Marshal(netInfo)
	if err != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", err.Error())
	}

	if debug {
		fmt.Printf("\n\n%s Memory Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(netInfoBytes)
	}

	return netInfo, nil
}

// GetLoadInfo collects local system load info, parses it, and returns the data or an error.
//
// Returned data example:
//
//	{
//		loadInfo: {
//			load1: 0.58,
//			load5: 0.87,
//			load15: 0.9
//		}
//	}
func GetLoadInfo(debug bool) (*load.AvgStat, error) {
	loadInfo, err := load.Avg()
	if err != nil {
		return nil, fmt.Errorf("[load.Avg() Error] %v", err.Error())
	}
	loadInfoBytes, err := json.Marshal(loadInfo)
	if err != nil {
		return nil, fmt.Errorf("[json.Marshal Error] %v", err.Error())
	}

	if debug {
		fmt.Printf("\n\n%s Memory Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(loadInfoBytes)
	}

	return loadInfo, nil
}
