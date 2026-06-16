import type { Config } from 'tailwindcss';

const config: Config = {
  content: ['./app/**/*.{ts,tsx}', './components/**/*.{ts,tsx}', './features/**/*.{ts,tsx}'],
  theme: {
    extend: {
      colors: {
        castle: {
          cream: '#FFF7E8',
          pink: '#F9A8D4',
          purple: '#8B5CF6',
          night: '#312E81',
        },
      },
    },
  },
  plugins: [],
};

export default config;
