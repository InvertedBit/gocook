/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./views/*.templ"], // This is where your HTML templates / JSX files are located
  theme: {
    extend: {
      fontFamily: {
        sans: ["Funnel Sans", "sans-serif"],
        mono: ["Funnel Sans", "monospace"],
        serif: ["Funnel Sans", "serif"],
      },
    },
  },
  plugins: [],
};
