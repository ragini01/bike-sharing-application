FROM golang:1.20-alpine AS build

WORKDIR /app

COPY . /app

COPY main.go go.* /app/


RUN CGO_ENABLED=0 go build -o main .

FROM scratch
COPY --from=build main  main

EXPOSE 8080

CMD ["./main"]


