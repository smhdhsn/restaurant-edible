package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/smhdhsn/restaurant-edible/internal/config"
	"github.com/smhdhsn/restaurant-edible/internal/server/resource"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
)

// Server contains server's services.
type Server struct {
	eRes   *resource.EdibleResource
	conf   *config.ServerConf
	router *mux.Router
}

// New creates a new http server.
func New(c *config.ServerConf, eRes *resource.EdibleResource) (*Server, error) {
	r := mux.NewRouter().StrictSlash(true)

	apiGroup := r.PathPrefix("/api").Subrouter()
	apiGroup.
		Methods(http.MethodGet).
		Path("/menu").
		HandlerFunc(eRes.MenuHandler.GetMenu)

	return &Server{
		eRes:   eRes,
		conf:   c,
		router: r,
	}, nil
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(c *config.ServerConf) error {
	cor := cors.New(cors.Options{
		AllowedOrigins: []string{},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	log.Info(fmt.Sprintf("%s server started listening on port <%d>", s.conf.Protocol, s.conf.Port))
	return http.ListenAndServe(fmt.Sprintf("%s:%d", c.Host, c.Port), cor.Handler(s.router))
}
