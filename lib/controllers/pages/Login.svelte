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

<script lang="ts">
    import Layout from "$lib/components/Layout.svelte";
    import Center from "$lib/components/Center.svelte";
    import Router from "$frizzante/components/Router.svelte";
    import {getContext} from "svelte";
    import type {ServerContext} from "$frizzante/types.ts";
    import {action} from "$frizzante/scripts/action.ts";
    import {href} from "$frizzante/scripts/href.ts";

    const server = getContext("server") as ServerContext<{ error: string }>
</script>
<Router/>
<Layout title="Login">
    <Center>
        <article>
            <h1 class="LoginTitle">Login</h1>
            <form {...action("Login")}>
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button>Continue</button>
            </form>
            <br/>
            <p class="AdditionalOptions">
                or

                <a {...href("Register")}>register a new account</a>
                {#if server.data.error}
                    <br/>
                    <span class="pico-color-red-600">{server.data.error}</span>
                {/if}
            </p>
        </article>
    </Center>
</Layout>