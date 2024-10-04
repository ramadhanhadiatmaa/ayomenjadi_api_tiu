FROM golang:1.21.11 AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o amquizdua

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/amquizdua .

EXPOSE 8452

CMD [ "./amquizdua" ]