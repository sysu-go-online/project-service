package router

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"github.com/sysu-go-online/project-service/controller"
	"github.com/urfave/negroni"
)

var upgrader = websocket.Upgrader{}

// GetServer return web server
func GetServer() *negroni.Negroni {
	r := mux.NewRouter()

	// user collection
	r.Handle("", controller.ErrorHandler(controller.CreateProjectHandler)).Methods("POST")
	r.Handle("/", controller.ErrorHandler(controller.CreateProjectHandler)).Methods("POST")
	r.Handle("", controller.ErrorHandler(controller.ListProjectsHandler)).Methods("GET")
	r.Handle("/", controller.ErrorHandler(controller.ListProjectsHandler)).Methods("GET")

	// Use classic server and return it
	handler := cors.Default().Handler(r)
	s := negroni.Classic()
	s.UseHandler(handler)
	return s
}
