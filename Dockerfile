FROM golang:bullseye

COPY . /app/

RUN GOPATH=/temp/ go install github.com/a-h/templ/cmd/templ@v0.2.663

WORKDIR /app

RUN PATH=/temp/bin:$PATH make build

FROM debian:bullseye-slim

COPY --from=0 /app/app /app/app

WORKDIR /app

EXPOSE 3000

ENTRYPOINT [ "./app" ]
