<style>
    article {
        max-width: 30rem;
    }

    button {
        width: 100%;
    }

    .LoginTitle, .AdditionalOptions {
        text-align: center;
    }

</style>

<script>
    import Layout from "$lib/components/Layout.svelte";
    import Form from "$frizzante/components/Form.svelte";
    import {getContext} from "svelte";
    import Center from "$lib/components/Center.svelte";

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
            <h1 class="LoginTitle">Login</h1>
            <Form>
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button type="submit">Continue</button>
            </Form>
            <p class="AdditionalOptions">
                or <a href={path("Register")}>register a new account</a>
                {#if data.error}
                    <br/>
                    <span class="pico-color-red-600">{data.error}</span>
                {/if}
            </p>
        </article>
    </Center>
</Layout>