package greethandler

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type GreeterService interface {
	Greet(ctx context.Context, name string) (greeting string, err error)
}

type GreeterServiceFunc func(context.Context, string) (string, error)

func (g GreeterServiceFunc) Greet(ctx context.Context, name string) (greeting string, err error) {
	return g(ctx, name)
}

type GreetHandler struct {
	greeter GreeterService
}

func NewGreetHandler(
	greeter GreeterService,
) *GreetHandler {
	return &GreetHandler{
		greeter: greeter,
	}
}

func (g *GreetHandler) Greet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	name := mux.Vars(r)["name"]

	greeting, err := g.greeter.Greet(ctx, name)
	if err != nil {
		fmt.Errorf("failed to greet %s", name)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprint(w, greeting)
}
