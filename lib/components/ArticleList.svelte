<script>
    import {getContext} from "svelte";

    /**
     * @typedef Article
     * @property {string} articleId
     * @property {string} title
     * @property {number} createdAt
     * @property {string} accountId
     */

    /**
     * @typedef Data
     * @property {Article[]} articles
     */

    /** @type {Data} */
    const data = getContext("data")
</script>

<hr/>
<ul>
    {#each data.articles as article}
        {@const createdAt = new Date(article.createdAt * 1000)}
        <li class="list-row">
            <div>
                <span>{article.title}</span>
                <span>{createdAt}</span>
            </div>
        </li>
    {/each}
</ul>