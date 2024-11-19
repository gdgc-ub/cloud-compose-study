import { resolve } from "path";
import { defineConfig } from "vite";

export default defineConfig({
    build: {
        lib: {
            entry: [resolve(__dirname, "resources/htmx.js")],
            formats: ["es"],            
            fileName: "htmx",
        },
        outDir: "static/js",
        emptyOutDir: false
    }
});