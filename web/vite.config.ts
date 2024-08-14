import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import dotenv from 'dotenv';

dotenv.config();

export default defineConfig({
    server: {
        host: "0.0.0.0",
        port: 7001,
        proxy: {
            "/user": process.env.VITE_API_URL || "http://localhost:6001",
        },
        cors: true
    },
    plugins: [
        vue(),
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    }
})
