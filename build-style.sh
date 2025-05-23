#!/bin/bash

if ! command -v tailwindcss &> /dev/null
then
    echo "Błąd: 'tailwindcss' is not installed or missing in PATH."
    exit 1
fi

tailwindcss -i ./tailwind/config.css -o ./web/bundle.css --minify

cp ./web/bundle.css ./docs/bundle.css
