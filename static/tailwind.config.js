/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./html/*.html",],
  theme: {
    extend: {
      colors: {
        transparent: 'transparent',
        current: 'currentColor',
        'color1': '#d3daf3',
        'color2': '#9092fb',
        'color3': '#f33862',
        'color4': '#2C3639',
        'color5': '#676afb',
      backgroundImage: {
        'frontend': "url('https://i.pinimg.com/originals/14/fc/d1/14fcd189633936157ee354b6a092169d.png')",
        'footer-texture': "url('/img/footer-texture.png')",
        }
      },
    },
  },
  plugins: [],
}
