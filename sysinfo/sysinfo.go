package sysinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func FileSys() (errors int) {
	outputDirRead, _ := os.Open("./proc")
	procFiles, err := outputDirRead.ReadDir(0)

	if err != nil {
		fmt.Println(err)
		errors++
	}

	for _, procFile := range procFiles {
		fmt.Println(procFile.Name(), procFile.IsDir())
	}

	return
}

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

func GetMemInfo(c *fiber.Ctx) *mem.VirtualMemoryStat {
	memInfo, _ := mem.VirtualMemory()
	memInfoBytes, err := json.Marshal(memInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}

	if c.Locals("debug") == true {
		fmt.Printf("\n\n%s Memory Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(memInfoBytes)
	}

	return memInfo
}

func GetCPUInfo(c *fiber.Ctx) []cpu.InfoStat {
	cpuInfo, _ := cpu.Info()
	cpuInfoBytes, err := json.Marshal(cpuInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}

	if c.Locals("debug") == true {
		fmt.Printf("\n\n%s CPU Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(cpuInfoBytes)
	}

	return cpuInfo
}

func GetHostInfo(c *fiber.Ctx) *host.InfoStat {
	hostInfo, _ := host.Info()
	hostInfoBytes, err := json.Marshal(hostInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}

	if c.Locals("debug") == true {
		fmt.Printf("\n\n%s Host Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(hostInfoBytes)
	}

	return hostInfo
}

func GetNetInfo(c *fiber.Ctx) []net.InterfaceStat {
	netInfo, _ := net.Interfaces()
	netInfoBytes, err := json.Marshal(netInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}

	if c.Locals("debug") == true {
		fmt.Printf("\n\n%s Net Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(netInfoBytes)
	}

	return netInfo
}

func GetLoadInfo(c *fiber.Ctx) *load.AvgStat {
	loadInfo, _ := load.Avg()
	loadInfoBytes, err := json.Marshal(loadInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}
	if c.Locals("debug") == true {
		fmt.Printf("\n\n%s Load Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(loadInfoBytes)
	}

	return loadInfo
}
