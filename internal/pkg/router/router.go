package router

import (
	_ "embed"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"log"
	"net/http"
)

//go:embed assets/notFound.html
var notFound []byte

func Serve() {
	router := httprouter.New()


	router.NotFound = http.HandlerFunc(noFound)
	log.Fatal(http.ListenAndServe(":" + "1234", corsHandler(router)))
}

func noFound(w http.ResponseWriter, r *http.Request){
	headers := w.Header()
	headers["Content-Type"] = []string{"text/html", "charset=utf-8"}
	_, _ = fmt.Fprint(w, string(notFound))
}

// Handle CORS request
func corsHandler(r *httprouter.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		AllowedMethods: []string{"GET","OPTIONS"},
		AllowCredentials: false,
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	return c.Handler(r)
}