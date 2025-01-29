/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: [
      {
        terminal: {
          "primary": "#ff6b6b",      // Red color for text highlights
          "secondary": "#a3a3a3",   // Neutral secondary text
          "accent": "#d4d4d4",      // Accent color for hover effects
          "neutral": "#525252",     // Neutral gray
          "base-100": "#181818",    // Background color (matches bg-gray-900)
          "base-200": "#272727",    // Slightly lighter background
          "base-300": "#3f3f3f",    // Border colors
          "base-content": "#d1d5db" // Text color (matches text-gray-300)
        },
      },
    ],
  },
}

