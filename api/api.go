package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/mattwiater/golangdocker/config"
	"github.com/mattwiater/golangdocker/sysinfo"
)

func readAPIIndex(c *fiber.Ctx) error {
	apiRoutes := sysinfo.GetAPIRoutes(c)
	c.Status(200).JSON(&fiber.Map{
		"apiRoutes": apiRoutes,
	})
	return nil
}

func readMemInfo(c *fiber.Ctx) error {
	memInfo := sysinfo.GetMemInfo(c)
	c.Status(200).JSON(&fiber.Map{
		"memInfo": memInfo,
	})
	return nil
}

func readCPUInfo(c *fiber.Ctx) error {
	cpuInfo := sysinfo.GetCPUInfo(c)
	c.Status(200).JSON(&fiber.Map{
		"cpuInfo": cpuInfo,
	})
	return nil
}

func readHostInfo(c *fiber.Ctx) error {
	hostInfo := sysinfo.GetHostInfo(c)
	c.Status(200).JSON(&fiber.Map{
		"hostInfo": hostInfo,
	})
	return nil
}

func readNetInfo(c *fiber.Ctx) error {
	netInfo := sysinfo.GetNetInfo(c)
	c.Status(200).JSON(&fiber.Map{
		"netInfo": netInfo,
	})
	return nil
}

func readLoadInfo(c *fiber.Ctx) error {
	loadInfo := sysinfo.GetLoadInfo(c)
	c.Status(200).JSON(&fiber.Map{
		"loadInfo": loadInfo,
	})
	return nil
}

func hello(c *fiber.Ctx) error {
	c.Status(200).JSON(&fiber.Map{
		"hello": "world",
	})
	return nil
}

func SetupRoute(cfg config.Config) *fiber.App {
	app := *fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("port", cfg.Server.Port)
		c.Locals("debug", cfg.Options.Debug)
		return c.Next()
	})

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${method}:${path}: ${status} (${latency}) | Bytes In: ${bytesReceived} Bytes Out: ${bytesSent}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "America/Los_Angeles",
	}))

	app.Get("/", hello)
	app.Get("/api/v1", readAPIIndex)
	app.Get("/api/metrics", monitor.New(monitor.Config{Title: "golangdocker Metrics Page"}))
	app.Get("/api/v1/mem", readMemInfo)
	app.Get("/api/v1/cpu", readCPUInfo)
	app.Get("/api/v1/host", readHostInfo)
	app.Get("/api/v1/net", readNetInfo)
	app.Get("/api/v1/load", readLoadInfo)

	return &app
}
