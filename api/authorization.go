//go:generate protoc --gogofaster_out=plugins=grpc:. --micro_out=. --proto_path=$GOPATH/src:$GOPATH/pkg/mod:. authorization.proto

// Package authorization defines the protocol buffers API for the authorization service.
package authorization
