import { defineConfig } from 'vite';

export default defineConfig({
    build: {
        outDir: 'assets/',
        manifest: true,
        rollupOptions: {
            input: {
                'gc-components': 'webcomponents/gc-components.js',
            }
        }
    }
});
