<style>
    .container {
        position: relative;
        display: grid;
        overflow-y: auto;
        padding: 1rem;
    }
</style>

<script lang="ts">
    import { fade } from "svelte/transition"
    import type { Snippet } from "svelte"

    type Props = {
        title: string
        children: Snippet
    }

    const { title, children }: Props = $props()
</script>

<svelte:head>
    <meta charset="UTF-8" />
    <meta
        name="viewport"
        content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0"
    />
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css" />
    <title>{title}</title>
</svelte:head>

<main class="container" transition:fade>
    {@render children()}
</main>
