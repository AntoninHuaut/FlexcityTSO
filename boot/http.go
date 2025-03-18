package boot

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func LoadHttpServer() {
	r := chi.NewRouter()

	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	tsoRouter := chi.NewRouter()
	tsoRouter.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("articles"))
	})

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/tso", tsoRouter)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", HttpPort), r); err != nil {
		log.Fatalf("failed to start http server: %s", err)
	}
}
