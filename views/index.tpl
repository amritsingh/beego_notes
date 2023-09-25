{{ define "index.tpl"}}
    {{ template "layouts/header.tpl" .}}
      <div class="container mt-4">
        <h1>{{.Title}}</h1>

        <h2 class="mb-4"><a href="/notes" class="mb-4">All Note</a></h2>

        <p>
            <div class="btn-group" role="group">
                {{if .LoggedIn}}
                    <form action="/logout" method="POST">
                        <button type="submit" class="btn btn-outline-danger">Logout</button>
                    </form>
                {{else}}
                    <a class="btn btn-outline-primary" href="/login" role="button">Login</a>
                    <a class="btn btn-outline-primary" href="/signup" role="button">Sign Up</a>
                {{end}}
            </div>
        </p>
      </div>
    {{ template "layouts/footer.tpl" .}}
{{ end }}