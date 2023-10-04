/** @type {import('tailwindcss').Config} */
export default {
    content: ['./src/**/*.{html,js,svelte,ts}'],
    theme: {
        extend: {
            colors: {
                'text': '#1a1a14',
                'background': '#f4f4f0',
                'primary': '#706c93',
                'secondary': '#e9e9e2',
                'accent': '#706c93',
            },
        }
    },
    plugins: []
};
