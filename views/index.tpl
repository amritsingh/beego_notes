{{ define "index.tpl"}}
    {{ template "layouts/header.tpl" .}}
    <h1>{{.Title}}</h1>
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
    {{ template "layouts/footer.tpl" .}}
{{ end }}