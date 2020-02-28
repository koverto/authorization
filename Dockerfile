FROM golang:1.14-alpine AS builder
WORKDIR /go/src/github.com/koverto/authorization
ENV CGO_ENABLED=0 GO111MODULE=on

RUN apk add bash build-base curl git

COPY go.mod go.sum ./
RUN go mod download

COPY api api/
COPY cmd cmd/
COPY internal internal/
RUN go test ./... && \
    go install ./cmd/...

FROM scratch
COPY --from=builder /go/bin/authorization /go/bin/
ENTRYPOINT [ "/go/bin/authorization" ]
