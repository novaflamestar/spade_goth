docker d: build
	docker build -t spade_goth .

drun:
	docker run -p 3000:3000 spade_goth

templ:
	templ generate --path ${PWD}/view
go:
	go build -o app .
build: templ go
run:
	./app
full: build run
