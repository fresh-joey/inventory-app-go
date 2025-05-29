FROM golang:1.24 AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN go mod tidy

# Build the Go application with CGO disabled for a static binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o inventory .


FROM alpine:3.20

WORKDIR /app

COPY --from=build /app/inventory .

COPY --from=build /app/assets ./assets

EXPOSE 8085

CMD [ "./inventory" ]
