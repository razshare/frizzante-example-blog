<script lang="ts">
    import Layout from "$lib/components/Layout.svelte"
    import Center from "$lib/components/Center.svelte"
    import {href} from "$lib/scripts/core/href.ts";
</script>

<Layout title="Session Expired">
    <Center>
        <h1>Your session has expired.</h1>

        <a {...href("/login")}>
            <button> Back to login. </button>
        </a>
    </Center>
</Layout>
