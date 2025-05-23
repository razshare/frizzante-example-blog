<script lang="ts">
    import Router from "$frizzante/components/Router.svelte";
    import ArticleList from "$lib/components/ArticleList.svelte";
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";
</script>

<Router/>
<Layout title="Login">
    <Push top>
        <Menu/>
        <ArticleList/>
    </Push>
</Layout>