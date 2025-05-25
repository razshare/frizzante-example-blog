<script>
    import {action} from "$frizzante/scripts/action.ts";
</script>

<form {...action("/article")}>
    <span>Title</span><br/>
    <input type="text" name="title">
    <span>Content</span><br/>
    <textarea name="content" cols="30" rows="10"></textarea>
    <button>Publish</button>
</form>