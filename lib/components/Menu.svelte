<script lang="ts">
    import Link from "$lib/components/Link.svelte";
    import Action from "$lib/components/Action.svelte";

    type Props = {
        server: ServerProperties<{}>
    }

    let {
        server = $bindable(),
    }: Props = $props()
</script>

<nav>
    <ul>
        <li><strong>Forum</strong></li>
    </ul>
    <ul>
        <li>
<!--            <button>-->
                <Link bind:server to="Account">Account</Link>
<!--            </button>-->
        </li>
        <li>
            <Action bind:server of="Logout">
                <button>Logout</button>
            </Action>
        </li>
    </ul>
</nav>