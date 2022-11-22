# Golangdocker

| !["Go"](_assets/logo-golang.png?raw=true "Go") | !["Docker"](_assets/logo-docker.png?raw=true "Docker") | !["Kubernetes"](_assets/logo-k8s.png?raw=true "Kubernetes") |
|:-------------|:-------------:|:-------------|

## To Do

- [x] Dockerize, multi-stage binary build
- [x] Sysinfo for API data example
- [x] Fiber: API
- [x] Usage in [Kubernetes](https://kubernetes.io/) (See: [blob/master/k8s/K8S-README.md](blob/master/k8s/K8S-README.md))
- [ ] TLS? In single container or via K8s?

## Prerequisites

The following programs will need to be installed:

* [Go](https://go.dev/learn/)
* [Docker](https://www.docker.com/get-started/)

Required for Kubernetes itegration:

* A running [Kubernetes](https://kubernetes.io/) cluster
* A [Docker Hub](https://hub.docker.com/) account
* Apache Benchmark (For Ubuntu, it's part of the Apache2 Utilities package, e.g.: `apt-get install apache2-utils `)

While the idea is to get this up and running quickly, it is not a deep dive into Go, Docker, or K8S. Basic knowledge of these technologies is required.

## Summary

This repository is a work in progress, but I'll do my best to keep the Master branch in a working state. Initially, this project was to create a boilerplate for containerizing Go binaries for use a K8s cluster.

## App

Our build is simple, just a compiled Go binary that runs in a container. This binary collects local resources/stats for display as JSON via these API Endpoints using [Fiber](https://docs.gofiber.io/):

```
http://192.168.0.99:5000/api/v1
http://192.168.0.99:5000/api/v1/mem
http://192.168.0.99:5000/api/v1/cpu
http://192.168.0.99:5000/api/v1/host
http://192.168.0.99:5000/api/v1/net
http://192.168.0.99:5000/api/v1/load
```

This walkthrough is not meant to be groundbreaking by any means, but rather to get something minimal, working, and useful up and running quickly.

For example, we can peek into the container via the API endpoint `api/v1/host` and see the docker assigned `hostname: "b189564db0c5"` and verify that it is one running a single process `procs: 1`:

```
{
hostInfo: {
  hostname: "b189564db0c5",
  uptime: 1238849,
  bootTime: 1667920883,
  procs: 1,
  os: "linux",
  platform: "",
  platformFamily: "",
  platformVersion: "",
  kernelVersion: "5.4.0-110-generic",
  kernelArch: "x86_64",
  virtualizationSystem: "docker",
  virtualizationRole: "guest",
  hostId: "12345678-1234-5678-90ab-cddeefaabbcc"
  }
}
```

## Docker

Using [multi-stage builds](https://docs.docker.com/build/building/multi-stage/#use-multi-stage-builds), we will use a very simple `Dockerfile` to containerize our app.

```
FROM golang:alpine as app
WORKDIR /go/src/app
COPY . .
RUN apk add git
RUN CGO_ENABLED=0 go install -ldflags '-extldflags "-static"' -tags timetzdata

FROM scratch
COPY --from=app /go/bin/golangdocker /golangdocker
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/golangdocker"]
```

## [TO DO] Build

`docker build -t mattwiater/golangdocker .`

## [TO DO] Run

`docker run -it -p 5000:5000 --rm mattwiater/golangdocker`

## [TO DO] Scripts

`bash docker_build.sh`
`bash docker_run.sh`

## [TO DO] Notes

Assumptions:

Since we initialized the project with: 

`go mod init github.com/mattwiater/golangdocker`

And each package is in it's own directory: `sysinfo`, `api`, `common`, etc., in order to use these local packages within the `main` Go package, you must enter each directory and type: `go build`

Then, in `main.go`, you can include them like this:

```
...
"github.com/mattwiater/golangdocker/sysinfo"
"github.com/mattwiater/golangdocker/api"
"github.com/mattwiater/golangdocker/common"
...
```

**Note on local packages:** In order to make use of your local package functions, along with running the `go build` command, ensure that your functions are Capital-cased. Otherwise Go will throw an error saying that your method is undefined. Only functions that begin with a capital letter are exported from packages, otherwise they are considered private.