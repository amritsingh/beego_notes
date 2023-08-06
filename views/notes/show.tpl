<!DOCTYPE html>
<html>
<body>

  <h2>{{.note.Name}}</h2>
  <p>{{ .note.Content }}</p>

  <p>
    <a href="/notes/edit/{{.note.Id}}" role="button">Edit</a>
    <a href="/notes/{{.note.Id}}" onclick="deleteNote(this)">Delete</a>
  </p>

  <script>
    function deleteNote(link) {
      var xhr = new XMLHttpRequest();
      xhr.open("DELETE", link.href, true);
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