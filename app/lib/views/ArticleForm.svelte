<style>
    .button-wrapper {
        display: grid;
    }
    .publish {
        justify-self: center;
    }
    .back {
        justify-self: end;
    }
    .details {
        text-align: center;
        padding: 1rem;
    }
</style>

<script lang="ts">
    import { action } from "$lib/utilities/frz/scripts/action.ts"
    import Layout from "$lib/components/Layout.svelte"
    import { href } from "$lib/utilities/frz/scripts/href.ts"
    import Icon from "$lib/components/Icon.svelte"
    import { mdiArrowLeft, mdiPlus } from "@mdi/js"

    type Props = { error: string }
    let { error }: Props = $props()
</script>

<Layout title="Article Form">
    <div class="button-wrapper">
        <a class="back" {...href("/board")}>
            <Icon value={mdiArrowLeft} />
            <span>Back</span>
        </a>
    </div>
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
            <button class="publish">
                <span>Publish</span>
                <Icon value={mdiPlus} />
            </button>
        </div>
    </form>
</Layout>
