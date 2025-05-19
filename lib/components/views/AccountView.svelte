<script lang="ts">
    import Router from "$lib/components/Router.svelte";
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";

    type Data = {
        accountId: string
        displayName: string
    }

    type Props = {
        server: ServerProperties<Data>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<Layout title="Login">
    <Push top>
        <Menu bind:server/>
        <h1>{server.data.displayName}</h1>
        <small>{server.data.accountId}</small>
    </Push>
</Layout>