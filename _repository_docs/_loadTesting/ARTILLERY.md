# ARTILLERY

[WORK IN PROGRESS]

## To Do
- [ ] Explore custom metrics options in more depth.
- [ ] Generate applicable reports for comparison between bare go app, dockerized app, and k8s replicas.

## Install

`npm install -g artillery@latest`

## Plugins

[Official: Per-endpoint (URL) metrics](https://www.artillery.io/docs/guides/plugins/plugin-metrics-by-endpoint#useonlyrequestnames)

`npm install artillery-plugin-metrics-by-endpoint`

## Custom Scripts

Reference: https://www.artillery.io/docs/guides/guides/extension-apis#example

Working project example here: [custom-artillery-functions.js](../../../../blob/master/_repository_docs/_loadTesting/custom-artillery-functions.js)

This simple example makes use of a custom Fiber middleware wrapper that captures the time spent on the server in each API call and sets a `Server-Timing` response header, e.g.: `Server-Timing: route;dur=16`. See the [RouteTimerHandler()](../../../../blob/master/api/api.go) function in `api/api.go`.

## Load Tests

In order to benchmark the different run processes, we need to start the app differently before sending a load test. You will alos want to run these test form a different physical machine that where you're running the container from. Keep in mind that these are not real world load tests, as we are mostly testing to targets within the same network. These tests are mainly for comparisons of ruinning the app with different mechanisims, e.g: go app, inside Docker container, within K8s w/ replicas.

## No container, bare app

E.g.: `bash go_run.sh`

```
clear && \
artillery run --output golangdocker-bare.json --target http://192.168.0.91:5000/api golangdocker-loadtest.yml && \
	artillery report golangdocker-bare.json
```

## Docker container

E.g.: `bash docker_run.sh`

```
clear && \
artillery run --output golangdocker-docker.json --target http://192.168.0.91:5000/api golangdocker-loadtest.yml && \
	artillery report golangdocker-docker.json
```

## K8s

Assumes working K8s cluster and manual scaling of replicas for each test, e.g.:

```
clear && \
artillery run --output golangdocker-k8s-3-replica.json --target http://golang.0nezer0.com/api golangdocker-loadtest.yml && \
	artillery report golangdocker-k8s-3-replica.json
```

```
clear && \
artillery run --output golangdocker-k8s-2-replica.json --target http://golang.0nezer0.com/api golangdocker-loadtest.yml && \
	artillery report golangdocker-k8s-2-replica.json
```

...