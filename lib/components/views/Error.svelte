<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types.ts";
    import Center from "$lib/components/Center.svelte";

    const server = getContext("server") as ServerContext<{ error: string }>
</script>

<Layout title="Error">
    <Center>
        {server.data.error}
    </Center>
</Layout>