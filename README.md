# Golangdocker

| !["Go"](_repository_assets/logo-golang.png?raw=true "Go") | !["Docker"](_repository_assets/logo-docker.png?raw=true "Docker") | !["Fiber"](_repository_assets/logo-golang-fiber.png?raw=true "Fiber") | !["Kubernetes"](_repository_assets/logo-k8s.png?raw=true "Kubernetes") |
|:-------------:|:-------------:|:-------------:|:-------------:|

## Documentation

**Start here:** [Extensive project documentation is here.](https://mwiater.github.io/golangdocker/)

**Bonus:** For information on how this project relates to the <strong>12-factor-app methodology</strong>, see my article: [Creating a compiled Golang binary for use in a minimal Docker container as defined by the 12 Factor App methodology](https://medium.com/dev-genius/creating-a-compiled-golang-binary-for-use-in-a-minimal-docker-container-ae1c5a720aab)
			</div>

<hr>

## Setup/Install

### Requirements

* Go: https://go.dev/learn/
* Docker: https://www.docker.com/get-started/

#### My development environment:

* `more /etc/os-release`: <strong>Ubuntu 20.04.5 LTS</strong>
* `go version`: <strong>go1.21.2 linux/amd64</strong>
* `docker -v`: <strong>Docker version 24.0.6, build ed223bc</strong>

#### Simple

```
git clone git@github.com:mwiater/golangdocker.git
cd golangdocker
go mod tidy
go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/godoc
go install gotest.tools/gotestsum@latest
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.51.2
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
go mod tidy
go install github.com/swaggo/swag/cmd/swag@latest
go install golang.org/x/tools/cmd/godoc
go install gotest.tools/gotestsum@latest
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.51.2
```

When you're finished with the environment, you can deactivate it: `conda deactivate`

Or, remove it completely: `conda env remove -n golangdocker`

<hr>

## Application

For convenience, I've included a makefile. Just type `make` to see all of the targets (comments added for clarity):

```
Targets in this Makefile:

make docker-build          // Build the golang binary in the docker container
make docker-run            // Run the docker container
make docker-run-multi      // Run 2 instances of the docker container on different ports
make golang-build          // Build the golang binary into: ./bin
make golang-build-arm64    // Build an example golang binary in a different architecture into: ./bin
make golang-godoc          // Generate and run GoDoc server, port 6060
make golang-lint           // Lint the application
make golang-run            // Run the application, no compilation
make golang-test           // Run the tests

For details on these commands, see the bash scripts in the 'scripts/' directory.
```

**NOTE:** Many of the bash scripts execute helpers before the main command, e.g.: `swag init`, `gofmt`, etc. There are exit status checks in place so that, for example, if `gofmt` fails prior to the build (usually because of a syntax error), the script will report the error and exit before trying to build the go binary--which would likely fail due to theerror found via `gofmt`. Here is an example of the script pattern:

```
...
echo -e "${CYANBOLD}Building Swagger docs...${RESET}"
swag init
status=$?
if test $status -ne 0
then
	echo -e "${REDBOLD}...Error: 'swag init' command failed:${RESET}"
	echo ""
	exit 1
fi
echo -e "${GREENBOLD}...Complete.${RESET}"
echo ""

echo -e "${CYANBOLD}Formatting *.go files...${RESET}"
for i in *.go **/*.go ; do
	gofmt -w "$i"
	status=$?
	if test $status -ne 0
	then
		echo -e "${REDBOLD}...Error: 'gofmt' command failed!${RESET}"
		echo ""
		exit 1
	fi
	echo "Formatted: $i"
done;
echo -e "${GREENBOLD}...Complete${RESET}"
echo ""
...
```

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