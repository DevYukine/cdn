FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/cdn

COPY . .

RUN go get -d -v

RUN go build -o /go/bin/cdn

FROM scratch

COPY --from=builder /go/bin/cdn /go/bin/cdn

ENTRYPOINT ["/go/bin/cdn"]