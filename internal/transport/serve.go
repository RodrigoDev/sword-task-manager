package transport

import (
	"context"
	"net"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

// Serve serves the request
func Serve(ctx context.Context, l net.Listener, h http.Handler) error {
	s := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      h,
	}

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := s.Serve(l); err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		return s.Shutdown(ctx)
	})

	return g.Wait()
}
