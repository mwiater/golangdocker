package api_test

import (
	"embed"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/mattwiater/golangdocker/api"
	"github.com/mattwiater/golangdocker/config"
	"github.com/stretchr/testify/assert"
)

//go:generate bash copy_env_file.sh

//go:embed .env
var fs embed.FS

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
		description         string
		route               string
		expectedCode        int
		expectedContentType string
		expectedBody        string
	}{
		{
			description:         "false root route",
			route:               "/",
			expectedCode:        302,
			expectedContentType: "application/json",
			expectedBody:        "",
		},
		{
			description:         "routes route",
			route:               "/api/v1",
			expectedCode:        200,
			expectedContentType: "application/json",
			expectedBody:        "{\"apiRoutes\":[\"/\",\"/api/v1\",\"/api/v1/docs/*\",\"/api/v1/metrics\",\"/api/v1/resource/\",\"/api/v1/resource/all\",\"/api/v1/resource/cpu\",\"/api/v1/resource/host\",\"/api/v1/resource/load\",\"/api/v1/resource/memory\",\"/api/v1/resource/network\"]}",
		},
		{
			description:         "cpu route",
			route:               "/api/v1/resource/cpu",
			expectedCode:        200,
			expectedContentType: "application/json",
			expectedBody:        "",
		},
		{
			description:         "host route",
			route:               "/api/v1/resource/host",
			expectedCode:        200,
			expectedContentType: "application/json",
			expectedBody:        "",
		},
		{
			description:         "load route",
			route:               "/api/v1/resource/load",
			expectedCode:        200,
			expectedContentType: "application/json",
			expectedBody:        "",
		},
		{
			description:         "mem route",
			route:               "/api/v1/resource/memory",
			expectedCode:        200,
			expectedContentType: "application/json",
			expectedBody:        "",
		},
		{
			description:         "net route",
			route:               "/api/v1/resource/network",
			expectedCode:        200,
			expectedContentType: "application/json",
			expectedBody:        "",
		},
		{
			description:         "metrics route",
			route:               "/api/v1/metrics",
			expectedCode:        200,
			expectedContentType: "text/html; charset=utf-8",
			expectedBody:        "",
		},
		{
			description:         "docs route",
			route:               "/api/v1/docs/index.html",
			expectedCode:        200,
			expectedContentType: "text/html; charset=utf-8",
			expectedBody:        "",
		},
		{
			description:         "forced 404 route",
			route:               "/api/v1/404",
			expectedCode:        404,
			expectedContentType: "text/html; charset=utf-8",
			expectedBody:        "",
		},
	}

	cfg, err := config.AppConfig(fs)
	if err != nil {
		log.Fatal("Error: config.AppConfig()")
	}

	app := api.SetupRoute(cfg)

	go app.Listen(":" + cfg["SERVERPORT"]) //nolint

	waitForServer(cfg["SERVERPORT"])

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
			fmt.Println(test)
			t.Errorf("Status Code: %#v != http.StatusOK: %#v\n", res.StatusCode, http.StatusOK)
		}

		contentType := res.Header.Get("Content-Type")
		if contentType != test.expectedContentType && res.StatusCode == 200 {
			t.Errorf("Content-Type: %#v != %#v", contentType, test.expectedContentType)
		}

		if test.expectedBody != "" {
			body, err := io.ReadAll(res.Body)
			assert.Nilf(t, err, test.description)
			assert.Equalf(t, test.expectedBody, string(body), test.description)
		}
	}
}
