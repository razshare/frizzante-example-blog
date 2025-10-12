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
    import Layout from "$lib/components/Layout.svelte"
    import Icon from "$lib/components/icons/Icon.svelte"
    import { mdiArrowLeft, mdiPlus } from "@mdi/js"
    import { href } from "$lib/scripts/core/href.ts"
    import { action } from "$lib/scripts/core/action.ts"
    import type { Props } from "$gen/types/main/lib/routes/form/Props"
    let { error }: Props = $props()
</script>

<Layout title="Article Form">
    <div class="button-wrapper">
        <a class="back" {...href("/board")}>
            <Icon path={mdiArrowLeft} />
            <span>Back</span>
        </a>
    </div>
    <form method="POST" {...action("/form/add")}>
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
                <Icon path={mdiPlus} />
            </button>
        </div>
    </form>
</Layout>
