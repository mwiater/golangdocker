# Golangdocker

| !["Go"](_repository_assets/logo-golang.png?raw=true "Go") | !["Docker"](_repository_assets/logo-docker.png?raw=true "Docker") | !["Fiber"](_repository_assets/logo-golang-fiber.png?raw=true "Fiber") | !["Kubernetes"](_repository_assets/logo-k8s.png?raw=true "Kubernetes") |
|:-------------:|:-------------:|:-------------:|:-------------:|

## Documentation

[Extensive documentation is here.](https://mwiater.github.io/golangdocker/)

<hr>

## Application

**Or dive right into the code!**

```
git clone https://github.com/mwiater/golangdocker
cd golangdocker
go get
```

<hr>

### Golang App

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