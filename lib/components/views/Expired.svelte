<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Center from "$lib/components/Center.svelte";
    import {href} from "$frizzante/scripts/href.ts";
</script>

<Layout title="Session expired">
    <Center>
        <h1>Your session has expired.</h1>

        <a {...href("login")}>
            <button>
                Back to login.
            </button>
        </a>
    </Center>
</Layout>