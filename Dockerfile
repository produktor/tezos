# Builder container
FROM golang:alpine  as builder

# Set enviroument
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

#RUN mkdir /app/
ADD . /app
WORKDIR /app

# Download depencies
RUN go mod download

# Build and show an progress
RUN go build ./cmd/main

#COPY migrations ./migrations
#RUN adduser -S -D -H -h /app appuser
#USER appuser

# Empty alpine to just run "main" binary
FROM alpine

COPY --from=builder /app/main /app/
COPY --from=builder /app/migrations /app/

EXPOSE 80
EXPOSE 3050

WORKDIR /app

CMD ["./main"]
