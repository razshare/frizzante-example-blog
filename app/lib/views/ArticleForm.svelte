<style>
    .button-wrapper {
        display: grid;
    }
    button {
        justify-self: end;
    }
</style>

<script lang="ts">
    import { action } from "$lib/utilities/scripts/action.ts"
    import Layout from "$lib/components/Layout.svelte";
</script>

<Layout title="New Article">
    <form method="POST" {...action()}>
        <span>Title</span><br />
        <input type="text" name="title" />
        <span>Content</span><br />
        <textarea name="content" cols="30" rows="10"></textarea>
        <div class="button-wrapper">
            <button>Publish</button>
        </div>
    </form>
</Layout>
