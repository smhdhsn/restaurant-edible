package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/smhdhsn/restaurant-menu/internal/config"
	"github.com/smhdhsn/restaurant-menu/internal/service"
	"github.com/smhdhsn/restaurant-menu/util/cli"
)

// Server contains server's services.
type Server struct {
	menu   *service.MenuService
	order  *service.OrderService
	router *mux.Router
}

// New creates a new http server.
func New(
	mService *service.MenuService,
	oService *service.OrderService,
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
