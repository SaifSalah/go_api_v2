FROM golang:1.16 AS builder_stage

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS Production
COPY --from=builder_stage /app .
CMD ["./app"]
