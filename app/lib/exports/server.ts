import Account from '$lib/views/Account.svelte'
import Board from '$lib/views/Board.svelte'
import Expired from '$lib/views/Expired.svelte'
import Login from '$lib/views/Login.svelte'
import Register from '$lib/views/Register.svelte'
import type {Component} from "svelte";

export const views: Record<string, Component> = {
    "Account": Account,
    "Board": Board,
    "Expired": Expired,
    "Login": Login,
    "Register": Register,
}