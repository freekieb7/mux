package main

import (
	"github.com/freekieb7/mux"
	"github.com/freekieb7/mux/middleware"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.Group("a", func(router mux.Router) {
		router.Use(middleware.NewDefaultLoggingMiddleware())

		router.Get("/a", PathCheckHandler)
		router.Get("b", PathCheckHandler)
		router.Get("{c}", DynamicHandler)
	})

	//router.Get("/a", PathCheckHandler)
	//router.Get("/a/b", PathCheckHandler)
	//router.Get("/a/b/c", PathCheckHandler)
	//router.Get("/a/b/{d}", PathCheckHandler)
	//router.Get("/a/b/{e:[a-z]+}", PathCheckHandler)
	//router.Get("/test/{id:[0-9]+}/asd", PathCheckHandler)
	//router.Get("/test/{id:[0-9]+}/{bla}", PathCheckHandler)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func PathCheckHandler(response http.ResponseWriter, request *http.Request) {
	_, _ = response.Write([]byte("Path: " + request.URL.Path))
}

func DynamicHandler(response http.ResponseWriter, request *http.Request) {
	_, _ = response.Write([]byte("Path param: " + mux.PathParam(request, "c") + "\n"))
	_, _ = response.Write([]byte("Query param: " + mux.QueryParam(request, "test")))
}