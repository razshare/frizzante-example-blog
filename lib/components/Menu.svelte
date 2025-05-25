<style>
    li {
        padding: 1rem;
    }
</style>

<script lang="ts">
    import {href} from "$frizzante/scripts/href.ts";
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types.ts";
    const server = getContext("server") as ServerContext<any>
</script>

<nav>
    <ul>
        <li><strong>Forum</strong></li>
    </ul>
    <ul>
        <li>
            {#if "Account" === server.view}
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