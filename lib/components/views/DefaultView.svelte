<script lang="ts">
    import LoginView from "$lib/components/views/LoginView.svelte";
    import Router from "$lib/components/Router.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<LoginView bind:server/>