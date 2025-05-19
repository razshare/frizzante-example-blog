<script lang="ts">
    import LoginView from "$lib/components/views/LoginView.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<LoginView bind:server/>
