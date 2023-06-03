/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        black: "#272343",
        white: "#FBFBFB",
        primary: "#395AFF",
        primary_light: "#455dd6",

        gradient_start: "#12c2e9",
        gradient_center: "#c471ed",
        gradient_end: "#f64f59",
      },
      fontFamily: {
        open_sans: ["Open Sans", "sans-serif"],
      },
      gridTemplateColumns: {
        fluid: "repeat(auto-fill, minmax(280px, 1fr))",
      },
    },
  },
  plugins: [],
};
