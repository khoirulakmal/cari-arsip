/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./*.{html,js}"],
  theme: {
    extend: {
      fontFamily: {
        headline: ['"Source Serif 4"', "ui-serif"], // Adds a new `font-display` class
      }
    },
  },
  plugins: [
    require('daisyui'),
  ],
}

