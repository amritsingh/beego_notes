{{ template "layouts/header.tpl" . }}

<div class="container mt-4">
    <div class="card">
        <div class="card-body">
            <h2 class="card-title">{{ .note.Name }}</h2>
            <p class="card-text">{{ .note.Content }}</p>
            <div class="btn-group" role="group" aria-label="Note Actions">
                <a href="/notes/edit/{{ .note.Id }}" class="btn btn-primary" role="button">Edit</a>
                <a class="btn btn-danger" role="button" onclick="deleteNote()">Delete</a>
            </div>
            <div class="btn-group" role="group" aria-label="Notes Nav">
                <a href="/notes/new" class="btn btn-primary" role="button">New Note</a>
                <a href="/notes" class="btn btn-info" role="button">Notes List</a>
            </div>
        </div>
    </div>
</div>

<script>
    function deleteNote() {
      var xhr = new XMLHttpRequest();
      xhr.open("DELETE", "/notes/{{.note.Id}}", true);
      xhr.onload = function() {
        if (xhr.status === 200) {
          // Redirect the page
          window.location.replace("/notes");
        }
      };
      xhr.send();
    }
</script>

{{ template "layouts/footer.tpl" . }}
