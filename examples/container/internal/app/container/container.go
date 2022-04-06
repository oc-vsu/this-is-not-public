package container

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func Initialize() *App {
	app := &App{
		Router: router(),
	}

	return app
}

func (app *App) Run(port string) {
	log.Println("starting service on port: %s", port)
	log.Fatal(http.ListenAndServe(port, app.Router))
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/hello-world", HelloWorld).Methods(http.MethodGet, http.MethodOptions)

	return router
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("hello world"))
}
