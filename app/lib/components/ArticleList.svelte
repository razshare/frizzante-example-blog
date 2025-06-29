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
    import { href } from "../../frizzante/scripts/href.ts"
    import type { Article } from "$lib/types.ts"
    import Icon from "$lib/components/Icon.svelte"
    import { mdiArrowLeft, mdiArrowRight } from "@mdi/js"

    type Props = {
        page: number
        hasMore: boolean
        articles: Article[]
    }

    let { page, articles, hasMore }: Props = $props()
</script>

{#if articles.length === 0}
    <p class="no-articles">No article available.</p>
{:else}
    {#snippet navigator()}
        <div class="items">
            {#if page > 0}
                <a class="previous" {...href(`/board?page=${page - 1}`)}>
                    <Icon value={mdiArrowLeft} />
                    <span>Previous</span>
                </a>
            {:else}
                <span class="previous">
                    <Icon value={mdiArrowLeft} />
                    <span>Previous</span>
                </span>
            {/if}

            {#if hasMore}
                <a class="next" {...href(`/board?page=${page + 1}`)}>
                    <span>Next</span>
                    <Icon value={mdiArrowRight} />
                </a>
            {:else}
                <span class="next">
                    <span>Next</span>
                    <Icon value={mdiArrowRight} />
                </span>
            {/if}
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
