<style>
    article {
        max-width: 30rem;
    }

    .login-title, .additional-options {
        text-align: center;
    }

</style>

<script>
    import Layout from "$lib/components/Layout.svelte";
    import {getContext} from "svelte";
    import Form from "$frizzante/components/Form.svelte";

    const path = getContext("Path")
</script>

<Layout title="Register">
    <article>
        <h1 class="login-title">Register</h1>
        <Form method="POST">
            <input type="email" name="Id" placeholder="Email" aria-label="Email">
            <input type="text" name="DisplayName" placeholder="Display Name" aria-label="DisplayName">
            <input type="password" name="Password" placeholder="Password" aria-label="Password">
            <button type="submit">Continue</button>
            <p class="additional-options">
                or <a href={path("Login")}>login</a>
            </p>
        </Form>
    </article>
</Layout>