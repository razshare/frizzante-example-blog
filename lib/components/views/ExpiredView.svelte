<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Center from "$lib/components/Center.svelte";
    import Link from "$lib/components/Link.svelte";
    import Router from "$lib/components/Router.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<Layout title="Session expired">
    <Center>
        <h1>Your session has expired.</h1>

        <Link bind:server to="Login">
            <button>Back to login.</button>
        </Link>
    </Center>
</Layout>