<script lang="ts">
    import Router from "$lib/components/Router.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<span>Hello, this is Account.svelte!</span>
<!-- Content. -->