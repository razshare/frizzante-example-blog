<script lang="ts">
    import ArticleList from "$lib/components/ArticleList.svelte"
    import Layout from "$lib/components/Layout.svelte"
    import Menu from "$lib/components/Menu.svelte"
    import Push from "$lib/components/Push.svelte"
    import type {Article} from "$lib/types.ts";
    
    let {page, hasMore, articles} = $props() as {
        page: number,
        hasMore: boolean,
        articles: Article[],
    }
</script>

<Layout title="Login">
    <Push top>
        <Menu />
        <br/>
        <br />
        <ArticleList {page} {hasMore} {articles} />
    </Push>
</Layout>
