import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// This plugin fixes a bug in Vite v4, where paths that include
// a dot (i.e. a fullstop '.') return a 404. This bug has been fixed
//  in v5, so remove this plugin when upgarding to Vite v5.
import pluginRewriteAll from 'vite-plugin-rewrite-all'

export default defineConfig({
  plugins: [vue(), pluginRewriteAll()],
  build: {
    outDir: '../public', // Output directory for built assets
    emptyOutDir: true, // Required since outDir is outside project directory
    assetsDir: 'assets', // Subdirectory for assets
  },
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:8080',
      },
    },
  },
})
