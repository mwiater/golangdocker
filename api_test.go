package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"testing"

	"github.com/mattwiater/golangdocker/api"
	"github.com/mattwiater/golangdocker/config"
	"github.com/stretchr/testify/assert"
)

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
			body, err := ioutil.ReadAll(res.Body)
			assert.Nilf(t, err, test.description)
			assert.Equalf(t, test.expectedBody, string(body), test.description)
		}
	}
}
