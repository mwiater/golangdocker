# Golangdocker

| !["Go"](_repository_assets/logo-golang.png?raw=true "Go") | !["Docker"](_repository_assets/logo-docker.png?raw=true "Docker") | !["Fiber"](_repository_assets/logo-golang-fiber.png?raw=true "Fiber") | !["Kubernetes"](_repository_assets/logo-k8s.png?raw=true "Kubernetes") |
|:-------------:|:-------------:|:-------------:|:-------------:|

## Documentation

**Start here:** [Extensive documentation is here.](https://mwiater.github.io/golangdocker/)

<hr>

## Setup/Install

### Requirements

* Go: https://go.dev/learn/
* Docker: https://www.docker.com/get-started/

#### My development environment:
`more /etc/os-release`: <strong>Ubuntu 20.04.5 LTS</strong>
`go version`: <strong>go1.18.5 linux/amd64</strong>
`docker -v`: <strong>Docker version 20.10.12, build 20.10.12-0ubuntu2~20.04.1</strong>

#### Simple

```
git clone git@github.com:mwiater/golangdocker.git
cd golangdocker
go get
go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/godoc
go install gotest.tools/gotestsum@latest
```

#### Preferred: Anaconda

* Follow setup for your system here: https://docs.anaconda.com/anaconda/install/
* Anaconda distributions: https://www.anaconda.com/products/distributio

Once installed, you'll also need a compiler for your system, e.g. for Ubuntu: `conda install gxx_linux-64`

Create the environment: `conda create -c conda-forge -n golangdocker go`

Verify: `conda info --envs`

```
# conda environments:
#
base                     /home/matt/anaconda3
golangdocker             /home/matt/anaconda3/envs/golangdocker
```

Activate: `conda activate golangdocker`

```
git clone git@github.com:mwiater/golangdocker.git
cd golangdocker
go get
go install github.com/swaggo/swag/cmd/swag@latest
```

When you're finished with the environment, you can deactivate it: `conda deactivate`

Or, remove it completely: `conda env remove -n golangdocker`

<hr>

## Application

### Basic run, without compilation

Command line: `go run .`

or via the Makefile: `make golang-run`

Output will be similar to:

```
 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2.40.1                   │
 │               http://127.0.0.1:5000               │
 │       (bound on host 0.0.0.0 and port 5000)       │
 │                                                   │
 │ Handlers ............ 22  Processes ........... 1 │
 │ Prefork ....... Disabled  PID ............ 783295 │
 └───────────────────────────────────────────────────┘
```

And will be accessible at: `http://{your-host-ip-address}:5000/`

The above base URL will have the following endpoints, where anything under the `resource` path will return current system information, and other endpoints provide meta information, e.g.: list of endpoints, Swagger documentation, and basic API metrics:

```
/
/api/v1
/api/v1/docs/
/api/v1/metrics
/api/v1/resource/
/api/v1/resource/all
/api/v1/resource/cpu
/api/v1/resource/host
/api/v1/resource/load
/api/v1/resource/memory
/api/v1/resource/network
```

<hr>

### Docker

#### Build

Similarly to the above, you can run the application in a Docker container.

Configure your .env file, if desired:

```
SERVERPORT=5000
DOCKERPORT=5000
DEBUG=false
DOCKERIMAGE=mattwiater/golangdocker
```

Note the **DOCKERIMAGE** name, your Docker image will be build with this variable.

First, build via the Makefile: `make docker-build`

Once you have built your image successfully, check the output of the `docker images` command:

```
REPOSITORY                TAG       IMAGE ID       CREATED          SIZE
mattwiater/golangdocker   latest    053f21052659   1 minute ago     26.4MB
...
```

You should see your new image in the list.

#### Run

`make docker-run`

Same output as the App output above as you will be dropped into an interactive Docker shell...