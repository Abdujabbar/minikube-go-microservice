package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/julienschmidt/httprouter"
)

//Router initilize
func Router() *httprouter.Router {
	router := httprouter.New()
	isReady := &atomic.Value{}
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()
	isReady.Store(false)
	router.GET("/", welcomeHandler)
	router.GET("/version", versionHandler)
	router.GET("/readyz", readyz(isReady))
	router.GET("/healthz", healthz)
	return router
}
