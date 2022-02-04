package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/smhdhsn/food/internal/config"
	"github.com/smhdhsn/food/util/cli"
)

// Server contains server's services.
type Server struct {
	router *mux.Router
}

// New creates a new http server.
func New() (*Server, error) {
	r := mux.NewRouter().StrictSlash(true)

	apiGroup := r.PathPrefix("/api").Subrouter()
	apiGroup.
		Methods(http.MethodGet).
		Path("/menu").
		HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
			fmt.Println("You're here.")
		})

	return &Server{
		r,
	}, nil
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(host string, port int) error {
	conf, err := config.LoadConf()
	if err != nil {
		return err
	}

	c := cors.New(cors.Options{
		AllowedOrigins: conf.ClientURI,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	fmt.Printf("%sStarting server%s: <http://%s:%d>\n", cli.GREEN, cli.RESET, host, port)
	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), c.Handler(s.router))
}
