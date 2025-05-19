<script lang="ts">
    import {href} from "$frizzante/scripts/href.ts";
    import {action} from "$frizzante/scripts/action.ts";
</script>

<nav>
    <ul>
        <li><strong>Forum</strong></li>
    </ul>
    <ul>
        <li>
            <a {...href("Account")}>
                <button>Account</button>
            </a>
        </li>
        <li>
            <form {...action("Logout")}>
                <button>Logout</button>
            </form>
        </li>
    </ul>
</nav>