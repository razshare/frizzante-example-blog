<script lang="ts">
    type Article = {
        articleId: string
        title: string
        createdAt: number
        accountId: string
    }

    type Props = {
        server: ServerProperties<{ articles: Article[] }>
    }

    let {server = $bindable()}: Props = $props()
</script>

<hr/>
<ul>
    {#each server.data.articles as article}
        {@const createdAt = new Date(article.createdAt * 1000)}
        <li class="list-row">
            <div>
                <span>{article.title}</span>
                <span>{createdAt}</span>
            </div>
        </li>
    {/each}
</ul>