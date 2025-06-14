<style>
    .items {
        display: flex;
        gap: 1rem;
    }
    .previous {
        flex-grow: 1;
        text-align: start;
    }
    .next {
        flex-grow: 1;
        text-align: end;
    }
    .article {
        padding: 1rem;
        display: grid;
        grid-template-areas:
        "title created-at"
        "content content";
        grid-template-columns: 1fr auto;
    }
    .title {
        grid-area: title;
    }
    .created-at {
        grid-area: created-at;
    }
    .content {
        grid-area: content;
    }
    .no-articles {
        text-align: center;
    }
</style>

<script lang="ts">
    import {href} from "$lib/utilities/scripts/href.ts";
    import type {Article} from "$lib/types.ts";

    type Props = {
        page:number,
        hasMore:boolean,
        articles:Article[],
    }

    let {page, articles, hasMore}:Props = $props()
</script>

{#if articles.length === 0}
    <p class="no-articles">No article available.</p>
{:else}
    {#snippet navigator()}
        <div class="items">
            <button disabled={page <= 0} class="previous" {...href(`/board?page=${page-1}`)}>&lt; Previous</button>
            <button disabled={!hasMore} class="next" {...href(`/board?page=${page+1}`)}>Next &gt;</button>
        </div>
    {/snippet}

    {@render navigator()}

    {#each articles as article (article.ID)}
        {@const createdAt = new Date(article.CreatedAt * 1000).toLocaleString()}
        <hr />
        <div class="article">
            <h1 class="title">{article.Title}</h1>
            <span class="created-at">{createdAt}</span>
            <div class="content">{article.Content}</div>
        </div>
    {/each}

    {@render navigator()}
{/if}