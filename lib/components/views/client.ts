export const views: Record<string, Promise<unknown>> = {
    "Account": import('./Account.svelte'),
    "Board": import('./Board.svelte'),
    "Expired": import('./Expired.svelte'),
    "Login": import('./Login.svelte'),
    "Register": import('./Register.svelte'),
}