<!DOCTYPE html>
<html>
<body>

  <h2>{{.note.Name}}</h2>
  <p>{{ .note.Content }}</p>

  <p>
    <a href="/notes/edit/{{.note.Id}}" role="button">Edit</a>
  </p>

</body>
</html>