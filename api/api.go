package api

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/mattwiater/golangdocker/common"
	"github.com/mattwiater/golangdocker/sysinfo"
	"github.com/shirou/gopsutil/v3/host"
	fiberSwagger "github.com/swaggo/fiber-swagger"

	_ "github.com/mattwiater/golangdocker/docs"
)

// apiFalseRoot ... Redirect root of API server to "/api/v1"
// @Summary Redirect root of API server to "/api/v1"
// @Description Redirect root of API server to "/api/v1"
// @Tags API Routes Root
// @Accept */*
// @Produce json
// @Success 200 {object} []string
// @Router / [get]
func apiFalseRoot(c *fiber.Ctx) error {
	err := c.Redirect("/api/v1")
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	return nil
}

// readAPIIndex ... Get all API routes
// @Summary Get all API routes
// @Description Get all API routes
// @Tags API Routes
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1 [get]
func readAPIIndex(c *fiber.Ctx) error {
	apiRoutes, err := sysinfo.GetAPIRoutes(c)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	c.Status(200).JSON(&fiber.Map{"apiRoutes": apiRoutes}) //nolint
	return nil
}

// readMemInfo ... Get system memory info
// @Summary Get system memory info
// @Description Get system memory info
// @Tags System Memory
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/mem [get]
func readMemInfo(c *fiber.Ctx) error {
	debug := false
	if c.Locals("debug") == true {
		debug = true
	}
	memInfo, err := sysinfo.GetMemInfo(debug)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	c.Status(200).JSON(&fiber.Map{"memInfo": memInfo}) //nolint
	return nil
}

// readCPUInfo ... Get system cpu info
// @Summary Get system cpu info
// @Description Get system cpu info
// @Tags System CPU
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/cpu [get]
func readCPUInfo(c *fiber.Ctx) error {
	debug := false
	if c.Locals("debug") == true {
		debug = true
	}
	cpuInfo, err := sysinfo.GetCPUInfo(debug)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	c.Status(200).JSON(&fiber.Map{"cpuInfo": cpuInfo}) //nolint
	return nil
}

// readHostInfo ... Get system host info
// @Summary Get system host info
// @Description Get system host info
// @Tags System Host
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/host [get]
func readHostInfo(c *fiber.Ctx) error {
	debug := false
	if c.Locals("debug") == true {
		debug = true
	}
	hostInfo, err := sysinfo.GetHostInfo(debug)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	c.Status(200).JSON(&fiber.Map{"hostInfo": hostInfo}) //nolint
	return nil
}

// readNetInfo ... Get system network info
// @Summary Get system network info
// @Description Get system network info
// @Tags System Network
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/net [get]
func readNetInfo(c *fiber.Ctx) error {
	debug := false
	if c.Locals("debug") == true {
		debug = true
	}
	netInfo, err := sysinfo.GetNetInfo(debug)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	c.Status(200).JSON(&fiber.Map{"netInfo": netInfo}) //nolint
	return nil
}

// readLoadInfo ... Get system load info
// @Summary Get system load info
// @Description Get system load info
// @Tags System Load
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/load [get]
func readLoadInfo(c *fiber.Ctx) error {
	debug := false
	if c.Locals("debug") == true {
		debug = true
	}
	loadInfo, err := sysinfo.GetLoadInfo(debug)
	if err != nil {
		return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
	}
	c.Status(200).JSON(&fiber.Map{"loadInfo": loadInfo}) //nolint
	return nil
}

// New creates a new middleware handler that wraps all other middleware
func RouteTimerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		if err != nil {
			return fiber.NewError(fiber.StatusServiceUnavailable, err.Error())
		}
		defer func() {
			c.Append("Server-timing", fmt.Sprintf("route;dur=%v", time.Since(start).Milliseconds()))
		}()
		return nil
	}
}

func CustomHeaders(c *fiber.Ctx) error {
	hostInfo, _ := host.Info()
	c.Append("Hostname", fmt.Sprintf("%v", hostInfo.Hostname))
	c.Append("Hostid", fmt.Sprintf("%v", hostInfo.HostID))
	return c.Next()
}

// SetupRoute ... Setup Fiber API routes
// @Summary Setup Fiber API routes
// @Description Setup Fiber API routes
// @Tags Fiber API
func SetupRoute(cfg map[string]string) *fiber.App {
	if cfg["DEBUG"] == "true" {
		fmt.Println(common.ConsoleInfo("Multi-stage image build tests:"))
		sysinfo.TestTZ()
		sysinfo.TestTLS()
		fmt.Println("")
	}

	app := *fiber.New()
	app.Use(RouteTimerHandler())
	app.Use(CustomHeaders)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("port", cfg["SERVERPORT"])
		c.Locals("debug", cfg["DEBUG"])
		return c.Next()
	})
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${method}:${path}: ${status} (${latency}) | Bytes In: ${bytesReceived} Bytes Out: ${bytesSent}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "America/Los_Angeles",
	}))

	app.Get("/", apiFalseRoot)
	app.Get("/api/v1", readAPIIndex)
	app.Get("/api/v1/metrics", monitor.New(monitor.Config{Title: "golangdocker Metrics Page"}))
	app.Get("/api/v1/mem", readMemInfo)
	app.Get("/api/v1/cpu", readCPUInfo)
	app.Get("/api/v1/host", readHostInfo)
	app.Get("/api/v1/net", readNetInfo)
	app.Get("/api/v1/load", readLoadInfo)

	// Routes for Swagger API Docs
	app.Get("/api/v1/docs/*", fiberSwagger.WrapHandler)

	return &app
}
