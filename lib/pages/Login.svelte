<style>
    article {
        max-width: 30rem;
    }

    button {
        width: 100%;
    }

    .login-title, .additional-options {
        text-align: center;
    }

</style>

<script>
    import Layout from "$lib/components/Layout.svelte";
    import {getContext} from "svelte";
    import Submit from "$frizzante/components/Submit.svelte";
    import Form from "$frizzante/components/Form.svelte";

    /**
     * @typedef Data
     * @property {string} Error
     * @property {string} Message
     * @property {boolean} Logged
     */

    /** @type {Data} */
    const data = getContext("Data")
    const path = getContext("Path")
</script>

<Layout title="Login">
    <article>
        <h1 class="login-title">Login</h1>
        <Form method="POST">
            <input type="email" name="Id" placeholder="Email" aria-label="Email">
            <input type="password" name="Password" placeholder="Password" aria-label="Password">
            <button type="submit">Continue</button>
        </Form>
        <p class="additional-options">
            or <a href={path("Register")}>register a new account</a>
            {#if data.Error}
                <br/>
                <span class="pico-color-red-600">{data.Error}</span>
            {:else if data.Message}
                <br/>
                <span class="pico-color-blue-600">{data.Message}</span>
                {#if data.Logged}
                    <Submit form="{{Logout:true}}">
                        <button>Logout</button>
                    </Submit>
                {/if}
            {/if}
        </p>
    </article>
</Layout>