package main

import (
	"github.com/codegangsta/negroni"
	"jezz.api.mobile/routers"
	"jezz.api.mobile/settings"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":8080", n)
}
