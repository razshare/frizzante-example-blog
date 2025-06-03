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
    import {action} from "$frizzante/scripts/action.ts";
    import {href} from "$frizzante/scripts/href.ts";
    import {getContext} from "svelte";
    import type {View} from "$frizzante/types.ts";

    const view = getContext("view") as View<unknown>
</script>

<Layout title="Login">
    <Center>
        <article>
            <h1 class="LoginTitle">Login</h1>
            <form {...action("/login")}>
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button>Continue</button>
            </form>
            <br/>
            <p class="AdditionalOptions">
                or <a {...href("/register")}>register a new account</a>
            </p>
            {#if view.error}
                <Center>
                    <div class="error">{view.error}</div>
                </Center>
            {/if}
        </article>
    </Center>
</Layout>