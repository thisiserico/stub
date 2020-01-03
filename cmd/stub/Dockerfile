FROM golang:1.13-alpine as builder

RUN apk --update add make
RUN adduser -D -g '' stub

WORKDIR /go/src/github.com/thisiserico/stub
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build


FROM scratch

COPY --from=builder /go/src/github.com/thisiserico/stub .
COPY --from=builder /etc/passwd /etc/passwd

USER stub
ENTRYPOINT ["./stub"]

