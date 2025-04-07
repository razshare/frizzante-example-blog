<script>
    import Submit from "$frizzante/components/Submit.svelte";
    import {getContext} from "svelte";

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
        <li><a href="#">About</a></li>
        <li><a href="#">Services</a></li>
        <li><a href="#">Products</a></li>
        <li>
            <Submit action="logout">
                <button>Logout</button>
            </Submit>
        </li>
    </ul>
</nav>