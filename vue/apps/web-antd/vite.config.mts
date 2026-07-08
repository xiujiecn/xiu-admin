import { fileURLToPath } from 'node:url';

import { defineConfig } from '@vben/vite-config';

const emptyNodeLoaderModule = '\0empty-node-loader-module';
const emptyNodeLoaderPath = fileURLToPath(
  new URL('./src/shims/node-loader-empty.ts', import.meta.url),
);

export default defineConfig(async ({ command }) => {
  const isBuild = command === 'build';

  return {
    application: {},
    vite: {
      build: {
        target: 'es2022',
      },
      plugins: isBuild
        ? [
            {
              enforce: 'pre',
              name: 'dtu-empty-node-loader',
              resolveId(source) {
                if (
                  source === 'jiti' ||
                  source.startsWith('jiti/') ||
                  source.includes('/node_modules/.pnpm/jiti@') ||
                  source.includes('/node_modules/jiti/')
                ) {
                  return emptyNodeLoaderModule;
                }
              },
              load(id) {
                if (id === emptyNodeLoaderModule) {
                  return [
                    'const noop = () => undefined;',
                    'const proxy = new Proxy(noop, {',
                    '  get() { return proxy; },',
                    '  apply() { return proxy; },',
                    '});',
                    'export const createJiti = () => proxy;',
                    'export const createRequire = () => proxy;',
                    'export default proxy;',
                  ].join('\n');
                }
              },
            },
          ]
        : [],
      resolve: {
        alias: isBuild
          ? [
              {
                find: /^jiti(\/.*)?$/,
                replacement: emptyNodeLoaderPath,
              },
            ]
          : [],
        conditions: isBuild
          ? ['browser', 'import', 'module', 'default']
          : ['development', 'browser', 'import', 'module', 'default'],
      },
      server: {
        proxy: {
          '/api': {
            changeOrigin: true,
            rewrite: (path) => path.replace(/^\/api/, ''),
            target: 'http://localhost:8199/api/v1',
            ws: true,
          },
          '/upload': {
            changeOrigin: true,
            rewrite: (path) => path.replace(/^\/upload/, ''),
            target: 'http://localhost:8199/upload',
            ws: true,
          },
        },
      },
    },
  };
});
