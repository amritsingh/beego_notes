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

<a href="/notes/new">New Note</a>

</body>
</html>