---
layout: null
---
{% include header.html %}

  <div class="container">
    <div class="row">
      <div class="col">
        <h1 id="gosec">Gosec</h1>
        <p>[IN PROGRESS]</p>
        <h2 id="to-do">To Do</h2>
        <ul>
          <li>Results same/similar things to <code>golangci-lint</code>, redundant?</li>
        </ul>
        <h2 id="install-to-bin-">Install to <code>./bin/</code></h2>
        <p><code>curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sh -s v2.14.0</code>  <i class="fa-duotone fa-copy fa-fw code-copy-button"></i>
        </p>
        <h2 id="run-with-all-options-">Run with all options:</h2>
        <p><code>bin/gosec ./...</code>  <i class="fa-duotone fa-copy fa-fw code-copy-button"></i>  &nbsp; #=&gt;</p>
        <pre><code>[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">33</span> Including <span class="hljs-string">rules:</span> <span class="hljs-keyword">default</span>
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">33</span> Excluding <span class="hljs-string">rules:</span> <span class="hljs-keyword">default</span>
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">33</span> Import <span class="hljs-string">directory:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">33</span> Import <span class="hljs-string">directory:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker/api
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">33</span> Import <span class="hljs-string">directory:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker/common
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">33</span> Import <span class="hljs-string">directory:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker/config
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">34</span> Import <span class="hljs-string">directory:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker/docs
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">34</span> Checking <span class="hljs-string">package:</span> common
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">34</span> Checking <span class="hljs-string">file:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/common/</span>colorOutput.go
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">34</span> Checking <span class="hljs-string">file:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/common/</span>common.go
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">34</span> Import <span class="hljs-string">directory:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker/sysinfo
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">34</span> Checking <span class="hljs-string">package:</span> config
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">34</span> Checking <span class="hljs-string">file:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/config/</span>config.go
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">package:</span> docs
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">file:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/docs/</span>docs.go
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">package:</span> sysinfo
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">file:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/sysinfo/</span>sysinfo.go
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">package:</span> api
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">file:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.go
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">package:</span> main
[gosec] <span class="hljs-number">2022</span><span class="hljs-regexp">/12/</span><span class="hljs-number">05</span> <span class="hljs-number">09</span>:<span class="hljs-number">21</span>:<span class="hljs-number">35</span> Checking <span class="hljs-string">file:</span> <span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker/main.go
<span class="hljs-string">Results:</span>


[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/sysinfo/</span>sysinfo.<span class="hljs-string">go:</span><span class="hljs-number">38</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    37:</span>         } <span class="hljs-keyword">else</span> {
  &gt; <span class="hljs-number">38</span>:                 rsp.Body.Close()
<span class="hljs-symbol">    39:</span>                 fmt.Printf(<span class="hljs-string">"     %s Successfully established https connection to: %s\n"</span>, common.ConsoleSuccess(<span class="hljs-string">"[ ✓ SUCCESS ]"</span>), common.ConsoleBold(url))



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/common/</span>colorOutput.<span class="hljs-string">go:</span><span class="hljs-number">27</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    26:</span>         out.Write([]<span class="hljs-keyword">byte</span>(<span class="hljs-string">"\n\n"</span>))
  &gt; <span class="hljs-number">27</span>:         out.WriteTo(os.Stdout)
<span class="hljs-symbol">    28:</span> }



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.<span class="hljs-string">go:</span><span class="hljs-number">120</span><span class="hljs-number">-122</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    119:</span>        <span class="hljs-string">loadInfo :</span>= sysinfo.GetLoadInfo(c)
  &gt; <span class="hljs-number">120</span>:        c.Status(<span class="hljs-number">200</span>).JSON(&amp;fiber.Map{
  &gt; <span class="hljs-number">121</span>:                <span class="hljs-string">"loadInfo"</span>: loadInfo,
  &gt; <span class="hljs-number">122</span>:        })
<span class="hljs-symbol">    123:</span>        <span class="hljs-keyword">return</span> nil



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.<span class="hljs-string">go:</span><span class="hljs-number">104</span><span class="hljs-number">-106</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    103:</span>        <span class="hljs-string">netInfo :</span>= sysinfo.GetNetInfo(c)
  &gt; <span class="hljs-number">104</span>:        c.Status(<span class="hljs-number">200</span>).JSON(&amp;fiber.Map{
  &gt; <span class="hljs-number">105</span>:                <span class="hljs-string">"netInfo"</span>: netInfo,
  &gt; <span class="hljs-number">106</span>:        })
<span class="hljs-symbol">    107:</span>        <span class="hljs-keyword">return</span> nil



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.<span class="hljs-string">go:</span><span class="hljs-number">88</span><span class="hljs-number">-90</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    87:</span>         <span class="hljs-string">hostInfo :</span>= sysinfo.GetHostInfo(c)
  &gt; <span class="hljs-number">88</span>:         c.Status(<span class="hljs-number">200</span>).JSON(&amp;fiber.Map{
  &gt; <span class="hljs-number">89</span>:                 <span class="hljs-string">"hostInfo"</span>: hostInfo,
  &gt; <span class="hljs-number">90</span>:         })
<span class="hljs-symbol">    91:</span>         <span class="hljs-keyword">return</span> nil



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.<span class="hljs-string">go:</span><span class="hljs-number">72</span><span class="hljs-number">-74</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    71:</span>         <span class="hljs-string">cpuInfo :</span>= sysinfo.GetCPUInfo(c)
  &gt; <span class="hljs-number">72</span>:         c.Status(<span class="hljs-number">200</span>).JSON(&amp;fiber.Map{
  &gt; <span class="hljs-number">73</span>:                 <span class="hljs-string">"cpuInfo"</span>: cpuInfo,
  &gt; <span class="hljs-number">74</span>:         })
<span class="hljs-symbol">    75:</span>         <span class="hljs-keyword">return</span> nil



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.<span class="hljs-string">go:</span><span class="hljs-number">56</span><span class="hljs-number">-58</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    55:</span>         <span class="hljs-string">memInfo :</span>= sysinfo.GetMemInfo(c)
  &gt; <span class="hljs-number">56</span>:         c.Status(<span class="hljs-number">200</span>).JSON(&amp;fiber.Map{
  &gt; <span class="hljs-number">57</span>:                 <span class="hljs-string">"memInfo"</span>: memInfo,
  &gt; <span class="hljs-number">58</span>:         })
<span class="hljs-symbol">    59:</span>         <span class="hljs-keyword">return</span> nil



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.<span class="hljs-string">go:</span><span class="hljs-number">40</span><span class="hljs-number">-42</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    39:</span>         <span class="hljs-string">apiRoutes :</span>= sysinfo.GetAPIRoutes(c)
  &gt; <span class="hljs-number">40</span>:         c.Status(<span class="hljs-number">200</span>).JSON(&amp;fiber.Map{
  &gt; <span class="hljs-number">41</span>:                 <span class="hljs-string">"apiRoutes"</span>: apiRoutes,
  &gt; <span class="hljs-number">42</span>:         })
<span class="hljs-symbol">    43:</span>         <span class="hljs-keyword">return</span> nil



[<span class="hljs-regexp">/home/</span>matt<span class="hljs-regexp">/_apps/</span>golangdocker<span class="hljs-regexp">/api/</span>api.<span class="hljs-string">go:</span><span class="hljs-number">26</span>] - G104 (CWE<span class="hljs-number">-703</span>): Errors unhandled. (<span class="hljs-string">Confidence:</span> HIGH, <span class="hljs-string">Severity:</span> LOW)
<span class="hljs-symbol">    25:</span> func apiFalseRoot(c *fiber.Ctx) error {
  &gt; <span class="hljs-number">26</span>:         c.Redirect(<span class="hljs-string">"/api/v1"</span>)
<span class="hljs-symbol">    27:</span>         <span class="hljs-keyword">return</span> nil
<span class="hljs-symbol">


Summary:</span>
  <span class="hljs-string">Gosec  :</span> <span class="hljs-number">2.14</span><span class="hljs-number">.0</span>
  <span class="hljs-string">Files  :</span> <span class="hljs-number">7</span>
  <span class="hljs-string">Lines  :</span> <span class="hljs-number">635</span>
  <span class="hljs-string">Nosec  :</span> <span class="hljs-number">0</span>
  <span class="hljs-string">Issues :</span> <span class="hljs-number">9</span>
</code></pre>
      </div>
    </div>
  </div>

{% include footer.html %}