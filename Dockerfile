FROM golang:latest as builder

WORKDIR /build

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o guzfolio server/*

WORKDIR /dist

RUN cp /build/guzfolio ./guzfolio
RUN cp /build/docker/wait-for-it.sh ./wait-for-it.sh

FROM alpine:latest

RUN apk add --no-cache bash

COPY --from=builder /dist /

EXPOSE 8080

CMD ["./guzfolio"]