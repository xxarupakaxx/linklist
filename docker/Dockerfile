FROM golang:1.17-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/xxarupakaxx/linklist
COPY . .
RUN go mod download
RUN go build main.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/xxarupakaxx/linklist /app

CMD /app/main $PORT
