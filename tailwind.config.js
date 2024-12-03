/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './templates/**/*.{templ,html}',
  ],
  theme: {
    extend: {
      fontFamily: {
        inter: ['Inter', 'sans-serif'],
      },
      height: {
        '18': '4.5rem'
      },

      aspectRatio: {
        '4/3': '4 / 3',
        '3/4': '3 / 4',
        '4/1': '4 / 1',
        '9/16': '9 / 16',
      },
    },
  },
  plugins: [],
}
