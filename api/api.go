package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mattwiater/golangdocker/sysinfo"
)

func readAPIIndex(c *fiber.Ctx) error {
	apiRoutes := sysinfo.GetAPIRoutes(c.App())
	c.Status(200).JSON(&fiber.Map{
		"apiRoutes": apiRoutes,
	})
	return nil
}

func readMemInfo(c *fiber.Ctx) error {
	memInfo := sysinfo.GetMemInfo()
	c.Status(200).JSON(&fiber.Map{
		"memInfo": memInfo,
	})
	return nil
}

func readCPUInfo(c *fiber.Ctx) error {
	cpuInfo := sysinfo.GetCPUInfo()
	c.Status(200).JSON(&fiber.Map{
		"cpuInfo": cpuInfo,
	})
	return nil
}

func readHostInfo(c *fiber.Ctx) error {
	hostInfo := sysinfo.GetHostInfo()
	c.Status(200).JSON(&fiber.Map{
		"hostInfo": hostInfo,
	})
	return nil
}

func readNetInfo(c *fiber.Ctx) error {
	netInfo := sysinfo.GetNetInfo()
	c.Status(200).JSON(&fiber.Map{
		"netInfo": netInfo,
	})
	return nil
}

func readLoadInfo(c *fiber.Ctx) error {
	loadInfo := sysinfo.GetLoadInfo()
	c.Status(200).JSON(&fiber.Map{
		"loadInfo": loadInfo,
	})
	return nil
}

func SetupRoute() *fiber.App {
	app := *fiber.New()

	app.Get("/api/v1", readAPIIndex)
	app.Get("/api/v1/mem", readMemInfo)
	app.Get("/api/v1/cpu", readCPUInfo)
	app.Get("/api/v1/host", readHostInfo)
	app.Get("/api/v1/net", readNetInfo)
	app.Get("/api/v1/load", readLoadInfo)

	return &app
}
