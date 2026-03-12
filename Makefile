build:
	@if not exist build mkdir build
	go build -o build/url_shortener.exe ./cmd/server/main.go

run:
	go run ./cmd/server/main.go

clean: 
	@if exist build rmdir /s /q build