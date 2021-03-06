# backend
FROM golang:1.17.3-alpine as builder
RUN apk add --update git

WORKDIR /go/src/github.com/xxarupakaxx/linklist
COPY go.mod go.sum ./
RUN go mod download

COPY docker .
RUN go build -o /linklist -ldflags "-s -w"

# runtime image
FROM alpine:3.14.2
WORKDIR /app

RUN apk --update --no-cache add tzdata \
  && cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
  && apk del tzdata \
  && mkdir -p /usr/share/zoneinfo/Asia \
  && ln -s /etc/localtime /usr/share/zoneinfo/Asia/Tokyo
RUN apk --update --no-cache add ca-certificates \
  && update-ca-certificates \
  && rm -rf /usr/share/ca-certificates

COPY --from=builder /linklist ./

ENTRYPOINT ./linklist
