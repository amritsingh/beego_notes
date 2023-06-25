<!DOCTYPE html>
<html>
<body>

  <script>
    function sendDelete(event, href){
      var xhttp = new XMLHttpRequest();
      event.preventDefault();
      xhttp.onreadystatechange = function() {
        // return if not ready state -> ~4
        if(this.readyState !== 4) {
            return;
        }

        if(this.readyState === 4) {
            // Redirect the page
            window.location.replace(this.responseURL);
        }
      };
      xhttp.open("DELETE", href, true);
      xhttp.send();
    }
  </script>

  <h2>{{.note.Name}}</h2>
  <p>{{ .note.Content }}</p>

  <p>
    <a href="/notes/edit/{{.note.Id}}" role="button">Edit</a>
    <a href="/notes/{{.note.Id}}" onclick="sendDelete(event, this.href)" role="button">Delete</a>
  </p>

</body>
</html>