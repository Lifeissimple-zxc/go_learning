package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewHTTPServer,
			fx.Annotate(NewServeMux, fx.ParamTags(`name:"echo"`, `name:"hello"`)),
			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),
				fx.ResultTags(`name:"echo"`),
			),
			fx.Annotate(
				NewHelloHandler,
				fx.As(new(Route)),
				fx.ResultTags(`name:"hello"`),
			),
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}

// NewHTTPServer builds an HTTP server that will serve requests
// when the Fx application starts
func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	srv := &http.Server{Addr: ":8080", Handler: mux}
	// lc-related commands tell the Fx how the server should be started & stopped
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			// fmt.Println("Starting HTTP server at", srv.Addr) //pre-log way of doing things
			log.Info("Starting HTTP server", zap.String("addr", srv.Addr))
			// Concurrent call to start the server
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

// Route is an http.Handler that knows the mux pattern
// under which it will be registered
type Route interface {
	http.Handler

	// Reports the path at which this is registered
	Pattern() string
}

// This implements the Route interface
type EchoHandler struct {
	log *zap.Logger
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

// NewEchoHandler builds a new EchoHandler.
func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

// ServerHTTP handles an HTTP request to the /echo endpoint. Receiver func.
func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		// fmt.Fprintln(os.Stderr, "Failed to handle request:", err) // pre-log method of sending a message
		h.log.Warn("Failed to handle request", zap.Error(err))
	}
}

func NewServeMux(route1, route2 Route) *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle(route1.Pattern(), route1)
	mux.Handle(route2.Pattern(), route2)
	return mux
}

type HelloHandler struct {
	log *zap.Logger
}

// Implement our route interface in HelloHandler via a receiver
func (*HelloHandler) Pattern() string {
	return "/hello"
}

// Now we also define a constructor for the above handler
func NewHelloHandler(log *zap.Logger) *HelloHandler {
	return &HelloHandler{log: log} // Create an instance of our handler and return its memory address
}

// ServeHTTP implementation for hello handler
func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("Failed to read request", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError) // this responds to client?
		return
	}

	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
