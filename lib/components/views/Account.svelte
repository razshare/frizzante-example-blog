<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";
    import {getContext} from "svelte";
    import type {View} from "$frz/types.ts";

    type Account = {
        ID: string
        DisplayName: string
    }

    const view = getContext("view") as View<Account>
</script>

<Layout title="Login">
    <Push top>
        <Menu/>
        <h1>{view.data.DisplayName}</h1>
        <small>{view.data.ID}</small>
    </Push>
</Layout>