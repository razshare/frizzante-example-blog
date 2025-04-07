<script>
    import Icon from "$lib/components/Icon.svelte"
    import {getContext} from "svelte";
    import {mdiDotsVertical} from "@mdi/js";

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
<ul class="list bg-base-100 rounded-box shadow-md">
    {#each data.articles as article}
        {@const createdAt = new Date(article.createdAt * 1000)}
        <li class="list-row">
            <div>
                <div>{article.title}</div>
                <div class="text-xs uppercase font-semibold opacity-60">{createdAt}</div>
            </div>
        </li>
    {/each}
</ul>