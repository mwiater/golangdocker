# Golangdocker

| !["Go"](_assets/logo-golang.png?raw=true "Go") | !["Docker"](_assets/logo-docker.png?raw=true "Docker") | !["Kubernetes"](_assets/logo-k8s.png?raw=true "Kubernetes") |
|:-------------|:-------------:|:-------------|

## Summary

This repository is a work in progress, but I'll do my best to keep the Master branch in a working state. Initially, this project was to create a boilerplate for containerizing Go binaries for use a K8s cluster. For now, just origanizing my notes in order to be able to replicate this process from end-to-end. The idea is to keep this narrow and succinct and be able to use this as a simple boilerplate for Go containers.

## Assumptions

* IP Addresses: For the most part, disregard the hard-coded IP addresses in here (e.g.: my K8s cluster and VM IPs (192.168.*.*)). You'll have to sub in your own for your particular envionment. Right now, laziness!
* I'm noticing a few instances where I'm using both `container` and `pod` to mean the same thing. Until I make them more consistent, assume they are interchangeable. A pod is basiically a container in in K8s context. While a `pod` can technically have multiple containers, for this demonstration, assume a 1:1 relationship.

## To Do

- [x] Dockerize, multi-stage binary build
- [x] [Sysinfo](https://github.com/shirou/gopsutil) collecting system information
- [x] [Fiber](https://docs.gofiber.io/): Simple Rest API for demonstration
- [x] Usage in [Kubernetes](https://kubernetes.io/) (See: [_k8s/K8S_README.md](../../blob/master/_k8s/K8S_README.md))
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

## App

### Build

The build command uses the local [Dockerfile](../../blob/master/Dockerfile) to build the image. Substitute your own Docker image tag for mine wherever you see it (`mattwiater/golangdocker`), e.g.: `{your-docker-hub-account-username/golangdocker}`

`docker build -t mattwiater/golangdocker .`

### Run

Start the container in an interactive shell, with the host port `5000` (the machine you're running Docker on) mapping to the container port (the port the app is running on within the Docker container) `5000` for simplicity. The app is currently hardcoded to listen on port `5000` via: `err := app.Listen(":5000")`

`docker run -it -p 5000:5000 --rm mattwiater/golangdocker`

You should see teh default Fiber message, e.g.:

```
 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.40.0                   │
 │               http://127.0.0.1:5000               │
 │       (bound on host 0.0.0.0 and port 5000)       │
 │                                                   │
 │ Handlers ............ 14  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ................. 1 │
 └───────────────────────────────────────────────────┘
```

On your host machine, you can now access the container via `http://{your-host-ip-address}:5000`

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