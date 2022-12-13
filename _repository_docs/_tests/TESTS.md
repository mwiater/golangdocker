# Tests

Very simple tests are in: [api_test.go](../../../../blob/master/api_test.go)

Run via: `clear && go test -v $(go list ./... | grep -v /docs | grep -v /config | grep -v /api)` #=>


```
=== RUN   TestAPIRoutes

 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.40.1                   │
 │               http://127.0.0.1:5000               │
 │       (bound on host 0.0.0.0 and port 5000)       │
 │                                                   │
 │ Handlers ............ 22  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ........... 2214242 │
 └───────────────────────────────────────────────────┘

[2022-12-13T11:01:06] GET:/: 302 (     0s) | Bytes In: 0 Bytes Out: 0
[2022-12-13T11:01:06] GET:/api/v1: 200 (     0s) | Bytes In: 0 Bytes Out: 136
[2022-12-13T11:01:06] GET:/api/v1/cpu: 200 (    1ms) | Bytes In: 0 Bytes Out: 3593
[2022-12-13T11:01:06] GET:/api/v1/host: 200 (    1ms) | Bytes In: 0 Bytes Out: 338
[2022-12-13T11:01:06] GET:/api/v1/load: 200 (     0s) | Bytes In: 0 Bytes Out: 54
[2022-12-13T11:01:06] GET:/api/v1/mem: 200 (    1ms) | Bytes In: 0 Bytes Out: 706
[2022-12-13T11:01:06] GET:/api/v1/net: 200 (    2ms) | Bytes In: 0 Bytes Out: 1559
[2022-12-13T11:01:06] GET:/api/v1/metrics: 200 (     0s) | Bytes In: 0 Bytes Out: 6186
[2022-12-13T11:01:06] GET:/api/v1/docs/index.html: 200 (     0s) | Bytes In: 0 Bytes Out: 3519
[2022-12-13T11:01:06] GET:/api/v1/404: 404 (     0s) | Bytes In: 0 Bytes Out: 22
--- PASS: TestAPIRoutes (0.13s)
PASS
ok      github.com/mattwiater/golangdocker      0.190s
=== RUN   ExamplePrettyPrintJSONToConsole
--- PASS: ExamplePrettyPrintJSONToConsole (0.00s)
=== RUN   ExampleUniqueSlice
--- PASS: ExampleUniqueSlice (0.00s)
PASS
ok      github.com/mattwiater/golangdocker/common       0.005s
=== RUN   ExampleTestTZ
--- PASS: ExampleTestTZ (0.00s)
=== RUN   ExampleTestTLS
--- PASS: ExampleTestTLS (0.35s)
PASS
ok      github.com/mattwiater/golangdocker/sysinfo      0.365s
```

## Test Cache

To clear the test cache, run: `go clean -testcache`