<script lang="ts">
    import Router from "$lib/components/Router.svelte";
    import Layout from "$lib/components/Layout.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<Layout title="">
    lorem

</Layout>