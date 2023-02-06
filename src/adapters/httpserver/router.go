package httpserver

import (
	"github.com/go-project-skeleton/src/adapters/httpserver/greethandler"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	greetPath       = "/greet/{name}"
	healthCheckPath = "/internal/healthcheck"
	metricsPath     = "/internal/metrics"
)

func NewRouter(
	greeter greethandler.GreeterService,
) *mux.Router {
	greetingHandler := greethandler.NewGreetHandler(greeter)

	router := mux.NewRouter()

	router.Handle(greetPath, http.HandlerFunc(greetingHandler.Greet))
	return router
}
