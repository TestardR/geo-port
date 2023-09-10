FROM golang:1.21-alpine3.18 as builder

WORKDIR /src
ADD . /src
RUN go build -o /bin/geo-port

FROM scratch
COPY --from=builder /bin/geo-port /bin/geo-port

ENTRYPOINT ["/bin/geo-port"]