<script lang="ts">
    import {getContext} from "svelte";
    import type {View} from "$frz/types.ts";

    type Article = {
        ID: string
        Title: string
        CreatedAt: number
        AccountId: string
    }

    const server = getContext("view") as View<Article[]>
</script>

<hr/>
<ul>
    {#each server.data as article(article.ID)}
        {@const createdAt = new Date(article.CreatedAt * 1000)}
        <li class="list-row">
            <div>
                <span>{article.Title}</span>
                <span>{createdAt}</span>
            </div>
        </li>
    {/each}
</ul>