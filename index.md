---
layout: null
---
{% include header.html %}
    <div class="row">
      <div class="col">
        <h1 id="golangdocker">Golangdocker</h1>
        <table>
          <thead></thead>
          <tbody>
            <tr>
              <th style="text-align:center"><img
                  src="https://raw.githubusercontent.com/mwiater/golangdocker/master/_repository_docs/_repository_assets/logo-golang.png?raw=true"
                  alt="&quot;Go&quot;" title="Go"></th>
              <th style="text-align:center"><img
                  src="https://raw.githubusercontent.com/mwiater/golangdocker/master/_repository_docs/_repository_assets/logo-docker.png?raw=true"
                  alt="&quot;Docker&quot;" title="Docker"></th>
              <th style="text-align:center"><img
                  src="https://raw.githubusercontent.com/mwiater/golangdocker/master/_repository_docs/_repository_assets/logo-golang-fiber.png?raw=true"
                  alt="&quot;Fiber&quot;" title="Fiber"></th>
              <th style="text-align:center"><img
                  src="https://raw.githubusercontent.com/mwiater/golangdocker/master/_repository_docs/_repository_assets/logo-k8s.png?raw=true"
                  alt="&quot;Kubernetes&quot;" title="Kubernetes"></th>
            </tr>
          </tbody>
        </table>
        <h2 id="table-of-contents">Table of Contents</h2>
        <ul>
          <li><a href="#project-summary">Project Summary</a></li>
          <li><a href="#project-topics">Project Topics</a></li>
          <li><a href="#assumptions">Project Assumptions</a></li>
          <li><a href="#prerequisites">Project Prerequisites</a></li>
          <li>Jump directly into the Project <a href="#app">app code</a></li>
        </ul>
        <h2 id="summary">Summary</h2>
        <p>This repository is a work in progress, but I&#39;ll do my best to keep the Master branch in a working
          state. Initially, this project was to create a boilerplate for containerizing Go binaries for use in a K8s
          cluster. For now, just origanizing my notes in order to be able to replicate this process from end-to-end.
          The idea is to keep this narrow and succinct and be able to use this as a simple boilerplate for Go
          containers.</p>
        <p>While my intention is to keep this image minimal, I also want to add in my most used go tools, e.g.:
          <code>godoc</code>, <code>gofmt</code>, <code>golangci-lint</code>, etc. I will try to make these additions
          for only tools and packages I find necessary, but it will be an opinionated balance.
        </p>
        <h2 id="project-topics">Project Topics</h2>
        <p>This project is in three distinct parts, each which build on the previous:</p>
        <h5 id="1-a-simple-but-functioanl-rest-api-app-written-in-go-this-rest-api-incorporates-">1) A simple but
          functioanl rest API app written in Go. This rest API incorporates:</h5>
        <ul>
          <li>The <a href="https://docs.gofiber.io/api/middleware/monitor">Fiber Monitor middleware</a> (API endpoint:
            <code>/api/v1/metrics</code>).
          </li>
          <li>Creating and serving API documentation (using <code>swag init</code>) based on <a
              href="https://swagger.io/docs/specification/about/">Swagger specifications</a>:
            <code>/api/v1/docs/</code>).
          </li>
          <li>A <code>YAML</code> configuration pattern for setting app variables.</li>
          <li>Basic Go endpoint tests via <code>go test</code>.</li>
          <li>Building a binary of the app and embedding external files (both native compilation and cross-compilation
            for <code>armv6</code> as an example) so that it is portable and self contained.</li>
          <li>Go Tools<ul>
              <li>File formatting for *.go files using <code>gofmt</code>.</li>
              <li>Code linting for *.go files using <code>golangci-lint</code>.</li>
              <li>Code documentation via <code>godocs</code>.</li>
            </ul>
          </li>
        </ul>
        <h5 id="2-using-the-app-in-a-docker-container-covering-">2) Using the app in a Docker container, covering:
        </h5>
        <ul>
          <li><a href="https://docs.docker.com/engine/reference/commandline/build/">Docker build</a> concepts.</li>
          <li><a href="https://docs.docker.com/engine/reference/commandline/run/">Docker run</a> concepts.</li>
          <li>Docker image versioning.</li>
          <li>Ways to make use of bash scripts for repetative tasks.</li>
        </ul>
        <h5 id="3-using-the-docker-container-in-kubernetes">3) Using the Docker container in Kubernetes</h5>
        <ul>
          <li>This section is the most incomplete, but should be in a working state.</li>
          <li>You should already have a working K8s cluster available for this section.</li>
          <li>Does not provide much background, assumes some basic knowledge using <code>kubectl</code>.</li>
          <li>This app will be deployed as a load-balanced Service across a Control Plane and 3 Worker nodes.</li>
        </ul>
        <h2 id="assumptions">Assumptions</h2>
        <ul>
          <li><strong>IP Addresses:</strong> For the most part, disregard the hard-coded IP addresses in here (e.g.:
            my K8s cluster and VM IPs (192.168.<em>.</em>)). You&#39;ll have to sub in your own for your particular
            envionment. Right now, laziness!</li>
          <li><strong>Container vs. Pod:</strong> I&#39;m noticing a few instances where I&#39;m using both
            <code>container</code> and <code>pod</code> to mean the same thing in the K8s section. Until I make them
            more consistent, assume they are interchangeable. A pod is basiically a container in in K8s context. While
            a <code>pod</code> can technically have multiple containers, for this demonstration, assume a 1:1
            relationship.
          </li>
          <li><strong>System</strong> My system and architecture is below, you&#39;ll have to adjust your commands if
            you&#39;re departing from Linux/x86_64.</li>
        </ul>
        <p><code>uname -a</code></p>
        <pre><code><span class="hljs-selector-tag">Linux</span> <span class="hljs-selector-tag">mjw-udoo-01</span> <span class="hljs-selector-tag">5</span><span class="hljs-selector-class">.4</span><span class="hljs-selector-class">.0-110-generic</span> <span class="hljs-selector-id">#124-Ubuntu</span> <span class="hljs-selector-tag">SMP</span> <span class="hljs-selector-tag">Thu</span> <span class="hljs-selector-tag">Apr</span> <span class="hljs-selector-tag">14</span> <span class="hljs-selector-tag">19</span><span class="hljs-selector-pseudo">:46</span><span class="hljs-selector-pseudo">:19</span> <span class="hljs-selector-tag">UTC</span> <span class="hljs-selector-tag">2022</span> <span class="hljs-selector-tag">x86_64</span> <span class="hljs-selector-tag">x86_64</span> <span class="hljs-selector-tag">x86_64</span> <span class="hljs-selector-tag">GNU</span>/<span class="hljs-selector-tag">Linux</span>
</code></pre>
        <h2 id="to-do">To Do</h2>
        <ul>
          <li>[x] Cross-compiling: Add scripts: <code>go_build.sh</code> and <code>go_build_arm64.sh</code> as
            examples.</li>
          <li>[ ] Generate Postman collection for reference?</li>
          <li>[ ] Turn these to-dos into issues!</li>
          <li>[ ] K8s: Use version tagging instead of <code>:latest</code> to provide an example of rolling updates.
            (Started: <a href="../../blob/master/_repository_docs/_k8s/K8S_README.md">K8S_README.md</a>)</li>
          <li>[ ] TLS? In single container or via K8s? To update in: <a href="../../blob/master/certs/">certs/</a>
          </li>
        </ul>
        <h2 id="prerequisites">Prerequisites</h2>
        <p>The following programs will need to be installed:</p>
        <ul>
          <li><a href="https://go.dev/learn/">Go</a></li>
          <li><a href="https://www.docker.com/get-started/">Docker</a></li>
        </ul>
        <p>You must set the DOCKERIMAGE environment variable when running the Docker Bash scripts. For only the
          current session, just enter:
          <code>export DOCKERIMAGE={your-docker-hub-account-username}/{your-docker-hub-image-name}</code>, e.g.:
          <code>export DOCKERIMAGE=mattwiater/golangdocker</code>. If you want that var to be permanent accross ssh
          sessions, add the export line to your ~/.bashrc (or equivalent) file.
        </p>
        <p>Required for Kubernetes itegration:</p>
        <ul>
          <li>A running <a href="https://kubernetes.io/">Kubernetes</a> cluster</li>
          <li>A <a href="https://hub.docker.com/">Docker Hub</a> account</li>
        </ul>
        <p>Optional:</p>
        <ul>
          <li>Artillery (nodejs): <a href="../../blob/master/_repository_docs/_loadTesting/ARTILLERY.md">Load
              Testing</a></li>
        </ul>
        <p>While the idea is to get this up and running quickly, it is not a deep dive into Go, Docker, or K8S. Basic
          knowledge of these technologies is required.</p>
        <p>For example, we can peek into the container via the API endpoint <code>api/v1/host</code> and see the
          docker assigned <code>hostname: &quot;b189564db0c5&quot;</code> and verify that it is one running a single
          process <code>procs: 1</code>:</p>
        <pre><code>{
<span class="hljs-symbol">hostInfo:</span> {
<span class="hljs-symbol">  hostname:</span> <span class="hljs-string">"b189564db0c5"</span>,
<span class="hljs-symbol">  uptime:</span> <span class="hljs-number">1238849</span>,
<span class="hljs-symbol">  bootTime:</span> <span class="hljs-number">1667920883</span>,
<span class="hljs-symbol">  procs:</span> <span class="hljs-number">1</span>,
<span class="hljs-symbol">  os:</span> <span class="hljs-string">"linux"</span>,
<span class="hljs-symbol">  platform:</span> <span class="hljs-string">""</span>,
<span class="hljs-symbol">  platformFamily:</span> <span class="hljs-string">""</span>,
<span class="hljs-symbol">  platformVersion:</span> <span class="hljs-string">""</span>,
<span class="hljs-symbol">  kernelVersion:</span> <span class="hljs-string">"5.4.0-110-generic"</span>,
<span class="hljs-symbol">  kernelArch:</span> <span class="hljs-string">"x86_64"</span>,
<span class="hljs-symbol">  virtualizationSystem:</span> <span class="hljs-string">"docker"</span>,
<span class="hljs-symbol">  virtualizationRole:</span> <span class="hljs-string">"guest"</span>,
<span class="hljs-symbol">  hostId:</span> <span class="hljs-string">"12345678-1234-5678-90ab-cddeefaabbcc"</span>
  }
}
</code></pre>
        <h2 id="makefile">Makefile</h2>
        <p>[IN PROGRESS]</p>
        <p>There is a <code>Makefile</code> for convenience. At the moment, it&#39;s just acting as a script-runner.
          To viee the executable tagets, just type: <code>make</code>:</p>
        <pre><code>Targets in this Makefile:

docker-build
docker-<span class="hljs-keyword">run</span><span class="bash">
golang-build
</span>golang-build-arm64
golang-<span class="hljs-keyword">run</span><span class="bash"></span>
</code></pre>
        <h2 id="app">App</h2>
        <h3 id="config">Config</h3>
        <p>There is a simple app config pattern using: <code>./config/appConfig.yml</code></p>
        <pre><code><span class="hljs-comment"># config.yml</span>

<span class="hljs-attr">server:</span>
<span class="hljs-attr">  port:</span> <span class="hljs-number">5000</span>
<span class="hljs-attr">options:</span>
<span class="hljs-attr">  debug:</span> <span class="hljs-literal">false</span>
</code></pre>
        <p>Keeping this simple for now, just want to have a boilerplate config pattern within the app for future use.
        </p>
        <ul>
          <li>Port: The Port that the app listens on, deafult: <code>5000</code></li>
          <li>Debug: More console output on API requests, deafult: <code>false</code></li>
        </ul>
        <p>For <code>debug</code>, this will print out the JSON response to the console, depending on the endpoint
          requested. For <code>/api/v1/host</code>, you get something like this:</p>
        <pre><code>[ ★ <span class="hljs-meta">INFO</span> ] Host <span class="hljs-meta">Info</span>:
{
        <span class="hljs-string">"hostname"</span>: <span class="hljs-string">"mjw-udoo-01"</span>,
        <span class="hljs-string">"uptime"</span>: <span class="hljs-number">11093</span>,
        <span class="hljs-string">"bootTime"</span>: <span class="hljs-number">1669484114</span>,
        <span class="hljs-string">"procs"</span>: <span class="hljs-number">176</span>,
        <span class="hljs-string">"os"</span>: <span class="hljs-string">"linux"</span>,
        <span class="hljs-string">"platform"</span>: <span class="hljs-string">"ubuntu"</span>,
        <span class="hljs-string">"platformFamily"</span>: <span class="hljs-string">"debian"</span>,
        <span class="hljs-string">"platformVersion"</span>: <span class="hljs-string">"20.04"</span>,
        <span class="hljs-string">"kernelVersion"</span>: <span class="hljs-string">"5.4.0-110-generic"</span>,
        <span class="hljs-string">"kernelArch"</span>: <span class="hljs-string">"x86_64"</span>,
        <span class="hljs-string">"virtualizationSystem"</span>: <span class="hljs-string">"kvm"</span>,
        <span class="hljs-string">"virtualizationRole"</span>: <span class="hljs-string">"host"</span>,
        <span class="hljs-string">"hostId"</span>: <span class="hljs-string">"3a114467-105a-48a5-9419-32654a9b2076"</span>
}
</code></pre>
        <h3 id="testing-developing-app">Testing/Developing App</h3>
        <p>while developing/testing the app, you can run it natively (not in a Docker container) via:</p>
        <p><code>go run main.go</code></p>
        <p>Or, for convenience and formatting, run: <code>bash go_run.sh</code></p>
        <p>Site will be available at: <a href="http://192.168.0.91:5000/api/v1">http://192.168.0.91:5000/api/v1</a>
          (substitute your own IP address)</p>
        <p>This step should be completed first before running via Docker to ensure everything is working properly.</p>
        <h3 id="building-the-docker-container">Building the Docker container</h3>
        <p>NOTE: The steps will refer to the docker image: <code>mattwiater/golangdocker</code>. You should change
          these steps to match your own image name, e.g.: <code>{your-docker-hub-account-username}/golangdocker</code>
        </p>
        <p>The build command uses the local <a href="../../blob/master/Dockerfile">Dockerfile</a> to build the image.
          Substitute your own Docker image tag for mine wherever you see it (<code>mattwiater/golangdocker</code>),
          e.g.: <code>{your-docker-hub-account-username}/golangdocker</code></p>
        <p><code>docker build -t mattwiater/golangdocker .</code></p>
        <p>Or, for convenience, run:
          <code>bash docker_build.sh &#39;{your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}&#39;</code>
        </p>
        <p>Once you have built your image successfully, check the output of <code>docker images</code> #=&gt;</p>
        <pre><code>REPOSITORY                <span class="hljs-keyword">TAG</span>       <span class="hljs-title">IMAGE</span> ID       CREATED          SIZE
mattwiater/golangdocker   latest    <span class="hljs-number">053</span>f21052659   <span class="hljs-number">10</span> minutes ago   <span class="hljs-number">10.7M</span>B
...
</code></pre>
        <p>You should see your tagged image in the list, similar to the output above.</p>
        <h2 id="docker-build-notes">Docker Build Notes</h2>
        <p>Using <a href="https://docs.docker.com/build/building/multi-stage/#use-multi-stage-builds">multi-stage
            builds</a>, we will use a very simple <code>Dockerfile</code> to containerize our app.</p>
        <pre><code><span class="hljs-keyword">FROM</span> golang:alpine as app
<span class="hljs-keyword">WORKDIR</span><span class="bash"> /go/src/app
</span><span class="hljs-keyword">COPY</span><span class="bash"> . .
</span><span class="hljs-keyword">RUN</span><span class="bash"> apk add git
</span><span class="hljs-keyword">RUN</span><span class="bash"> CGO_ENABLED=0 go install -ldflags <span class="hljs-string">'-extldflags "-static"'</span> -tags timetzdata
</span>
<span class="hljs-keyword">FROM</span> scratch
<span class="hljs-keyword">COPY</span><span class="bash"> --from=app /go/bin/golangdocker /golangdocker
</span><span class="hljs-keyword">COPY</span><span class="bash"> --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
</span><span class="hljs-keyword">ENTRYPOINT</span><span class="bash"> [<span class="hljs-string">"/golangdocker"</span>]</span>
</code></pre>
        <h3 id="running-the-docker-container">Running the Docker container</h3>
        <p>Start the container in an interactive shell, with the host port <code>5000</code> (the machine you&#39;re
          running Docker on) mapping to the container port (the port the app is running on within the Docker
          container) <code>5000</code> for simplicity. The app is Port if configured here:
          <code>./config/appConfig.yml</code>
        </p>
        <p><code>docker run -it -p 5000:5000 --rm mattwiater/golangdocker</code></p>
        <p>Or, for convenience, run:
          <code>bash docker_run.sh &#39;{your-docker-hub-account-username}/{your-docker-hub-image-name}{:optional-version}&#39;</code>.
        </p>
        <p>You should see the default Fiber message, e.g.:</p>
        <pre><code> ┌───────────────────────────────────────────────────┐
 │                   Fiber v2<span class="hljs-number">.40</span><span class="hljs-number">.0</span>                   │
 │               http:<span class="hljs-comment">//127.0.0.1:5000               │</span>
 │       (bound on host <span class="hljs-number">0.0</span><span class="hljs-number">.0</span><span class="hljs-number">.0</span> and port <span class="hljs-number">5000</span>)       │
 │                                                   │
 │ Handlers ............ <span class="hljs-number">14</span>  Processes ........... <span class="hljs-number">1</span> │
 │ Prefork ....... Disabled  PID ................. <span class="hljs-number">1</span> │
 └───────────────────────────────────────────────────┘
</code></pre>
        <p>On your host machine, you can now access the container via <code>http://{your-host-ip-address}:5000</code>
        </p>
        <p>Our build is simple, just a compiled Go binary that runs in a container. This binary collects local
          resources/stats for display as JSON via these API Endpoints using <a
            href="https://docs.gofiber.io/">Fiber</a>:</p>
        <h5 id="api-info-">API Info:</h5>
        <pre><code><span class="hljs-symbol">http:</span><span class="hljs-comment">//{your-host-ip-address}:5000/api/v1</span>
</code></pre>
        <h5 id="system-info-">System Info:</h5>
        <pre><code><span class="hljs-symbol">http:</span><span class="hljs-comment">//{your-host-ip-address}:5000/api/v1/mem</span>
<span class="hljs-symbol">http:</span><span class="hljs-comment">//{your-host-ip-address}:5000/api/v1/cpu</span>
<span class="hljs-symbol">http:</span><span class="hljs-comment">//{your-host-ip-address}:5000/api/v1/host</span>
<span class="hljs-symbol">http:</span><span class="hljs-comment">//{your-host-ip-address}:5000/api/v1/net</span>
<span class="hljs-symbol">http:</span><span class="hljs-comment">//{your-host-ip-address}:5000/api/v1/load</span>
</code></pre>
        <h5 id="api-metrics-">API Metrics:</h5>
        <p>For simplicity, the default <a href="https://docs.gofiber.io/api/middleware/monitor">Fiber Monitor
            middleware</a> is included and available at:</p>
        <p><code>http://{your-host-ip-address}:5000/api/v1/metrics</code></p>
        <h5 id="api-endpoint-documentation-via-swagger">API Endpoint Documentation via Swagger</h5>
        <p><code>go install github.com/swaggo/swag/cmd/swag@latest</code></p>
        <p><code>go get -u github.com/swaggo/fiber-swagger</code></p>
        <p>When updating documentation, you must run this to regenerate docs data: <code>swag init</code>
          (<code>swag init</code> is incorporated into the bash scripts for convenience, e.g.: <a
            href="../../blob/master/docker_run.sh">docker_run.sh</a>)</p>
        <p>Then, when you run the application, docs are avaialble at:</p>
        <p><code>http://{your-host-ip-address}:5000/api/v1/docs/index.html</code></p>
        <h2 id="tests">Tests</h2>
        <p>See: <a href="../../blob/master/_repository_docs/_tests/TESTS.md">Tests</a></p>
        <h2 id="linting-code-analysis">Linting: Code analysis</h2>
        <p>Basic linting option via <code>golangci-lint</code></p>
        <p>See: <a href="../../blob/master/_repository_docs/_linting/LINTING.md">Linting</a></p>
        <h2 id="gosec-security-analysis">Gosec: Security analysis</h2>
        <p>High level gosec usage example.</p>
        <p>See: <a href="../../blob/master/_repository_docs/_gosec/GOSEC.md">Gosec</a></p>
        <h2 id="-to-do-notes">[TO DO] Notes</h2>
        <p>Assumptions:</p>
        <p>Since we initialized the project with: </p>
        <p><code>go mod init github.com/mattwiater/golangdocker</code></p>
        <p>And each package is in it&#39;s own directory: <code>sysinfo</code>, <code>api</code>, <code>common</code>,
          etc., in order to use these local packages within the <code>main</code> Go package, you must enter each
          directory and type: <code>go build</code></p>
        <p>Then, in <code>main.go</code>, you can include them like this:</p>
        <pre><code>...
<span class="hljs-string">"github.com/mattwiater/golangdocker/sysinfo"</span>
<span class="hljs-string">"github.com/mattwiater/golangdocker/api"</span>
<span class="hljs-string">"github.com/mattwiater/golangdocker/common"</span>
...
</code></pre>
        <p><strong>Note on local packages:</strong> In order to make use of your local package functions, along with
          running the <code>go build</code> command, ensure that your functions are Capital-cased. Otherwise Go will
          throw an error saying that your method is undefined. Only functions that begin with a capital letter are
          exported from packages, otherwise they are considered private.</p>

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