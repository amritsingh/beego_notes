<!DOCTYPE html>
<html>
<body>

  <h2>{{.note.Name}}</h2>
  <p>{{ .note.Content }}</p>

  <p>
    <a href="/notes/edit/{{.note.Id}}" role="button">Edit</a>
    <a href="#content" class="skip-link" onclick="deleteNote()">Delete</a>
  </p>

  <p>
    <a href="/notes/new">New Note</a>
    <a href="/notes">Notes List</a>
  </p>

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

</body>
</html>