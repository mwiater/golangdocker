# Golangdocker

| !["Go"](_repository_docs/_repository_assets/logo-golang.png?raw=true "Go") | !["Docker"](_repository_docs/_repository_assets/logo-docker.png?raw=true "Docker") | !["Fiber"](_repository_docs/_repository_assets/logo-golang-fiber.png?raw=true "Fiber") | !["Kubernetes"](_repository_docs/_repository_assets/logo-k8s.png?raw=true "Kubernetes") |
|:-------------:|:-------------:|:-------------:|:-------------:|

## Summary

This repository is a work in progress, but I'll do my best to keep the Master branch in a working state. Initially, this project was to create a boilerplate for containerizing Go binaries for use a K8s cluster. For now, just origanizing my notes in order to be able to replicate this process from end-to-end. The idea is to keep this narrow and succinct and be able to use this as a simple boilerplate for Go containers.

## Topics

This project is in three parts, each which build on the previous:

##### 1) A simple but functioanl rest API app written in Go. This rest API incorporates: 
  * The Fiber Metrics middleware (API endpoint: `/api/v1/metrics`).
  * Creating and serving API documentation (using `swag init`) based on Swagger specifications: `/api/v1/docs/`).
  * A `YAML` configuration pattern for setting app variables.
  * Basic Go endpoint tests.
  * Building a binary of the app and embedding external files so that it is portable and self contained.
  * File formatting for *.go files using `gofmt`.

##### 2) Using the app in a Docker container, covering:
  * Docker build concepts.
  * Docker run concepts.
  * Docker image versioning.
  * Ways to make use of bash scripts for repetative tasks.

##### 3) Using the Docker container in Kubernetes
  * This section is the most incomplete, but should be in a working state.
  * You should already have a working K8s cluster available for this section.
  * Does not provide much background, assumes some basic knowledge using `kubectl`.
  * This app will be deployed as a load-balanced Service across a Control Plane and 3 Worker nodes.

## Assumptions

* **IP Addresses:** For the most part, disregard the hard-coded IP addresses in here (e.g.: my K8s cluster and VM IPs (192.168.*.*)). You'll have to sub in your own for your particular envionment. Right now, laziness!
* **Container vs. Pod:** I'm noticing a few instances where I'm using both `container` and `pod` to mean the same thing in the K8s section. Until I make them more consistent, assume they are interchangeable. A pod is basiically a container in in K8s context. While a `pod` can technically have multiple containers, for this demonstration, assume a 1:1 relationship.

## To Do
- [ ] Generate Postman collection for reference?
- [ ] Turn these to-dos into issues!
- [ ] K8s: Use version tagging instead of `:latest` to provide an example of rolling updates. (Started: [K8S_README.md](../../blob/master/_repository_docs/_k8s/K8S_README.md))
- [ ] TLS? In single container or via K8s? To update in: [certs/](../../blob/master/certs/)

## Prerequisites

The following programs will need to be installed:

* [Go](https://go.dev/learn/)
* [Docker](https://www.docker.com/get-started/)

Required for Kubernetes itegration:

* A running [Kubernetes](https://kubernetes.io/) cluster
* A [Docker Hub](https://hub.docker.com/) account

Optional:

* Artillery (nodejs): [Load Testing](../../blob/master/_repository_docs/_loadTesting/ARTILLERY.md)

While the idea is to get this up and running quickly, it is not a deep dive into Go, Docker, or K8S. Basic knowledge of these technologies is required.

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

## App

### Config

There is a simple app config pattern using: `./config/appConfig.yml`

```
# config.yml

server:
  port: 5000
options:
  debug: false
```

Keeping this simple for now, just want to have a boilerplate config pattern within the app for future use.

* Port: The Port that the app listens on, deafult: `5000`
* Debug: More console output on API requests, deafult: `false`

For `debug`, this will print out the JSON response to the console, depending on the endpoint requested. For `/api/v1/host`, you get something like this:

```
[ ★ INFO ] Host Info:
{
        "hostname": "mjw-udoo-01",
        "uptime": 11093,
        "bootTime": 1669484114,
        "procs": 176,
        "os": "linux",
        "platform": "ubuntu",
        "platformFamily": "debian",
        "platformVersion": "20.04",
        "kernelVersion": "5.4.0-110-generic",
        "kernelArch": "x86_64",
        "virtualizationSystem": "kvm",
        "virtualizationRole": "host",
        "hostId": "3a114467-105a-48a5-9419-32654a9b2076"
}
```

### Testing/Developing App

while developing/testing the app, you can run it natively (not in a Docker container) via:

`go run main.go`

Or, for convenience and formatting, run: `bash go_run.sh`

Site will be available at: http://192.168.0.91:5000/api/v1 (substitute your own IP address)

This step should be completed first before running via Docker to ensure everything is working properly.

### Building the Docker container

NOTE: The steps will refer to the docker image: `mattwiater/golangdocker`. You should change these steps to match your own image name, e.g.: `{your-docker-hub-account-username}/golangdocker`

The build command uses the local [Dockerfile](../../blob/master/Dockerfile) to build the image. Substitute your own Docker image tag for mine wherever you see it (`mattwiater/golangdocker`), e.g.: `{your-docker-hub-account-username}/golangdocker`

`docker build -t mattwiater/golangdocker .`

Or, for convenience, run: `bash docker_build.sh '{your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}'`

Once you have built your image successfully, check the output of `docker images` #=>

```
REPOSITORY                TAG       IMAGE ID       CREATED          SIZE
mattwiater/golangdocker   latest    053f21052659   10 minutes ago   10.7MB
...
```

You should see your tagged image in the list, similar to the output above.

## Docker Build Notes

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

### Running the Docker container

Start the container in an interactive shell, with the host port `5000` (the machine you're running Docker on) mapping to the container port (the port the app is running on within the Docker container) `5000` for simplicity. The app is Port if configured here: `./config/appConfig.yml`

`docker run -it -p 5000:5000 --rm mattwiater/golangdocker`

Or, for convenience, run: `bash docker_run.sh '{your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}'`.

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

##### API Info:

```
http://{your-host-ip-address}:5000/api/v1
```

##### System Info:

```
http://{your-host-ip-address}:5000/api/v1/mem
http://{your-host-ip-address}:5000/api/v1/cpu
http://{your-host-ip-address}:5000/api/v1/host
http://{your-host-ip-address}:5000/api/v1/net
http://{your-host-ip-address}:5000/api/v1/load
```

##### API Metrics:

For simplicity, the default [Fiber Monitor middleware](https://docs.gofiber.io/api/middleware/monitor) is included and available at:

`http://{your-host-ip-address}:5000/api/v1/metrics`


##### API Endpoint Documentation via Swagger

`go install github.com/swaggo/swag/cmd/swag@latest`

`go get -u github.com/swaggo/fiber-swagger`

When updating documentation, you must run this to regenerate docs data: `swag init` (`swag init` is incorporated into the bash scripts for convenience, e.g.: [docker_run.sh](../../blob/master/docker_run.sh))

Then, when you run the application, docs are avaialble at:

`http://{your-host-ip-address}:5000/api/v1/docs/index.html`

## Tests

See: [Tests](../../blob/master/_repository_docs/_tests/TESTS.md)

## Linting: Code analysis

Basic linting option via `golangci-lint`

See: [Linting](../../blob/master/_repository_docs/_linting/LINTING.md)

## Gosec: Security analysis

High level gosec usage example.

See: [Gosec](../../blob/master/_repository_docs/_gosec/GOSEC.md)

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
