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

<script>
    /**
     * @typedef Props
     * @property {import("svelte").Snippet<[]>} children
     * @property {boolean} [left]
     * @property {boolean} [right]
     * @property {boolean} [top]
     * @property {boolean} [bottom]
     */

    /** @type {Props} */
    let {
        children,
        left = false,
        right = false,
        top = false,
        bottom = false,
    } = $props()
</script>

<div
        class:left
        class:right
        class:top
        class:bottom
>
    {@render children()}
</div>