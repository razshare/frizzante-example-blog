<style global>
    .container {
        position: fixed;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        display: grid;
    }
</style>

<script lang="ts">
    import type {Snippet} from "svelte";

    type Props = {
        title: string
        children: Snippet
        server: ServerProperties<{}>
    }

    const {
        title,
        children,
        server = $bindable()
    }: Props = $props()
</script>


<svelte:head>
    <meta charset="UTF-8">
    <meta name="viewport"
          content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0">
    <title>{title}</title>
    <link
            rel="stylesheet"
            href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.min.css"
    >
    <link
            rel="stylesheet"
            href="https://cdn.jsdelivr.net/npm/@picocss/pico@2/css/pico.colors.min.css"
    >
</svelte:head>

<main class="container">
    {@render children()}
</main>