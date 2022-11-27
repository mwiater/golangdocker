//go:build integration
// +build integration

package main

import (
	// "fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/mattwiater/golangdocker/api"
	"github.com/mattwiater/golangdocker/config"
)

func WaitServer() {
	timeout := 10 * time.Millisecond
	for {
		conn, err := net.DialTimeout("tcp", "192.168.0.91:5000", timeout)
		if err == nil {
			conn.Close()
			return
		}
		time.Sleep(timeout)
	}
}

func TestIntegration(t *testing.T) {
	configData, _ := conf.ReadFile("config/appConfig.yml")

	cfg, err := config.AppConfig(configData)
	if err != nil {
		log.Fatal(err)
	}

	app := api.SetupRoute(cfg)

	go app.Listen(":" + strconv.Itoa(cfg.Server.Port))

	WaitServer()
	out, err := exec.Command("venom", "run", "tests/*.yml").Output()
	if err != nil {
		t.Fatalf("running venom: %s", string(out))
	}
}
