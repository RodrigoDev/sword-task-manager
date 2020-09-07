package transport

import (
	"context"
	"net"
	"net/http"
)

func ListenAndServe(ctx context.Context, address string, h http.Handler) error {
	l, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	defer l.Close()

	return Serve(ctx, l, h)
}
