import Account from './Account.svelte'
import Board from './Board.svelte'
import Expired from './Expired.svelte'
import Login from './Login.svelte'
import Register from './Register.svelte'
import type {Component} from "svelte";

export const views: Record<string, Component> = {
    "Account": Account,
    "Board": Board,
    "Expired": Expired,
    "Login": Login,
    "Register": Register,
}