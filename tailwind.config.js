/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
        "./internal/templates/**/*.templ",
        "./internal/templates/components/**/*.templ",
        "./internal/templates/pages/**/*.templ",
        "./internal/templates/shared/**/*.templ",
    ],
  theme: {
    extend: {
            fontFamily: {
                neurial: ['Neurial Grotesk', 'sans-serif'],
            }
        },
  },
  plugins: [],
}

