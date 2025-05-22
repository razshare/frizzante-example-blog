<style>
    article {
        max-width: 30rem;
    }

    .LoginTitle, .AdditionalOptions {
        text-align: center;
    }

    .error {
        color: #f45;
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
<Layout title="Register">
    <Center>
        <article>
            <h1 class="LoginTitle">Register</h1>
            <form {...action("Register")}>
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="text" name="displayName" placeholder="Display Name" aria-label="DisplayName">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button type="submit">Continue</button>
                <p class="AdditionalOptions">
                    or
                    <a {...href("Login")}>login</a>
                </p>
                {#if server.data.error}
                    <div class="error">{server.data.error}</div>
                {/if}
            </form>
        </article>
    </Center>
</Layout>