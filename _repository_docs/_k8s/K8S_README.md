# K8s

[IN PROGRESS]

## To Do
- [ ] Working, but walkthrough needs more definition, test commands to ensure cluster setup is correct.

## Assumptions

You have built the container on the Control Plane node:

`docker build -t {your-docker-hub-account-username}/golangdocker:v1 .`

Above we are going to use the `:v1` tag so that we can use [K8s Rolling Updates](https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/) when we make changes to the image. If you have built images in the previous sections, you'll likely see multiple versions of your image with different tags:

`docker images`

```
REPOSITORY                TAG       IMAGE ID       CREATED         SIZE
mattwiater/golangdocker   latest    e9b376df3a3f   24 minutes ago  11.1MB
mattwiater/golangdocker   v1        e9b376df3a3f   4 minutes ago   11.1MB
...
```

And pushed it to docker hub, e.g.: `docker push mattwiater/golangdocker:v1`

Docker Hub Note: This step is important for the remaining nodes to download and run the image without having to manually build it locally on each node. K8s can use local images to spawn pods, but that would require a manual build on each node (dowmlaoding the repo, building the image, and chaning the manifest entry for `imagePullPolicy: Always` to `imagePullPolicy: Never`), which we are skipping for this demonstration.

For rolling updates, we would just make the necessary updates to our code, build an image tagged with a new version, e.g.: `:v1.1`, `:v2`, etc., push it to docker hub, and then issue the command:

[NEED TO FIX]: `kubectl set image deployments/k8s-golang-api k8s-golang-api=mattwiater/golangdocker:v2`

[PROBLEM]: The command is not working with namespaced deployments, need to adjust.

The command above tells K8s to update the existing deployment to the newer version and it will take care of bringing down the old pods and spawning new pods with no downtime.

[TO DO: Need example after the initial deployment below is complete]

## Load Balancer

Since we want to make use of multiple container instances in our cluster which are accessible via a single external endpoint, we'll need to setup a load balancer.

The basic traffic path is for our setup is:

Ingress: Our domain maps to an exposed service so that we can reach the Service
Service: The load balancer which will route traffic from a singular endpoint to multiple Pods containers via internal Endpoints
Endpoints: Defines which target Pods to route traffic to: internal IP Address and Port

For this example, we'll use Metal-LB to do the heavy lifting.

### Metal LB

METAL-LB: https://metallb.universe.tf/installation/

Ensure that Strict ARP Mode is enabled in your cluster:

`kubectl edit configmap -n kube-system kube-proxy`

Eddit/Add the `mode` and `strictARP` fields to match below:

```
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
mode: "ipvs"
ipvs:
  strictARP: true
```

Next, set up the Metal-LB infrastructure and resources by applying the Metal-LB native Manifest via:

`kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml`

Your cluster will vary, but my setup has static IP Addresses on my local network:

`kubectl get nodes`

NAME          STATUS   ROLES           AGE    VERSION
mjw-udoo-01   Ready    control-plane   181d   v1.25.3
mjw-udoo-02   Ready    worker          181d   v1.25.3
mjw-udoo-03   Ready    worker          181d   v1.25.3
mjw-udoo-04   Ready    worker          181d   v1.25.3


Configure Metal-LB to add these IP addresses to the IPAddressPool (REF: https://metallb.universe.tf/usage/example/)


```
cat <<EOF | kubectl apply -f -
apiVersion: metallb.io/v1beta1
kind: IPAddressPool
metadata:
  name: udoo
  namespace: metallb-system
spec:
  addresses:
  - 192.168.0.91-192.168.0.94
EOF
```

Then, you can name your IPAddressPool and advertise it to the cluster. In my case, I've hjust named it `udoo`

```
cat <<EOF | kubectl apply -f -
apiVersion: metallb.io/v1beta1
kind: BGPAdvertisement
metadata:
  name: external
  namespace: metallb-system
spec:
  ipAddressPools:
  - udoo
EOF
```

Then, create the namespace and deployment for the App. The following code creates the `k8s-golang-api` namespace for the app to run in and be identified with. It is up to you to choose a name that makes sense, but be sure to adjust the the following YAML snippets to reflect your Namespave name in *all* of the `namespace:` fields.

### Create Namespace

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Namespace
metadata:
  name: k8s-golang-api
EOF
```

### Create Deployment

The following defines how K8s will deploy the Pods on your system. It defines the names, associated Namespaces, number of Replicas, Resource Limits, Ports, etc.

```
cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: k8s-golang-api
  namespace: k8s-golang-api
  labels:
    app: k8s-golang-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: k8s-golang-api
  template:
    metadata:
      labels:
        app: k8s-golang-api
    spec:
      containers:
        - name: k8s-golang-api
          image: 'mattwiater/golangdocker:latest'
          env:
          - name: K8S_NODE_NAME
            valueFrom:
              fieldRef:
                fieldPath: spec.nodeName
          - name: K8S_NODE_IP
            valueFrom:
              fieldRef:
                fieldPath: status.hostIP
          - name: K8S_POD_NAME
            valueFrom:
              fieldRef:
                fieldPath: metadata.name
          - name: K8S_POD_NAMESPACE
            valueFrom:
              fieldRef:
                fieldPath: metadata.namespace
          - name: K8S_POD_IP
            valueFrom:
              fieldRef:
                fieldPath: status.podIP
          imagePullPolicy: Always
          resources:
            requests:
              memory: "500Mi"
              cpu: "250m"
            limits:
              memory: "500Mi"
              
              cpu: "250m"
          ports:
            - containerPort: 5000
              protocol: TCP
EOF
```

The final two steps, Service and Ingress are responsible for routing external traffic into the cluster.

### Create Service

You can see that the service is accepting incoming traffic on port 80, and routing to the Pods named `k8s-golang-api` that are already running on Port 5000 (defined in the Deployment manifest above: `containerPort: 5000`)

```
cat <<EOF | kubectl apply -f -
apiVersion: v1
kind: Service
metadata:
  name: k8s-golang-api
  namespace: k8s-golang-api
spec:
  type: LoadBalancer
  selector:
    app: k8s-golang-api
  ports:
  - name: web
    port: 80
    targetPort: 5000
EOF
```

### Create Ingress

In my setup, I want the containers to be accessible via Port 80 at the domain `golang.0nezer0.com`. So the Ingress section below defines the domain mapping to the Serice section

```
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: k8s-golang-api-ingress
  namespace: k8s-golang-api
spec:
  defaultBackend:
    service:
      name: k8s-golang-api
      port:
        number: 80
  rules:
  - host: golang.0nezer0.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: k8s-golang-api
            port:
              number: 80
EOF
```

You can verify the setup via:

`kubectl describe ingress k8s-golang-api-ingress -n=k8s-golang-api`

```
Name:             k8s-golang-api-ingress
Labels:           <none>
Namespace:        k8s-golang-api
Address:
Ingress Class:    <none>
Default backend:  k8s-golang-api:80 (10.244.1.74:5000,10.244.2.104:5000,10.244.3.66:5000)
Rules:
  Host                Path  Backends
  ----                ----  --------
  golang.0nezer0.com
                      /   k8s-golang-api:80 (10.244.1.74:5000,10.244.2.104:5000,10.244.3.66:5000)
Annotations:          <none>
Events:               <none>
```

Note that the domain is listed and the Backend are pointing to the Service we created.

Ensure that you have an IP Address allocated for the Load Balancer:

`kubectl get svc -n=k8s-golang-api`

```
NAME             TYPE           CLUSTER-IP      EXTERNAL-IP    PORT(S)        AGE
k8s-golang-api   LoadBalancer   10.105.31.196   192.168.0.91   80:31188/TCP   21s
```

Assuming that your setup is also on your local network, make sure to add add an IP -> Domain mapping in `/etc/hosts` file on the machine you are accessing the cluster from:

```
192.168.0.91 golang.0nezer0.com
```

## Horizontal Pod Autoscaler (HPA)

`kubectl autoscale deployment -n k8s-golang-api  k8s-golang-api --cpu-percent=75 --memory-percent=75 --min=1 --max=3`