package web

import (
	"context"
	"fmt"
	"github.com/go-project-skeleton/src/adapters/httpserver/greethandler"
	"github.com/go-project-skeleton/src/domain/greet"
)

// App holds and creates dependencies for your application.
type App struct {
	Greeter greethandler.GreeterService
}

func newApp(applicationContext context.Context) (*App, error) {
	go handleInterrupts(applicationContext)
	return &App{
		Greeter: greethandler.GreeterServiceFunc(greet.HelloGreeter),
	}, nil
}

// this is just an example of how the services within an app could listen to the
// cancellation signal and stop their work gracefully. So it's probably a decent
// idea to pass the application context to services if you want to do some
// cleanup before finishing.
func handleInterrupts(ctx context.Context) {
	<-ctx.Done()
	fmt.Print(ctx, "shutting down")
}
