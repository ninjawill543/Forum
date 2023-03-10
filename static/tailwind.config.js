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
        'color3': '#F4882A',
        'color4': '#2C3639',
        'color5': '#676afb',
        'newcolor1': '#ffffff',
        'newcolor2': '#E5E7EB',
        'newcolor3': '#C5CAD3',
        'newcolor4': '#FF7700',
        'newcolor5': '#E5E7EB',
        'newcolor6': '#F2EFEF',
      backgroundImage: {
        'frontend': "url('https://i.pinimg.com/originals/14/fc/d1/14fcd189633936157ee354b6a092169d.png')",
        'footer-texture': "url('/img/footer-texture.png')",
        }
      },
    },
  },
  plugins: [],
}