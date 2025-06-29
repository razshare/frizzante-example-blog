import Board from "$lib/views/Board.svelte"
import Expired from "$lib/views/Expired.svelte"
import Login from "$lib/views/Login.svelte"
import Register from "$lib/views/Register.svelte"
import ArticleForm from "$lib/views/ArticleForm.svelte"

export const views = {
    ArticleForm: ArticleForm,
    Board: Board,
    Expired: Expired,
    Login: Login,
    Register: Register,
}
