<style>
    li {
        padding: 1rem;
    }
</style>

<script lang="ts">
    import {href} from "$lib/utilities/scripts/href.ts";
    import {getContext} from "svelte";
    import type {View} from "$lib/utilities/types.ts";

    const view = getContext("view") as View<unknown>
</script>

<nav>
    <ul>
        <li><strong>Forum</strong></li>
    </ul>
    <ul>
        <li>
            {#if "Account" === view.name}
                <a {...href("/board")}>
                    <button>Board</button>
                </a>
            {:else}
                <a {...href("/account")}>
                    <button>Account</button>
                </a>
            {/if}
        </li>
        <li>
            <a {...href("/logout")}>
                <button>Logout</button>
            </a>
        </li>
    </ul>
</nav>