<style>
    li {
        padding: 1rem;
    }
</style>

<script lang="ts">
    import { href } from "$frizzante/core/scripts/href.ts"
    import Icon from "$lib/components/Icon.svelte"
    import {mdiAccountPlus, mdiLogout, mdiPlusBox} from "@mdi/js"

    type Props = {
        verified: boolean
        expired: boolean
    }

    let { verified, expired }: Props = $props()
</script>

<nav>
    <ul>
        <li></li>
    </ul>
    <ul>
        {#if verified && !expired}
            <li>
                <a {...href("/logout")}>
                    <span>Logout</span>
                    <Icon value={mdiLogout} />
                </a>
            </li>
        {:else}
            <li>
                <a {...href("/register")}>
                    <span>Register</span>
                    <Icon value={mdiAccountPlus} />
                </a>
            </li>
            <li>
                <a {...href("/login")}>
                    <span>Login</span>
                    <Icon value={mdiLogout} />
                </a>
            </li>
        {/if}

        {#if verified && !expired}
            <li>
                <a {...href("/article-form")}>
                    <span>New Article</span>
                    <Icon value={mdiPlusBox} />
                </a>
            </li>
        {/if}
    </ul>
</nav>
