import { defineConfig } from "vite"
import { svelte } from "@sveltejs/vite-plugin-svelte"

const IS_DEV = (process.env.DEV ?? "0") === "1"

let sourcemap: "inline" | boolean = false
if (IS_DEV) {
    sourcemap = "inline"
}

// https://vite.dev/config/
export default defineConfig({
    plugins: [
        svelte({
            compilerOptions: {
                css: "injected",
                experimental: {
                    async: true,
                },
            },
        }),
    ],
    resolve: {
        alias: {
            "$lib": "./lib",
            "$exports.client": "./exports.client.ts",
            "$exports.server": "./exports.server.ts",
        },
    },
    build: {
        copyPublicDir: false,
        sourcemap,
        rollupOptions: {
            input: {
                index: "./index.html",
            },
        },
    },
})
