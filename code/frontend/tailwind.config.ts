import type { Config } from "tailwindcss";

const config: Config = {
  content: ["./app/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}"],
  theme: { extend: { colors: { background: "#FFFFFF", text: "#000000", accent: "#2563EB" } } },
  plugins: []
};

export default config;
