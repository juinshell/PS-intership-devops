#build stage
FROM golang:alpine AS builder
RUN apk -U --no-cache add git
WORKDIR /go/src/app
COPY . .
RUN go build -o /go/bin/app -v server/server.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
WORKDIR /
ENTRYPOINT /app