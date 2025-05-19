<script lang="ts">

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<span>Hello, this is Redirect.svelte!</span>
<!-- Content. -->