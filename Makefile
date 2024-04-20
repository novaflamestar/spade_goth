docker d: build
	docker build -t spade_goth .

drun:
	docker run -p 3000:3000 spade_goth

dfull: d drun

templ:
	templ generate --path ${PWD}/view
go:
	go build -o app .
tail:
	npx tailwindcss -i ./styles/tailwind.css -o ./styles/dist/styles.css
build: templ go
run:
	./app
full: build run
