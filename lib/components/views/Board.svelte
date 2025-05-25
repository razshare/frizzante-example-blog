<script lang="ts">
    import ArticleList from "$lib/components/ArticleList.svelte";
    import Layout from "$lib/components/Layout.svelte";
    import Menu from "$lib/components/Menu.svelte";
    import Push from "$lib/components/Push.svelte";
    import ArticleForm from "$lib/components/ArticleForm.svelte";
</script>

<Layout title="Login">
    <Push top>
        <Menu/>
        <ArticleForm/>
        <ArticleList/>
    </Push>
</Layout>