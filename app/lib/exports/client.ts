export const views: Record<string, Promise<unknown>> = {
    "Account": import('$lib/views/Account.svelte'),
    "Board": import('$lib/views/Board.svelte'),
    "Expired": import('$lib/views/Expired.svelte'),
    "Login": import('$lib/views/Login.svelte'),
    "Register": import('$lib/views/Register.svelte'),
}