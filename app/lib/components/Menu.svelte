<style>
    li {
        padding: 1rem;
    }
</style>

<script lang="ts">
    import Icon from "$lib/components/icons/Icon.svelte"
    import { mdiAccountPlus, mdiLogout, mdiPlusBox } from "@mdi/js"
    import { href } from "$lib/scripts/core/href.ts"

    type Props = {
        loggedIn: boolean
        expired: boolean
    }

    let { loggedIn, expired }: Props = $props()
</script>

<nav>
    <ul>
        <li></li>
    </ul>
    <ul>
        {#if loggedIn && !expired}
            <li>
                <a {...href("/logout")}>
                    <span>Logout</span>
                    <Icon path={mdiLogout} />
                </a>
            </li>
        {:else}
            <li>
                <a {...href("/register")}>
                    <span>Register</span>
                    <Icon path={mdiAccountPlus} />
                </a>
            </li>
            <li>
                <a {...href("/login")}>
                    <span>Login</span>
                    <Icon path={mdiLogout} />
                </a>
            </li>
        {/if}

        {#if loggedIn && !expired}
            <li>
                <a {...href("/form")}>
                    <span>New Article</span>
                    <Icon path={mdiPlusBox} />
                </a>
            </li>
        {/if}
    </ul>
</nav>
