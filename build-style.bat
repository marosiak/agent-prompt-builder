tailwindcss -i ./tailwind/config.css -o ./web/bundle.css --minify
copy .\web\bundle.css .\docs\bundle.css
