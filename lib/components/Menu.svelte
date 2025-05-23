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
            <a {...href("account")}>
                <button>Account</button>
            </a>
        </li>
        <li>
            <form {...action("logout")}>
                <button>Logout</button>
            </form>
        </li>
    </ul>
</nav>