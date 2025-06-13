<script>
    import {action} from "$lib/utilities/scripts/action.ts";
</script>

<form method="POST" {...action("/board")}>
    <span>Title</span><br/>
    <input type="text" name="title">
    <span>Content</span><br/>
    <textarea name="content" cols="30" rows="10"></textarea>
    <button>Publish</button>
</form>