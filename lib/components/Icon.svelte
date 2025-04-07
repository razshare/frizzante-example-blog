<script>
    /**
     * @type {string}
     */
    export let value
    export let size = "1.6rem"
    export let fill = "currentColor"
    export let stroke = "none"
</script>

<svg
        xmlns="http://www.w3.org/2000/svg"
        style="width:{size};height:{size}"
        fill={fill}
        viewBox="0 0 24 24"
        stroke={stroke}
>
    <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="1"
            d={value}
    />
</svg>
