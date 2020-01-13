FROM golang:1.13-alpine as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GOFLAGS=-mod=vendor

RUN apk --update add make
RUN adduser -D -g '' stub

WORKDIR /go/src/github.com/thisiserico/stub
COPY . .

RUN make build


FROM scratch

COPY --from=builder /go/src/github.com/thisiserico/stub .
COPY --from=builder /etc/passwd /etc/passwd

USER stub
ENTRYPOINT ["./stub"]

