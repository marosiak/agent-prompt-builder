build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/server
	go build -o output/app_binary ./cmd/server

release:
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o docs/web/app.wasm ./cmd/server

	wasm-opt ./docs/web/app.wasm \
		--enable-bulk-memory \
		-Oz \
		--vacuum \
		--remove-unused-module-elements \
		--dce \
		--output=./docs/web/app.wasm

	./build-style.sh
	cd docs && go run .././cmd/static/

install_linter:
	go install github.com/mgechev/revive

run:
	GOARCH=wasm GOOS=js go build -o web/app.wasm ./cmd/server
	go run -o output/app_binary ./cmd/server