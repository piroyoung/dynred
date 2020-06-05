FROM golang:1.14.4-alpine3.12 as builder

COPY ./dynred /go/dynred
WORKDIR /go/dynred
RUN go install -v ./cmd/dynred/


FROM alpine:3.12

COPY --from=builder /go/bin/ /go/bin/
COPY --from=builder /go/dynred/static /go/dynred/static

WORKDIR /go/dynred

CMD ["/go/bin/dynred"]