<style>
    article {
        max-width: 30rem;
    }

    .LoginTitle,
    .AdditionalOptions {
        text-align: center;
    }
</style>

<script lang="ts">
    import { action } from "../../frizzante/scripts/action.ts"
    import { href } from "../../frizzante/scripts/href.ts"
    import Layout from "$lib/components/Layout.svelte"
    import Center from "$lib/components/Center.svelte"

    type Props = { error: string }
    let { error }: Props = $props()
</script>

<Layout title="Register">
    <Center>
        <article>
            <h1 class="LoginTitle">Register</h1>
            <form method="POST" {...action("/register")}>
                <input
                    type="email"
                    name="id"
                    placeholder="Email"
                    aria-label="Email"
                />
                <input
                    type="text"
                    name="displayName"
                    placeholder="Display Name"
                    aria-label="DisplayName"
                />
                <input
                    type="password"
                    name="password"
                    placeholder="Password"
                    aria-label="Password"
                />
                <button type="submit">Continue</button>
                <p class="AdditionalOptions">
                    or
                    <a {...href("/login")}>login</a>
                </p>
                {#if error}
                    <Center>
                        <div class="error">{error}</div>
                    </Center>
                {/if}
            </form>
        </article>
    </Center>
</Layout>
