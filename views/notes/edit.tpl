{{ template "layouts/header.tpl" . }}

<div class="container mt-4">
    <h1>Edit Note</h1>

    <form action="/notes/{{ .note.Id }}" method="POST">
        <div class="form-group">
            <label for="name">Title:</label>
            <input type="text" class="form-control" id="name" name="name" value="{{ .note.Name }}">
        </div>

        <div class="form-group">
            <label for="content">Content:</label>
            <textarea class="form-control" id="content" name="content" rows="10">{{ .note.Content }}</textarea>
        </div>

        <div class="mb-2"></div>

        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
</div>

{{ template "layouts/footer.tpl" . }}
