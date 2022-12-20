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
        <h1 id="tests">Tests</h1>
        <p>Very simple tests are in: <a href="../../../../blob/master/api_test.go">api_test.go</a></p>
        <p>Run via:
          <code>clear &amp;&amp; go test -v $(go list ./... | grep -v /docs | grep -v /config | grep -v /api)</code>
          #=&gt;
        </p>
        <pre><code>=== RUN   TestAPIRoutes

 ┌───────────────────────────────────────────────────┐
 │                   Fiber v2<span class="hljs-number">.40</span><span class="hljs-number">.1</span>                   │
 │               <span class="hljs-string">http:</span><span class="hljs-comment">//127.0.0.1:5000               │</span>
 │       (bound on host <span class="hljs-number">0.0</span><span class="hljs-number">.0</span><span class="hljs-number">.0</span> and port <span class="hljs-number">5000</span>)       │
 │                                                   │
 │ Handlers ............ <span class="hljs-number">22</span>  Processes ........... <span class="hljs-number">1</span> │
 │ Prefork ....... Disabled  PID ........... <span class="hljs-number">2214242</span> │
 └───────────────────────────────────────────────────┘

[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span>/: <span class="hljs-number">302</span> (     <span class="hljs-number">0</span>s) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">0</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span><span class="hljs-string">v1:</span> <span class="hljs-number">200</span> (     <span class="hljs-number">0</span>s) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">136</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1/<span class="hljs-string">cpu:</span> <span class="hljs-number">200</span> (    <span class="hljs-number">1</span>ms) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">3593</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1/<span class="hljs-string">host:</span> <span class="hljs-number">200</span> (    <span class="hljs-number">1</span>ms) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">338</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1/<span class="hljs-string">load:</span> <span class="hljs-number">200</span> (     <span class="hljs-number">0</span>s) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">54</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1/<span class="hljs-string">mem:</span> <span class="hljs-number">200</span> (    <span class="hljs-number">1</span>ms) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">706</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1/<span class="hljs-string">net:</span> <span class="hljs-number">200</span> (    <span class="hljs-number">2</span>ms) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">1559</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1/<span class="hljs-string">metrics:</span> <span class="hljs-number">200</span> (     <span class="hljs-number">0</span>s) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">6186</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1<span class="hljs-regexp">/docs/</span>index.<span class="hljs-string">html:</span> <span class="hljs-number">200</span> (     <span class="hljs-number">0</span>s) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">3519</span>
[<span class="hljs-number">2022</span><span class="hljs-number">-12</span><span class="hljs-number">-13</span><span class="hljs-string">T11:</span><span class="hljs-number">01</span>:<span class="hljs-number">06</span>] <span class="hljs-string">GET:</span><span class="hljs-regexp">/api/</span>v1/<span class="hljs-number">404</span>: <span class="hljs-number">404</span> (     <span class="hljs-number">0</span>s) | Bytes <span class="hljs-string">In:</span> <span class="hljs-number">0</span> Bytes <span class="hljs-string">Out:</span> <span class="hljs-number">22</span>
--- <span class="hljs-string">PASS:</span> TestAPIRoutes (<span class="hljs-number">0.13</span>s)
PASS
ok      github.com<span class="hljs-regexp">/mattwiater/</span>golangdocker      <span class="hljs-number">0.190</span>s
=== RUN   ExamplePrettyPrintJSONToConsole
--- <span class="hljs-string">PASS:</span> ExamplePrettyPrintJSONToConsole (<span class="hljs-number">0.00</span>s)
=== RUN   ExampleUniqueSlice
--- <span class="hljs-string">PASS:</span> ExampleUniqueSlice (<span class="hljs-number">0.00</span>s)
PASS
ok      github.com<span class="hljs-regexp">/mattwiater/</span>golangdocker/common       <span class="hljs-number">0.005</span>s
=== RUN   ExampleTestTZ
--- <span class="hljs-string">PASS:</span> ExampleTestTZ (<span class="hljs-number">0.00</span>s)
=== RUN   ExampleTestTLS
--- <span class="hljs-string">PASS:</span> ExampleTestTLS (<span class="hljs-number">0.35</span>s)
PASS
ok      github.com<span class="hljs-regexp">/mattwiater/</span>golangdocker/sysinfo      <span class="hljs-number">0.365</span>s
</code></pre>
        <h2 id="test-cache">Test Cache</h2>
        <p>To clear the test cache, run: <code>go clean -testcache</code></p>

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