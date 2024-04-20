templ:
	templ generate --path ./view
go:
	go build -o app .
build: templ go
run:
	./app
full: build run
