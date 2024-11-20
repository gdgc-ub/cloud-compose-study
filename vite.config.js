import { resolve } from "path";
import { defineConfig } from "vite";

export default defineConfig({
    build: {
        lib: {
            entry: [resolve(__dirname, "resources/dev/htmx.js")],
            formats: ["es"],            
            fileName: "htmx",
        },
        outDir: "resources/static/js",
        emptyOutDir: false
    }
});