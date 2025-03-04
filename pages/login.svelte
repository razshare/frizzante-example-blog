<style>
    article {
        max-width: 30rem;
    }

    .login-title, .additional-options {
        text-align: center;
    }

</style>

<script>
    import Layout from "$lib/components/Layout.svelte";
    import {getContext} from "svelte";
    import {update} from "$frizzante/scripts/update.js";
    import Submit from "$frizzante/components/Submit.svelte";

    /**
     * @typedef Data
     * @property {string} error
     * @property {string} message
     * @property {boolean} logged
     */

    /** @type {Data} */
    const data = getContext("data")
    const path = getContext("path")
</script>

<Layout title="Login">
    <article>
        <h1 class="login-title">Login</h1>
        <form method="POST" action="?" onsubmit={update(getContext("data"))}>
            <input type="text" name="id" placeholder="Id" aria-label="Id">
            <input type="password" name="password" placeholder="Password" aria-label="Password">
            <button type="submit">Continue</button>
        </form>
        <p class="additional-options">
            or <a href={path("register")}>register a new account</a>
            {#if data.error}
                <br/>
                <span class="pico-color-red-600">{data.error}</span>
            {:else if data.message}
                <br/>
                <span class="pico-color-blue-600">{data.message}</span>
                {#if data.logged}
                    <Submit form="{{logout:true}}">
                        <button>Logout</button>
                    </Submit>
                {/if}
            {/if}
        </p>
    </article>
</Layout>