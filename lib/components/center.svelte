<style>
    div {
        justify-self: center;
        align-self: center;
    }
</style>

<script>
    /**
     * @typedef Props
     * @property {import("svelte").Snippet<[]>} children
     */

    /** @type {Props} */
    let {children} = $props()
</script>

<div>
    {@render children()}
</div>