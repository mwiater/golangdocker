# Tests

Very simple tests are in: [api_test.go](../../../../blob/master/api_test.go)

Run via: `go test` #=>

```
 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.40.1                   │
 │               http://127.0.0.1:5000               │
 │       (bound on host 0.0.0.0 and port 5000)       │
 │                                                   │
 │ Handlers ............ 20  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............ 473552 │
 └───────────────────────────────────────────────────┘

[2022-11-29T09:24:16] GET:/: 302 (     0s) | Bytes In: 0 Bytes Out: 0
[2022-11-29T09:24:16] GET:/api/v1: 200 (     0s) | Bytes In: 0 Bytes Out: 136
[2022-11-29T09:24:16] GET:/api/v1/cpu: 200 (    2ms) | Bytes In: 0 Bytes Out: 3593
[2022-11-29T09:24:16] GET:/api/v1/host: 200 (    4ms) | Bytes In: 0 Bytes Out: 337
[2022-11-29T09:24:16] GET:/api/v1/load: 200 (     0s) | Bytes In: 0 Bytes Out: 54
[2022-11-29T09:24:16] GET:/api/v1/mem: 200 (    1ms) | Bytes In: 0 Bytes Out: 708
[2022-11-29T09:24:16] GET:/api/v1/net: 200 (    3ms) | Bytes In: 0 Bytes Out: 1559
[2022-11-29T09:24:16] GET:/api/v1/metrics: 200 (     0s) | Bytes In: 0 Bytes Out: 6186
[2022-11-29T09:24:16] GET:/api/v1/docs/index.html: 200 (    1ms) | Bytes In: 0 Bytes Out: 3519
[2022-11-29T09:24:16] GET:/api/v1/404: 404 (     0s) | Bytes In: 0 Bytes Out: 22
PASS
ok      github.com/mattwiater/golangdocker      0.155s
```