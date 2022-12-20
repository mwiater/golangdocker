---
layout: null
---
{% include header.html %}

  <div class="container">
    <header role="banner">
      golangdocker
    </header>
    <div class="row">
      <div class="col">
        <h1 id="k8s">K8s</h1>
        <p>[IN PROGRESS]</p>
        <h2 id="to-do">To Do</h2>
        <ul>
          <li>[ ] Working, but walkthrough needs more definition, test commands to ensure cluster setup is correct.
          </li>
        </ul>
        <h2 id="assumptions">Assumptions</h2>
        <p>You have built the container on the Control Plane node:</p>
        <p><code>docker build -t {your-docker-hub-account-username}/golangdocker:v1 .</code></p>
        <p>Above we are going to use the <code>:v1</code> tag so that we can use <a
            href="https://kubernetes.io/docs/tutorials/kubernetes-basics/update/update-intro/">K8s Rolling Updates</a>
          when we make changes to the image. If you have built images in the previous sections, you&#39;ll likely see
          multiple versions of your image with different tags:</p>
        <p><code>docker images</code></p>
        <pre><code>REPOSITORY                TAG       IMAGE ID       CREATED         SIZE
mattwiater/golangdocker   latest    e9b376df3a3f   <span class="hljs-number">24</span> minutes ago  <span class="hljs-number">11.1</span>MB
mattwiater/golangdocker   v1        e9b376df3a3f   <span class="hljs-number">4</span> minutes ago   <span class="hljs-number">11.1</span>MB
...
</code></pre>
        <p>And pushed it to docker hub, e.g.: <code>docker push mattwiater/golangdocker:v1</code></p>
        <p>Docker Hub Note: This step is important for the remaining nodes to download and run the image without
          having to manually build it locally on each node. K8s can use local images to spawn pods, but that would
          require a manual build on each node (dowmlaoding the repo, building the image, and chaning the manifest
          entry for <code>imagePullPolicy: Always</code> to <code>imagePullPolicy: Never</code>), which we are
          skipping for this demonstration.</p>
        <p>For rolling updates, we would just make the necessary updates to our code, build an image tagged with a new
          version, e.g.: <code>:v1.1</code>, <code>:v2</code>, etc., push it to docker hub, and then issue the
          command:</p>
        <p>[NEED TO FIX]:
          <code>kubectl set image deployments/k8s-golang-api k8s-golang-api=mattwiater/golangdocker:v2</code>
        </p>
        <p>[PROBLEM]: The command is not working with namespaced deployments, need to adjust.</p>
        <p>The command above tells K8s to update the existing deployment to the newer version and it will take care of
          bringing down the old pods and spawning new pods with no downtime.</p>
        <p>[TO DO: Need example after the initial deployment below is complete]</p>
        <h2 id="load-balancer">Load Balancer</h2>
        <p>Since we want to make use of multiple container instances in our cluster which are accessible via a single
          external endpoint, we&#39;ll need to setup a load balancer.</p>
        <p>The basic traffic path is for our setup is:</p>
        <p>Ingress: Our domain maps to an exposed service so that we can reach the Service
          Service: The load balancer which will route traffic from a singular endpoint to multiple Pods containers via
          internal Endpoints
          Endpoints: Defines which target Pods to route traffic to: internal IP Address and Port</p>
        <p>For this example, we&#39;ll use Metal-LB to do the heavy lifting.</p>
        <h3 id="metal-lb">Metal LB</h3>
        <p>METAL-LB: <a href="https://metallb.universe.tf/installation/">https://metallb.universe.tf/installation/</a>
        </p>
        <p>Ensure that Strict ARP Mode is enabled in your cluster:</p>
        <p><code>kubectl edit configmap -n kube-system kube-proxy</code></p>
        <p>Eddit/Add the <code>mode</code> and <code>strictARP</code> fields to match below:</p>
        <pre><code><span class="hljs-symbol">apiVersion:</span> kubeproxy.config.k8s.io/v1alpha1
<span class="hljs-symbol">kind:</span> KubeProxyConfiguration
<span class="hljs-symbol">mode:</span> <span class="hljs-string">"ipvs"</span>
<span class="hljs-symbol">ipvs:</span>
<span class="hljs-symbol">  strictARP:</span> true
</code></pre>
        <p>Next, set up the Metal-LB infrastructure and resources by applying the Metal-LB native Manifest via:</p>
        <p>
          <code>kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml</code>
        </p>
        <p>Your cluster will vary, but my setup has static IP Addresses on my local network:</p>
        <p><code>kubectl get nodes</code></p>
        <p>NAME STATUS ROLES AGE VERSION
          mjw-udoo-01 Ready control-plane 181d v1.25.3
          mjw-udoo-02 Ready worker 181d v1.25.3
          mjw-udoo-03 Ready worker 181d v1.25.3
          mjw-udoo-04 Ready worker 181d v1.25.3</p>
        <p>Configure Metal-LB to add these IP addresses to the IPAddressPool (REF: <a
            href="https://metallb.universe.tf/usage/example/">https://metallb.universe.tf/usage/example/</a>)</p>
        <pre><code>cat &lt;&lt;EOF | kubectl apply -f -
<span class="hljs-attr">apiVersion:</span> metallb.io/v1beta1
<span class="hljs-attr">kind:</span> IPAddressPool
<span class="hljs-attr">metadata:</span>
<span class="hljs-attr">  name:</span> udoo
<span class="hljs-attr">  namespace:</span> metallb-system
<span class="hljs-attr">spec:</span>
<span class="hljs-attr">  addresses:</span>
<span class="hljs-bullet">  -</span> <span class="hljs-number">192.168</span><span class="hljs-number">.0</span><span class="hljs-number">.91</span><span class="hljs-bullet">-192.168</span><span class="hljs-number">.0</span><span class="hljs-number">.94</span>
EOF
</code></pre>
        <p>Then, you can name your IPAddressPool and advertise it to the cluster. In my case, I&#39;ve hjust named it
          <code>udoo</code>
        </p>
        <pre><code>cat &lt;&lt;EOF | kubectl apply -f -
<span class="hljs-attr">apiVersion:</span> metallb.io/v1beta1
<span class="hljs-attr">kind:</span> BGPAdvertisement
<span class="hljs-attr">metadata:</span>
<span class="hljs-attr">  name:</span> external
<span class="hljs-attr">  namespace:</span> metallb-system
<span class="hljs-attr">spec:</span>
<span class="hljs-attr">  ipAddressPools:</span>
<span class="hljs-bullet">  -</span> udoo
EOF
</code></pre>
        <p>Then, create the namespace and deployment for the App. The following code creates the
          <code>k8s-golang-api</code> namespace for the app to run in and be identified with. It is up to you to
          choose a name that makes sense, but be sure to adjust the the following YAML snippets to reflect your
          Namespave name in <em>all</em> of the <code>namespace:</code> fields.
        </p>
        <h3 id="create-namespace">Create Namespace</h3>
        <pre><code>cat &lt;&lt;EOF | kubectl apply -f -
<span class="hljs-attr">apiVersion:</span> v1
<span class="hljs-attr">kind:</span> Namespace
<span class="hljs-attr">metadata:</span>
<span class="hljs-attr">  name:</span> k8s-golang-api
EOF
</code></pre>
        <h3 id="create-deployment">Create Deployment</h3>
        <p>The following defines how K8s will deploy the Pods on your system. It defines the names, associated
          Namespaces, number of Replicas, Resource Limits, Ports, etc.</p>
        <pre><code>cat &lt;&lt;EOF | kubectl apply -f -
<span class="hljs-attr">apiVersion:</span> apps/v1
<span class="hljs-attr">kind:</span> Deployment
<span class="hljs-attr">metadata:</span>
<span class="hljs-attr">  name:</span> k8s-golang-api
<span class="hljs-attr">  namespace:</span> k8s-golang-api
<span class="hljs-attr">  labels:</span>
<span class="hljs-attr">    app:</span> k8s-golang-api
<span class="hljs-attr">spec:</span>
<span class="hljs-attr">  replicas:</span> <span class="hljs-number">3</span>
<span class="hljs-attr">  selector:</span>
<span class="hljs-attr">    matchLabels:</span>
<span class="hljs-attr">      app:</span> k8s-golang-api
<span class="hljs-attr">  template:</span>
<span class="hljs-attr">    metadata:</span>
<span class="hljs-attr">      labels:</span>
<span class="hljs-attr">        app:</span> k8s-golang-api
<span class="hljs-attr">    spec:</span>
<span class="hljs-attr">      containers:</span>
<span class="hljs-attr">        - name:</span> k8s-golang-api
<span class="hljs-attr">          image:</span> <span class="hljs-string">'mattwiater/golangdocker:latest'</span>
<span class="hljs-attr">          env:</span>
<span class="hljs-attr">          - name:</span> K8S_NODE_NAME
<span class="hljs-attr">            valueFrom:</span>
<span class="hljs-attr">              fieldRef:</span>
<span class="hljs-attr">                fieldPath:</span> spec.nodeName
<span class="hljs-attr">          - name:</span> K8S_NODE_IP
<span class="hljs-attr">            valueFrom:</span>
<span class="hljs-attr">              fieldRef:</span>
<span class="hljs-attr">                fieldPath:</span> status.hostIP
<span class="hljs-attr">          - name:</span> K8S_POD_NAME
<span class="hljs-attr">            valueFrom:</span>
<span class="hljs-attr">              fieldRef:</span>
<span class="hljs-attr">                fieldPath:</span> metadata.name
<span class="hljs-attr">          - name:</span> K8S_POD_NAMESPACE
<span class="hljs-attr">            valueFrom:</span>
<span class="hljs-attr">              fieldRef:</span>
<span class="hljs-attr">                fieldPath:</span> metadata.namespace
<span class="hljs-attr">          - name:</span> K8S_POD_IP
<span class="hljs-attr">            valueFrom:</span>
<span class="hljs-attr">              fieldRef:</span>
<span class="hljs-attr">                fieldPath:</span> status.podIP
<span class="hljs-attr">          imagePullPolicy:</span> Always
<span class="hljs-attr">          resources:</span>
<span class="hljs-attr">            requests:</span>
<span class="hljs-attr">              memory:</span> <span class="hljs-string">"500Mi"</span>
<span class="hljs-attr">              cpu:</span> <span class="hljs-string">"250m"</span>
<span class="hljs-attr">            limits:</span>
<span class="hljs-attr">              memory:</span> <span class="hljs-string">"500Mi"</span>

<span class="hljs-attr">              cpu:</span> <span class="hljs-string">"250m"</span>
<span class="hljs-attr">          ports:</span>
<span class="hljs-attr">            - containerPort:</span> <span class="hljs-number">5000</span>
<span class="hljs-attr">              protocol:</span> TCP
EOF
</code></pre>
        <p>The final two steps, Service and Ingress are responsible for routing external traffic into the cluster.</p>
        <h3 id="create-service">Create Service</h3>
        <p>You can see that the service is accepting incoming traffic on port 80, and routing to the Pods named
          <code>k8s-golang-api</code> that are already running on Port 5000 (defined in the Deployment manifest above:
          <code>containerPort: 5000</code>)
        </p>
        <pre><code>cat &lt;&lt;EOF | kubectl apply -f -
<span class="hljs-attr">apiVersion:</span> v1
<span class="hljs-attr">kind:</span> Service
<span class="hljs-attr">metadata:</span>
<span class="hljs-attr">  name:</span> k8s-golang-api
<span class="hljs-attr">  namespace:</span> k8s-golang-api
<span class="hljs-attr">spec:</span>
<span class="hljs-attr">  type:</span> LoadBalancer
<span class="hljs-attr">  selector:</span>
<span class="hljs-attr">    app:</span> k8s-golang-api
<span class="hljs-attr">  ports:</span>
<span class="hljs-attr">  - name:</span> web
<span class="hljs-attr">    port:</span> <span class="hljs-number">80</span>
<span class="hljs-attr">    targetPort:</span> <span class="hljs-number">5000</span>
EOF
</code></pre>
        <h3 id="create-ingress">Create Ingress</h3>
        <p>In my setup, I want the containers to be accessible via Port 80 at the domain
          <code>golang.0nezer0.com</code>. So the Ingress section below defines the domain mapping to the Serice
          section
        </p>
        <pre><code>cat &lt;&lt;EOF | kubectl apply -f -
<span class="hljs-attr">apiVersion:</span> networking.k8s.io/v1
<span class="hljs-attr">kind:</span> Ingress
<span class="hljs-attr">metadata:</span>
<span class="hljs-attr">  name:</span> k8s-golang-api-ingress
<span class="hljs-attr">  namespace:</span> k8s-golang-api
<span class="hljs-attr">spec:</span>
<span class="hljs-attr">  defaultBackend:</span>
<span class="hljs-attr">    service:</span>
<span class="hljs-attr">      name:</span> k8s-golang-api
<span class="hljs-attr">      port:</span>
<span class="hljs-attr">        number:</span> <span class="hljs-number">80</span>
<span class="hljs-attr">  rules:</span>
<span class="hljs-attr">  - host:</span> golang<span class="hljs-number">.0</span>nezer0.com
<span class="hljs-attr">    http:</span>
<span class="hljs-attr">      paths:</span>
<span class="hljs-attr">      - path:</span> /
<span class="hljs-attr">        pathType:</span> Prefix
<span class="hljs-attr">        backend:</span>
<span class="hljs-attr">          service:</span>
<span class="hljs-attr">            name:</span> k8s-golang-api
<span class="hljs-attr">            port:</span>
<span class="hljs-attr">              number:</span> <span class="hljs-number">80</span>
EOF
</code></pre>
        <p>You can verify the setup via:</p>
        <p><code>kubectl describe ingress k8s-golang-api-ingress -n=k8s-golang-api</code></p>
        <pre><code><span class="hljs-symbol">Name:</span>             k8s-golang-api-ingress
<span class="hljs-symbol">Labels:</span>           &lt;none&gt;
<span class="hljs-symbol">Namespace:</span>        k8s-golang-api
<span class="hljs-symbol">Address:</span>
Ingress Class:    &lt;none&gt;
<span class="hljs-meta">Default</span> backend:  k8s-golang-api:<span class="hljs-number">80</span> (<span class="hljs-number">10.244</span><span class="hljs-meta">.1</span><span class="hljs-meta">.74</span>:<span class="hljs-number">5000</span>,<span class="hljs-number">10.244</span><span class="hljs-meta">.2</span><span class="hljs-meta">.104</span>:<span class="hljs-number">5000</span>,<span class="hljs-number">10.244</span><span class="hljs-meta">.3</span><span class="hljs-meta">.66</span>:<span class="hljs-number">5000</span>)
<span class="hljs-symbol">Rules:</span>
  Host                Path  Backends
  ----                ----  --------
  golang.0nezer0.com
                      /   k8s-golang-api:<span class="hljs-number">80</span> (<span class="hljs-number">10.244</span><span class="hljs-meta">.1</span><span class="hljs-meta">.74</span>:<span class="hljs-number">5000</span>,<span class="hljs-number">10.244</span><span class="hljs-meta">.2</span><span class="hljs-meta">.104</span>:<span class="hljs-number">5000</span>,<span class="hljs-number">10.244</span><span class="hljs-meta">.3</span><span class="hljs-meta">.66</span>:<span class="hljs-number">5000</span>)
<span class="hljs-symbol">Annotations:</span>          &lt;none&gt;
<span class="hljs-symbol">Events:</span>               &lt;none&gt;
</code></pre>
        <p>Note that the domain is listed and the Backend are pointing to the Service we created.</p>
        <p>Ensure that you have an IP Address allocated for the Load Balancer:</p>
        <p><code>kubectl get svc -n=k8s-golang-api</code></p>
        <pre><code>NAME             TYPE           CLUSTER-<span class="hljs-built_in">IP</span>      EXTERNAL-<span class="hljs-built_in">IP</span>    PORT(S)        AGE
k8s-golang-api   LoadBalancer   <span class="hljs-number">10.105</span><span class="hljs-meta">.31</span><span class="hljs-meta">.196</span>   <span class="hljs-number">192.168</span><span class="hljs-meta">.0</span><span class="hljs-meta">.91</span>   <span class="hljs-number">80</span>:<span class="hljs-number">31188</span>/TCP   21s
</code></pre>
        <p>Assuming that your setup is also on your local network, make sure to add add an IP -&gt; Domain mapping in
          <code>/etc/hosts</code> file on the machine you are accessing the cluster from:
        </p>
        <pre><code>192<span class="hljs-selector-class">.168</span><span class="hljs-selector-class">.0</span><span class="hljs-selector-class">.91</span> <span class="hljs-selector-tag">golang</span><span class="hljs-selector-class">.0nezer0</span><span class="hljs-selector-class">.com</span>
</code></pre>
        <h2 id="horizontal-pod-autoscaler-hpa-">Horizontal Pod Autoscaler (HPA)</h2>
        <p>
          <code>kubectl autoscale deployment -n k8s-golang-api  k8s-golang-api --cpu-percent=75 --memory-percent=75 --min=1 --max=3</code>
        </p>

      </div>
    </div>
    <footer class="site-footer h-card">
      FOOTER
    </footer>
  </div>

  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js"
    integrity="sha384-kenU1KFdBIe4zVF0s0G1M5b4hcpxyD9F7jL+jjXkk+Q2h455rYXK/7HAuoJl+0I4"
    crossorigin="anonymous"></script>

</body>

</html>