<script>
    import Layout from "$lib/components/Layout.svelte";
    import Link from "$frizzante/components/Link.svelte";
    import Center from "$lib/components/Center.svelte";
</script>

<Layout title="Session expired">
    <Center>
        <h1>Your session has expired.</h1>

        <Link view="Login">
            <button>Back to login.</button>
        </Link>
    </Center>
</Layout>