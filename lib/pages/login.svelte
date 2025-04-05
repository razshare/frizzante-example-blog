<style>
    article {
        max-width: 30rem;
    }

    button {
        width: 100%;
    }

    .login_title, .additional_options {
        text-align: center;
    }

</style>

<script>
    import Layout from "$lib/components/layout.svelte";
    import {getContext} from "svelte";
    import Form from "$frizzante/components/form.svelte";
    import Center from "$lib/components/center.svelte";

    /**
     * @typedef Data
     * @property {string} error
     * @property {string} message
     * @property {boolean} verified
     */

    /** @type {Data} */
    const data = getContext("data")
    const path = getContext("path")
</script>

<Layout title="Login">
    <Center>
        <article>
            <h1 class="login_title">Login</h1>
            <Form>
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button type="submit">Continue</button>
            </Form>
            <p class="additional_options">
                or <a href={path("register")}>register a new account</a>
                {#if data.error}
                    <br/>
                    <span class="pico-color-red-600">{data.error}</span>
                {/if}
            </p>
        </article>
    </Center>
</Layout>