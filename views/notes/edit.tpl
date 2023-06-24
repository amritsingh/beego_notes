<!DOCTYPE html>
<html>
<body>

  <h1>Edit Note</h1>
  
  <form action="/notes/{{.note.Id}}" method="POST">
    <label for="name">Title:</label><br>
    <input type="text" id="name" name="name" value="{{.note.Name}}"><br>
    <br>
    <label for="content">Content:</label><br>
    <textarea name="content" id="content" rows="10" cols="30">{{.note.Content}}</textarea><br>
    <br>
    <input type="submit" value="Submit">
  </form>

</body>
</html>