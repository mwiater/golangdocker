# ARTILLERY

[WORK IN PROGRESS]

## Install

`npm install -g artillery@latest`

## Load Tests

In order to benchmark the different run processes, we need to start the app differently before sending a load test. You will alos want to run these test form a different physical machine that where you're running the container from. Keep in mind that these are not real world load tests, as we are mostly testing to targets within the same network.

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

Assumes working K8s cluster...

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