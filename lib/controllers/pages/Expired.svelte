<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Center from "$lib/components/Center.svelte";
    import Router from "$frizzante/components/Router.svelte";
    import {href} from "$frizzante/scripts/href.ts";
</script>

<Router/>
<Layout title="Session expired">
    <Center>
        <h1>Your session has expired.</h1>

        <a {...href("Login")}>
            <button>
                Back to login.
            </button>
        </a>
    </Center>
</Layout>