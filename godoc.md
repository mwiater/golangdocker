---
layout: null
---
{% include header.html %}

  <div class="container">
    <div class="row">
      <div class="col">
        <h1 id="godoc">Godoc</h1>
        <p>[IN PROGRESS]</p>
        <p><code>godoc -http=:6060</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </p>
        <p>Or via Makefile:
          <code>make golang-godoc</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> (which will execute the <a target="_blank" href="https://github.com/mwiater/golangdocker/blob/master/scripts/golang_godoc.sh">scripts/golang_godoc.sh</a> script)
        </p>

        <p>Access via browser at: <code>http://{your-ip-address}:6060/pkg/{app-module-name-in-go.mod}</code></p>
        <p>E.g.: <code>http://192.168.0.91:6060/pkg/github.com/mattwiater/golangdocker/</code></p>
      </div>
    </div>
  </div>

{% include footer.html %}