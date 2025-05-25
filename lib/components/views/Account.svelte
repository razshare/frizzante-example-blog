<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types.ts";

    type Account = {
        ID: string
        DisplayName: string
    }

    const server = getContext("server") as ServerContext<Account>
</script>

<Layout title="Login">
    <Push top>
        <Menu/>
        <h1>{server.data.DisplayName}</h1>
        <small>{server.data.ID}</small>
    </Push>
</Layout>