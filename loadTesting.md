---
layout: null
---
{% include header.html %}

  <div class="container">
    <div class="row">
      <div class="col">
        <h1 id="artillery">ARTILLERY</h1>
        <p>[IN PROGRESS]</p>
        <h2 id="to-do">To Do</h2>
        <ul>
          <li>Explore custom metrics options in more depth.</li>
          <li>Generate applicable reports for comparison between bare go app, dockerized app, and k8s replicas.
          </li>
        </ul>
        <h2 id="install">Install</h2>
        <p><code>npm install -g artillery@latest</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </p>
        <h2 id="plugins">Plugins</h2>
        <p><a
            href="https://www.artillery.io/docs/guides/plugins/plugin-metrics-by-endpoint#useonlyrequestnames">Official:
            Per-endpoint (URL) metrics</a></p>
        <p><code>npm install artillery-plugin-metrics-by-endpoint</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </p>

        <h2 id="test-phases">Test Phases</h2>

        <p>
        <pre><code>golangdocker-loadtest.yml</code></pre>
        </p>

        <pre>
<code>config:
    phases:
      - duration: 60
        arrivalRate: 5
        name: Warm up
      - duration: 120
        arrivalRate: 5
        rampTo: 50
        name: Ramp up load
      - duration: 600
        arrivalRate: 50
        name: Sustained load
    plugins:
      metrics-by-endpoint:
        useOnlyRequestNames: false
    processor: "custom-artillery-functions.js"
  scenarios:
    - name: "golang.0nezer0.com"
      flow:
  
      - get:
          url: "/v1"
          afterResponse: "customMetrics"
  
      - get:
          url: "/v1/cpu"
          afterResponse: "customMetrics"
  
      - get:
          url: "/v1/host"
          capture:
            - json: "$['hostInfo']['virtualizationSystem']"
              as: "virtualizationSystem"
            - json: "$['hostInfo']['hostname']"
              as: "hostname"
          afterResponse: "customMetrics"
      # - log: "{{ hostname }} [{{ virtualizationSystem }}]" # Here to ensure we are correctly load-balancing different pods in K8s deployment
  
      - get:
          url: "/v1/load"
          afterResponse: "customMetrics"
  
      - get:
          url: "/v1/mem"
          afterResponse: "customMetrics"
  
      - get:
          url: "/v1/net"
          afterResponse: "customMetrics"
</code>
</pre>


        <h2 id="custom-scripts">Custom Scripts</h2>
        <p>Reference: <a
            href="https://www.artillery.io/docs/guides/guides/extension-apis#example">https://www.artillery.io/docs/guides/guides/extension-apis#example</a>
        </p>
        <p>This simple example makes use of a custom Fiber middleware wrapper that captures the time spent on the
          server in each API call and sets a <code>Server-Timing</code> response header, e.g.:
          <code>Server-Timing: route;dur=16</code>. See the <a
            href="https://github.com/mwiater/golangdocker/blob/master/api/api.go">RouteTimerHandler()</a> function in <code>api/api.go</code>.
        </p>

        <p>
        <pre><code>custom-artillery-functions.js</code></pre>
        </p>

        <pre>
<code>//
// custom-artillery-functions.js
//

module.exports = {
  logHeaders: logHeaders,
  customMetrics: customMetrics
}

function logHeaders(requestParams, response, context, events, next) {
  // console.log(response.headers);
  return next();
}

function customMetrics(requestParams, response, context, events, next) {
  const latency = parseServerTimingLatency(response.headers["server-timing"], "route");
  const url = new URL(requestParams.url);
  const routePath = url.pathname.replaceAll("/", "_")
  events.emit("histogram", "route_latency"+routePath.trim(), latency);
  return next();
}

function parseServerTimingLatency(header, timingMetricName) {
  const serverTimings = header.split(",");

  for (let timing of serverTimings) {
    const timingDetails = timing.split(";");
    if (timingDetails[0] === timingMetricName) {
      return parseFloat(timingDetails[1].split("=")[1]);
    }
  }
}
</code>
</pre>

        <h2 id="load-tests">Load Tests</h2>
        <p>In order to benchmark the different run processes, we need to start the app differently before sending a
          load test. You will alos want to run these test form a different physical machine that where you&#39;re
          running the container from. Keep in mind that these are not real world load tests, as we are mostly testing
          to targets within the same network. These tests are mainly for comparisons of ruinning the app with
          different mechanisims, e.g: go app, inside Docker container, within K8s w/ replicas.</p>
        <h2 id="no-container-bare-app">No container, bare app</h2>
        <p>E.g.: <code>bash go_run.sh</code></p>
        <pre><code><span class="hljs-keyword">clear</span> &amp;&amp; \
artillery <span class="hljs-keyword">run</span> --<span class="hljs-keyword">output</span> golangdocker-bare.json --target http:<span class="hljs-comment">//192.168.0.91:5000/api golangdocker-loadtest.yml &amp;&amp; \</span>
    artillery report golangdocker-bare.json
</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </pre>
        <h2 id="docker-container">Docker container</h2>
        <p>E.g.: <code>bash docker_run.sh</code></p>
        <pre><code><span class="hljs-keyword">clear</span> &amp;&amp; \
artillery <span class="hljs-keyword">run</span> --<span class="hljs-keyword">output</span> golangdocker-docker.json --target http:<span class="hljs-comment">//192.168.0.91:5000/api golangdocker-loadtest.yml &amp;&amp; \</span>
    artillery report golangdocker-docker.json</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </pre>
        <h2 id="k8s">K8s</h2>
        <p>Assumes working K8s cluster and manual scaling of replicas for each test, e.g.:</p>
        <pre><code><span class="hljs-keyword">clear</span> &amp;&amp; \
artillery <span class="hljs-keyword">run</span> --<span class="hljs-keyword">output</span> golangdocker-k8s<span class="hljs-number">-3</span>-replica.json --target http:<span class="hljs-comment">//golang.0nezer0.com/api golangdocker-loadtest.yml &amp;&amp; \</span>
    artillery report golangdocker-k8s<span class="hljs-number">-3</span>-replica.json
</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </pre>
        <pre><code><span class="hljs-keyword">clear</span> &amp;&amp; \
artillery <span class="hljs-keyword">run</span> --<span class="hljs-keyword">output</span> golangdocker-k8s<span class="hljs-number">-2</span>-replica.json --target http:<span class="hljs-comment">//golang.0nezer0.com/api golangdocker-loadtest.yml &amp;&amp; \</span>
    artillery report golangdocker-k8s<span class="hljs-number">-2</span>-replica.json
</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </pre>

      </div>
    </div>
  </div>

{% include footer.html %}