FROM golang:1.20 as builder

WORKDIR /go/src/github.com/monotek/mongodb-exporter

COPY . .

RUN make build

FROM alpine:3.17

COPY --from=builder /go/src/github.com/monotek/mongodb-exporter/bin/mongodb_exporter /bin/mongodb_exporter

EXPOSE 9216

ENTRYPOINT [ "/bin/mongodb_exporter" ]
