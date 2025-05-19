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
    import Form from "$lib/components/Form.svelte";
    import Link from "$lib/components/Link.svelte";
    import Router from "$lib/components/Router.svelte";

    type Props = {
        server: ServerProperties<{ error: string }>
    }

    let {server = $bindable()}: Props = $props()
</script>
<Router bind:server/>
<Layout bind:server title="Login">
    <Center>
        <article>
            <h1 class="LoginTitle">Login</h1>
            <Form bind:server of="Login">
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button type="submit">Continue</button>
            </Form>
            <p class="AdditionalOptions">
                or
                <Link bind:server to="Register">register a new account</Link>

                {#if server.data.error}
                    <br/>
                    <span class="pico-color-red-600">{server.data.error}</span>
                {/if}
            </p>
        </article>
    </Center>
</Layout>