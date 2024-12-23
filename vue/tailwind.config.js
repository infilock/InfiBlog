/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors:{
        primary:'#4377FE',
        lightBlueBg:'#E8EEFF',
        lightBlackBg:'#383D4D',
        blackBg70:'#9B9EA6',
      }
    },
  },
  plugins: [],
}

