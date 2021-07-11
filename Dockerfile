# Stage 1
FROM golang:alpine as builder

LABEL mainteiner="Arthur Santana <artsantana1457@gmail.com>"

RUN apk update && apk add --no-cache git
RUN mkdir /build 
ADD . /build/
WORKDIR /build
RUN go get -d -v
RUN go build -o server .


# Stage 2
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/ /app/
WORKDIR /app
CMD ["./server"]

EXPOSE 8081
