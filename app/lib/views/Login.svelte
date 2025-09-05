<style>
    article {
        max-width: 30rem;
    }

    button {
        width: 100%;
    }

    .LoginTitle,
    .AdditionalOptions {
        text-align: center;
    }
</style>

<script lang="ts">
    import Layout from "$lib/components/Layout.svelte"
    import Center from "$lib/components/Center.svelte"
    import {action} from "$lib/scripts/core/action.ts";
    import {href} from "$lib/scripts/core/href.ts";

    type Props = { error: string }
    let { error }: Props = $props()
</script>

<Layout title="Login">
    <Center>
        <article>
            <h1 class="LoginTitle">Login</h1>
            <form method="POST" {...action("/login")}>
                <input
                    type="email"
                    name="id"
                    placeholder="Email"
                    aria-label="Email"
                />
                <input
                    type="password"
                    name="password"
                    placeholder="Password"
                    aria-label="Password"
                />
                <button>Continue</button>
            </form>
            <br />
            <p class="AdditionalOptions">
                or <a {...href("/register")}>register a new account</a>
                , <a {...href("/board")}>go back to board</a>
            </p>
            {#if error}
                <Center>
                    <div class="error">{error}</div>
                </Center>
            {/if}
        </article>
    </Center>
</Layout>
