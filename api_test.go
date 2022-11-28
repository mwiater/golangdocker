package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/mattwiater/golangdocker/api"
	"github.com/mattwiater/golangdocker/config"
	"github.com/stretchr/testify/assert"
)

func waitForServer(port string) {
	backoff := 50 * time.Millisecond

	for i := 0; i < 10; i++ {
		conn, err := net.DialTimeout("tcp", ":"+port, 1*time.Second)
		if err != nil {
			time.Sleep(backoff)
			continue
		}
		err = conn.Close()
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	log.Fatalf("Server on port %s not up after 10 attempts", port)
}

func TestAPIRoutes(t *testing.T) {
	tests := []struct {
		description   string
		route         string
		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "routes route",
			route:         "/api/v1",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"apiRoutes\":[\"/\",\"/api/metrics\",\"/api/v1\",\"/api/v1/cpu\",\"/api/v1/host\",\"/api/v1/load\",\"/api/v1/mem\",\"/api/v1/net\"]}",
		},
		{
			description:   "cpu route",
			route:         "/api/v1/cpu",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "host route",
			route:         "/api/v1/host",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "load route",
			route:         "/api/v1/load",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "mem route",
			route:         "/api/v1/mem",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
		{
			description:   "net route",
			route:         "/api/v1/net",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "",
		},
	}

	configData, _ := conf.ReadFile("config/appConfig.yml")

	cfg, err := config.AppConfig(configData)
	if err != nil {
		log.Fatal(err)
	}

	app := api.SetupRoute(cfg)

	go app.Listen(":" + strconv.Itoa(cfg.Server.Port))

	waitForServer(strconv.Itoa(cfg.Server.Port))

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)
		if err != nil {
			t.Errorf("ERROR: %#v", err)
		}

		if res.StatusCode != test.expectedCode {
			t.Errorf("Status Code: %#v != http.StatusOK: %#v\n", res.StatusCode, http.StatusOK)
		}

		contentType := res.Header.Get("Content-Type")
		if contentType != "application/json" {
			t.Errorf("Content-Type: %#v != 'application/json'", contentType)
		}

		if test.expectedBody != "" {
			body, err := io.ReadAll(res.Body)
			assert.Nilf(t, err, test.description)
			assert.Equalf(t, test.expectedBody, string(body), test.description)
		}
	}
}
