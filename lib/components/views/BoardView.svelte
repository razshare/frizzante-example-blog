<script lang="ts">
    import ArticleList from "$lib/components/ArticleList.svelte";
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";
    import Router from "$lib/components/Router.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<Layout bind:server title="Login">
    <Push top>
        <Menu bind:server/>
        <ArticleList bind:server/>
    </Push>
</Layout>