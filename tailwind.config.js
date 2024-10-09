/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
        "./internal/templates/**/*.templ",
        "./internal/templates/components/**/*.templ",
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

