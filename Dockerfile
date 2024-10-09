FROM golang:1.21.11 AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o amtiu

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/amtiu .

EXPOSE 8452

CMD [ "./amtiu" ]