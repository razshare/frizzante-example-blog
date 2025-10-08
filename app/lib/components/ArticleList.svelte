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
            "title created-at remove"
            "content content content";
        grid-template-columns: 1fr auto auto;
        gap: 1rem;
    }
    .title {
        grid-area: title;
    }
    .created-at {
        grid-area: created-at;
    }
    .remove {
        grid-area: remove;
    }
    .content {
        grid-area: content;
    }
    .no-articles {
        text-align: center;
    }
</style>

<script lang="ts">
    import Icon from "$lib/components/icons/Icon.svelte"
    import { mdiArrowLeft, mdiArrowRight } from "@mdi/js"
    import { href } from "$lib/scripts/core/href.ts"
    import type {sqlc} from "$gen/types/main/lib/routes/board/Props";
    type Props = {
        page: number
        hasMore: boolean
        articles: sqlc.Article[]
        loggedIn: boolean
    }
    let { page, hasMore, articles, loggedIn }: Props = $props()
</script>

{#if articles.length === 0}
    <p class="no-articles">No article available.</p>
{:else}
    {#snippet navigator()}
        <div class="items">
            {#if page > 0}
                <a class="previous" {...href(`/board?page=${page - 1}`)}>
                    <Icon path={mdiArrowLeft} />
                    <span>Previous</span>
                </a>
            {:else}
                <span class="previous">
                    <Icon path={mdiArrowLeft} />
                    <span>Previous</span>
                </span>
            {/if}

            {#if hasMore}
                <a class="next" {...href(`/board?page=${page + 1}`)}>
                    <span>Next</span>
                    <Icon path={mdiArrowRight} />
                </a>
            {:else}
                <span class="next">
                    <span>Next</span>
                    <Icon path={mdiArrowRight} />
                </span>
            {/if}
        </div>
    {/snippet}

    {@render navigator()}

    {#each articles as article (article.id)}
        {@const createdAt = new Date(article.createdAt * 1000).toLocaleString()}
        <hr />
        <div class="article">
            <h1 class="title">{article.title}</h1>
            <span class="created-at">{createdAt}</span>
            {#if loggedIn}
                <span class="remove">
                    <a {...href(`/article/remove?id=${article.id}`)}>[Remove]</a>
                </span>
            {/if}
            <div class="content">{article.content}</div>
        </div>
    {/each}

    {@render navigator()}
{/if}
