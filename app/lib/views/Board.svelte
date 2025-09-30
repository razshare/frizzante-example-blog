<script lang="ts">
    import ArticleList from "$lib/components/ArticleList.svelte"
    import Layout from "$lib/components/Layout.svelte"
    import Menu from "$lib/components/Menu.svelte"
    import Push from "$lib/components/Push.svelte"
    import type { Article } from "$lib/types.ts"

    type Props = {
        page: number
        hasMore: boolean
        articles: Article[]
        loggedIn: boolean
        expired: boolean
    }

    let { page, hasMore, articles, loggedIn, expired }: Props = $props()
</script>

<Layout title="Board">
    <Push top>
        <Menu {loggedIn} {expired} />
        <ArticleList {page} {hasMore} {articles} {loggedIn} />
    </Push>
</Layout>
