package main

import (
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New().Run()
}

// NewHTTPServer builds an HTTP server that will serve requests
// when the Fx application starts
func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	srv := &http.Server{Addr: ":8080"}
	return srv
}
