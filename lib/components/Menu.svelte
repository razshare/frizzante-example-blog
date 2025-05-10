<script>
    import Submit from "$frizzante/components/Submit.svelte";
    import {getContext} from "svelte";
    import Link from "$frizzante/components/Link.svelte";

    /**
     * @typedef Data
     * @property {boolean} Verified
     */

    /** @type {Data} */
    const data = getContext("data")
</script>

<nav>
    <ul>
        <li><strong>Forum</strong></li>
    </ul>
    <ul>
        <li>
            <button>
                <Link view="Account">Account</Link>
            </button>
        </li>
        <li>
            <Submit view="Logout">
                <button>Logout</button>
            </Submit>
        </li>
    </ul>
</nav>