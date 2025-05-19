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
    import Form from "$lib/components/Form.svelte";
    import Link from "$lib/components/Link.svelte";
    import Router from "$lib/components/Router.svelte";

    type Props = {
        server: ServerProperties<any>
    }

    let {server = $bindable()}: Props = $props()
</script>

<Router bind:server/>
<Layout title="Register">
    <Center>
        <article>
            <h1 class="LoginTitle">Register</h1>
            <Form bind:server of="Register">
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="text" name="displayName" placeholder="Display Name" aria-label="DisplayName">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button type="submit">Continue</button>
                <p class="AdditionalOptions">
                    or
                    <Link bind:server to="Login">login</Link>
                </p>
                {#if server.data.error}
                    <div class="error">{server.data.error}</div>
                {/if}
            </Form>
        </article>
    </Center>
</Layout>