{{ template "layouts/header.tpl" . }}
<div class="container mt-4">
    <h2>Create a New Note</h2>

    <form action="/notes" method="POST">
        <div class="form-group">
            <label for="name">Title:</label>
            <input type="text" class="form-control" id="name" name="name">
        </div>

        <div class="form-group">
            <label for="content">Content:</label>
            <textarea class="form-control" id="content" name="content" rows="10"></textarea>
        </div>

        <div class="mb-2"></div>

        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
</div>
{{ template "layouts/footer.tpl" . }}
