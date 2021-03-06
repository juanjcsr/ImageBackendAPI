package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"jezz.api.mobile/controllers"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
