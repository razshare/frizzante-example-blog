<script lang="ts">
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types";

    type Article = {
        ID: string
        Title: string
        CreatedAt: number
        AccountId: string
    }

    const server = getContext("server") as ServerContext<Article[]>
</script>

<hr/>
<ul>
    {#each server.data as article}
        {@const createdAt = new Date(article.CreatedAt * 1000)}
        <li class="list-row">
            <div>
                <span>{article.Title}</span>
                <span>{createdAt}</span>
            </div>
        </li>
    {/each}
</ul>