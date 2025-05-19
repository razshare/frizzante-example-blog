<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types.ts";
    import Router from "$frizzante/components/Router.svelte";

    let server = getContext("server") as ServerContext<{
        accountId: string
        displayName: string
    }>
</script>

<Router/>
<Layout title="Login">
    <Push top>
        <Menu/>
        <h1>{server.data.displayName}</h1>
        <small>{server.data.accountId}</small>
    </Push>
</Layout>