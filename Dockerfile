FROM golang:bookworm

COPY ./app /app/app

WORKDIR /app

EXPOSE 3000

ENTRYPOINT [ "./app" ]

