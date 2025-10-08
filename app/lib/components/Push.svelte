<style>
    .left {
        justify-self: start;
    }

    .right {
        justify-self: end;
    }

    .top {
        align-self: start;
    }

    .bottom {
        align-self: end;
    }
</style>

<script lang="ts">
    import type { Snippet } from "svelte"

    type Props = {
        left?: boolean
        right?: boolean
        top?: boolean
        bottom?: boolean
        children: Snippet
    }

    let { left = false, right = false, top = false, bottom = false, children }: Props = $props()
</script>

<div class:left class:right class:top class:bottom>
    {@render children()}
</div>
