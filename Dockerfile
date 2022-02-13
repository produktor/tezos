FROM golang:alpine

ENV GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

WORKDIR /dist

COPY main .

COPY migrations ./migrations

EXPOSE 80

CMD ["/dist/main"]
