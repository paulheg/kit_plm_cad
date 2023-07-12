FROM golang:latest AS build

COPY . /app
WORKDIR /app

RUN go mod download 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o server cmd/main.go

FROM alpine:latest

COPY --from=build /app/server /app/server
WORKDIR /app

EXPOSE 3000

CMD [ "./server" ]
