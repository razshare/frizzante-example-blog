<style>
    li {
        padding: 1rem;
    }
</style>

<script lang="ts">
    import { href } from "$lib/utilities/scripts/href.ts"
</script>

<nav>
    <ul>
        <li></li>
    </ul>
    <ul>
        <li>
            <a {...href("/article-form")}>
                <button>New Article</button>
            </a>
        </li>
        <li>
            <a {...href("/logout")}>
                <button>Logout</button>
            </a>
        </li>
    </ul>
</nav>
