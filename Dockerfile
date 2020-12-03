FROM golang:1.14

WORKDIR /go/src/github.com/magiskboy/katalog

COPY . .

RUN make build

FROM alpine:3.12

EXPOSE 8080

COPY --from=0 /go/src/github.com/magiskboy/katalog/katalog /usr/bin

CMD ["katalog", "rest"]
