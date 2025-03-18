package boot

import (
	"FlexcityTest/infrastructure/controller"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func LoadHttpServer() {
	assetController := controller.NewAssetController(assetUsecase)

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	assetRouter := chi.NewRouter()
	assetRouter.Post("/activation", assetController.Activation)

	r.Route("/v1", func(r chi.Router) {
		r.Mount("/assets", assetRouter)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), r); err != nil {
		log.Fatalf("failed to start http server: %s", err)
	}
}
