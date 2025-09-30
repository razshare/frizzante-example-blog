import Board from "$lib/views/Board.svelte"
import Expired from "$lib/views/Expired.svelte"
import Login from "$lib/views/Login.svelte"
import Register from "$lib/views/Register.svelte"
import Form from "$lib/views/Form.svelte"

export const views = {
    Form: Form,
    Board: Board,
    Expired: Expired,
    Login: Login,
    Register: Register,
}
