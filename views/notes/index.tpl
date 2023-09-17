{{ template "layouts/header.tpl" . }}

<div class="container mt-4">
    <h2 class="mb-4">All Notes</h2>

    <div class="row">
        {{ range $index, $note := .notes }}
        <div class="col-md-4 mb-4">
            <div class="card">
                <div class="card-body">
                    <h5 class="card-title"><a href="/notes/{{ $note.Id }}">{{ $note.Name }}</a></h5>
                    <p class="card-text">{{ $note.Content }}</p>
                </div>
            </div>
        </div>
        {{ end }}
    </div>

    <a href="/notes/new" class="btn btn-primary">New Note</a>
</div>

{{ template "layouts/footer.tpl" . }}
