package handlers

import (
	"sync/atomic"

	"github.com/julienschmidt/httprouter"
)

//Router initilize
func Router() *httprouter.Router {
	router := httprouter.New()
	isReady := &atomic.Value{}
	isReady.Store(false)
	router.GET("/", welcomeHandler)
	router.GET("/version", versionHandler)
	router.GET("/readyz", readyz(isReady))
	router.GET("/healthz", healthz)
	return router
}
