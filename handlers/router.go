package handlers

import (
	"github.com/julienschmidt/httprouter"
)

//Router initilize
func Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", welcome)
	return router
}
