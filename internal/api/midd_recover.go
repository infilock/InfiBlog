package api

import (
	"github.com/infilock/InfiBlog/pkg/res"
	"github.com/pkg/errors"
	"net/http"
)

func (mw *recoverMW) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if v := recover(); v != nil {
			res.Done(w, r, res.InternalServerError(errors.WithStack(errors.Errorf("panic recovered: %+v", v))))
		}
	}()
	mw.next.ServeHTTP(w, r)
}

func (h *handler) RecoverPanic() Middleware {
	return func(next http.Handler) http.Handler {
		return &recoverMW{
			next: next,
		}
	}
}
