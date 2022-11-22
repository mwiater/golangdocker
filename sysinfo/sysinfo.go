package sysinfo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

var debug = false

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
	files, err := ioutil.ReadDir("./proc")
	if err != nil {
		fmt.Println(err)
		errors++
	}

	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
	}

	return
}

func GetAPIRoutes(app *fiber.App) []string {
	routes := app.GetRoutes()
	routePaths := []string{}
	for _, route := range routes {
		routePaths = append(routePaths, route.Path)
	}
	routePaths = common.UniqueSlice(routePaths)
	sort.Sort(sort.StringSlice(routePaths))
	routePathsJSON, _ := json.Marshal(routePaths)
	if debug {
		fmt.Printf("\n\n%s API Routes:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(routePathsJSON)
	}

	return routePaths
}

func GetMemInfo() *mem.VirtualMemoryStat {
	memInfo, _ := mem.VirtualMemory()
	memInfoBytes, err := json.Marshal(memInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}
	if debug {
		fmt.Printf("\n\n%s Memory Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(memInfoBytes)
	}

	return memInfo
}

func GetCPUInfo() []cpu.InfoStat {
	cpuInfo, _ := cpu.Info()
	cpuInfoBytes, err := json.Marshal(cpuInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}
	if debug {
		fmt.Printf("\n\n%s CPU Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(cpuInfoBytes)
	}

	return cpuInfo
}

func GetHostInfo() *host.InfoStat {
	hostInfo, _ := host.Info()
	hostInfoBytes, err := json.Marshal(hostInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}
	if debug {
		fmt.Printf("\n\n%s Host Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(hostInfoBytes)
	}

	return hostInfo
}

func GetNetInfo() []net.InterfaceStat {
	netInfo, _ := net.Interfaces()
	netInfoBytes, err := json.Marshal(netInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}
	if debug {
		fmt.Printf("\n\n%s Net Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(netInfoBytes)
	}

	return netInfo
}

func GetLoadInfo() *load.AvgStat {
	loadInfo, _ := load.Avg()
	loadInfoBytes, err := json.Marshal(loadInfo)
	if err != nil {
		fmt.Println("ERR", err)
	}
	if debug {
		fmt.Printf("\n\n%s Load Info:\n\n", common.ConsoleInfo("[ ★ INFO ]"))
		common.PrettyPrintJSONToConsole(loadInfoBytes)
	}

	return loadInfo
}
