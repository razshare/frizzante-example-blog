<script lang="ts">
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types";

    type Article = {
        articleId: string
        title: string
        createdAt: number
        accountId: string
    }

    let server = getContext("server") as ServerContext<{ articles: Article[] }>
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