package handlers

import (
	"net/http"
	"sync/atomic"

	"github.com/julienschmidt/httprouter"
)

func readyz(isReady *atomic.Value) httprouter.Handle {
	return func(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
