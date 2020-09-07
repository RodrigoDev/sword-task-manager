package transport

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct{ router *mux.Router }

func New(opts ...Option) (*Handler, error) {
	h := &Handler{router: mux.NewRouter()}

	if err := h.Apply(opts...); err != nil {
		return nil, err
	}

	return h, nil
}

// Apply will apply the options passed.
func (h *Handler) Apply(opts ...Option) error {
	for _, opt := range opts {
		if err := opt(h); err != nil {
			return err
		}
	}
	return nil
}

// ServeHTTP calls the serve method on the router
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
