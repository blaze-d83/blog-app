/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
        "./internal/templates/*.{html, js, templ, go}",
        "./internal/templates/components/*.{html, js, templ, go}",
    ],
  theme: {
    extend: {},
  },
  plugins: [],
}

