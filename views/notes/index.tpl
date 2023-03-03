<!DOCTYPE html>
<html>
<body>

<h2>All notes</h2>

<dl>
  {{ range $index, $note := .notes }}
    <dt>{{ $note.Name }}</dt>
    <dd>- {{ $note.Content }}</dd>
  {{ end }}
</dl>

</body>
</html>