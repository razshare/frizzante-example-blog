<style>
    article {
        max-width: 30rem;
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

    const path = getContext("path")
</script>

<Layout title="Register">
    <Center>
        <article>
            <h1 class="login_title">Register</h1>
            <Form>
                <input type="email" name="id" placeholder="Email" aria-label="Email">
                <input type="text" name="display_name" placeholder="Display Name" aria-label="DisplayName">
                <input type="password" name="password" placeholder="Password" aria-label="Password">
                <button type="submit">Continue</button>
                <p class="additional_options">
                    or <a href={path("login")}>login</a>
                </p>
            </Form>
        </article>
    </Center>
</Layout>