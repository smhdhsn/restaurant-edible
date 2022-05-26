package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/smhdhsn/restaurant-menu/internal/config"

	log "github.com/smhdhsn/restaurant-menu/internal/logger"
	mServContract "github.com/smhdhsn/restaurant-menu/internal/service/contract/menu"
	oServContract "github.com/smhdhsn/restaurant-menu/internal/service/contract/order"
)

// Server contains server's services.
type Server struct {
	menu   mServContract.MenuService
	order  oServContract.OrderService
	conf   *config.ServerConf
	router *mux.Router
}

// New creates a new http server.
func New(
	c *config.ServerConf,
	mService mServContract.MenuService,
	oService oServContract.OrderService,
) (*Server, error) {
	r := mux.NewRouter().StrictSlash(true)

	oHandler := NewOrderHandler(oService)
	mHandler := NewMenuHandler(mService)

	apiGroup := r.PathPrefix("/api").Subrouter()
	apiGroup.
		Methods(http.MethodPost).
		Path("/order").
		HandlerFunc(oHandler.SubmitOrder)
	apiGroup.
		Methods(http.MethodGet).
		Path("/menu").
		HandlerFunc(mHandler.GetMenu)

	return &Server{
		mService,
		oService,
		c,
		r,
	}, nil
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen(c *config.ServerConf) error {
	_, err := config.LoadConf()
	if err != nil {
		return err
	}

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
