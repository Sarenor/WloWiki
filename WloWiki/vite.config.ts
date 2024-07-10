import { purgeCss } from 'vite-plugin-tailwind-purgecss';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit(), purgeCss()],
	server: {
        host: '0.0.0.0',
        port: 5173,
        strictPort: true, // Ensures Vite uses the specified port
        hmr: {
            protocol: 'ws', // WebSocket protocol for HMR
            host: 'web.astagan.pl', // or your server's hostname
            port: 5173 // or the port your VPN exposes
        },
		fs: {
		  allow: ['.']  // Allow serving files from all directories
		}
	  }
});
