FROM golang:1.21-alpine3.18 AS builder

WORKDIR /src

ADD . /src

RUN go env -w CGO_ENABLED=0
RUN go build -o /bin/geo-port

FROM scratch AS release
COPY --from=builder /bin/geo-port /geo-port

ENTRYPOINT ["/geo-port"]