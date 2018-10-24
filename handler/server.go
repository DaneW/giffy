package handler

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/danew/giffy/util"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var webRoot string

// NewServer configures and returns a HTTP Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter().StrictSlash(true)

	repo := initRepository()
	initRoutes(mx, formatter, repo)
	n.UseHandler(mx)

	return n
}

func initRepository() (repo gifRepository) {
	repo = newInMemoryRepository()
	return
}

func initRoutes(mx *mux.Router, formatter *render.Render, repo gifRepository) {
	root, err := os.Getwd()
	if err != nil {
		panic("Could not retrieve working directory")
	} else {
		webRoot = util.Env("WEBROOT", root)
	}

	api := mx.PathPrefix("/api/v1").Subrouter()

	api.Path("/gifs").Queries("q", "{query}").Queries("offset", "{offset}").HandlerFunc(getGIFsHandler(formatter, repo)).Methods("GET").Name("OffsetHandler")
	api.Path("/gifs").Queries("q", "{query}").HandlerFunc(getGIFsHandler(formatter, repo)).Methods("GET")
	api.Path("/gifs/{id}/upvote").HandlerFunc(upVoteGIF(formatter, repo)).Methods("PUT")
	api.Path("/gifs/{id}/downvote").HandlerFunc(downVoteGIF(formatter, repo)).Methods("PUT")

	mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/build")))
}
