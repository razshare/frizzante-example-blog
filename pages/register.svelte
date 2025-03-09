<style>
    article {
        max-width: 30rem;
    }
    .login-title, .additional-options{
        text-align: center;
    }

</style>

<script>
    import Layout from "$lib/components/Layout.svelte";
    import {getContext} from "svelte";
    import {update} from "$frizzante/scripts/update.js";
    const path = getContext("path")
</script>

<Layout title="Register">
    <article>
        <h1 class="login-title">Register</h1>
        <form method="POST" action="?" onsubmit={update(getContext("data"))}>
            <input type="text" name="id" placeholder="Id" aria-label="id">
            <input type="password" name="password" placeholder="Password" aria-label="Password">
            <button type="submit">Continue</button>
            <p class="additional-options">
                or <a href={path("login")}>login</a>
            </p>
        </form>
    </article>
</Layout>