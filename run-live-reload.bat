go install github.com/bokwoon95/wgo@latest
wgo cmd /c "set GOOS=js&& set GOARCH=wasm&& go build -o web/app.wasm" ./cmd/server