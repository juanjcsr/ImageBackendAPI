package routers

import (
	"api.jwt.auth/controllers"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle("/test/hello",
		negroni.New(
			negroni.HandlerFunc(controllers.HelloController),
		)).Methods("GET")

	return router
}
