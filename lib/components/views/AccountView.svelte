<script lang="ts">
    import Router from "$lib/components/Router.svelte";
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<Layout title="Login">
    <Push top>
        <Menu bind:server/>
        <h1>This is my account.</h1>
    </Push>
</Layout>