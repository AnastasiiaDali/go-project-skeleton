package web

import (
	"fmt"
	"github.com/go-project-skeleton/src/adapters/httpserver"
	"log"
)

func main() {
	ctx, done := listenForCancellationAndAddToContext()
	defer done()

	config, err := loadAppConfig()
	if err != nil {
		log.Fatalf("failed to load config - %v", err)
	}

	app, err := newApp(ctx)
	if err != nil {
		fmt.Errorf("failed to create app: %w", err)
	}

	router := httpserver.NewRouter(
		app.Greeter,
	)

	server := httpserver.NewWebServer(config.ServerConfig, router)

	fmt.Sprintf("Started. Listening on port: %s", config.ServerConfig.Port)
	if err = server.ListenAndServe(); err != nil {
		fmt.Errorf("http server listen failed: %w", err)
	}
}
