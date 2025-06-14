<style>
    .button-wrapper {
        display: grid;
    }
    button {
        justify-self: end;
    }
    .details {
        text-align: center;
        padding: 1rem;
    }
</style>

<script lang="ts">
    import { action } from "$lib/utilities/scripts/action.ts"
    import Layout from "$lib/components/Layout.svelte";

    type Props = {
        error: string
    }

    let {error}:Props = $props()
</script>

<Layout title="New Article">
    <form method="POST" {...action()}>
        <span>Title</span><br />
        <input type="text" name="title" />
        <span>Content</span><br />
        <textarea name="content" cols="30" rows="10"></textarea>
        {#if error}
            <div class="details">
                <div class="error">{error}</div>
            </div>
        {/if}
        <div class="button-wrapper">
            <button>Publish</button>
        </div>
    </form>
</Layout>
