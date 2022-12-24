---
layout: null
---
{% include header.html %}

  <div class="container">
    <div class="row">
      <div class="col">
        <h1 id="linting">Linting</h1>
        <p>[IN PROGRESS]</p>
        <h2 id="to-do">To Do</h2>
        <ul>
          <li>Create config file <a href="https://golangci-lint.run/usage/configuration/">Official Docs</a></li>
        </ul>
        <h2 id="golangci-lint">golangci-lint</h2>
        <p><code>go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </p>
        <p>Usage: <code>golangci-lint run</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> </p>
        <p>Or via Makefile:
          <code>make golang-lint</code> <i class="fa-duotone fa-copy fa-fw code-copy-button"></i> (which will execute the <a target="_blank" href="https://github.com/mwiater/golangdocker/blob/master/scripts/golang_lint.sh">scripts/golang_lint.sh</a> script)
        </p>
      </div>
    </div>
  </div>

{% include footer.html %}