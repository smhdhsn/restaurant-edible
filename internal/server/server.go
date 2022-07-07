package server

import (
	"fmt"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/smhdhsn/restaurant-edible/internal/config"
	"github.com/smhdhsn/restaurant-edible/internal/server/resource"

	log "github.com/smhdhsn/restaurant-edible/internal/logger"
	inventoryProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/inventory"
	menuProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/menu"
	recipeProto "github.com/smhdhsn/restaurant-edible/internal/protos/edible/recipe"
)

// Server contains server's services.
type Server struct {
	listener net.Listener
	grpc     *grpc.Server
	conf     *config.ServerConf
}

// New creates a new http server.
func New(c *config.ServerConf, eRes *resource.EdibleResource) (*Server, error) {
	// listen to a specific host and port for incoming requests.
	l, err := net.Listen(c.Protocol, fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		return nil, errors.Wrap(err, "failed to listen to port")
	}

	// instantiate gRPC server.
	s := grpc.NewServer()

	// register gRPC service handlers.
	inventoryProto.RegisterEdibleInventoryServiceServer(s, eRes.InventoryHandler)
	recipeProto.RegisterEdibleRecipeServiceServer(s, eRes.RecipeHandler)
	menuProto.RegisterEdibleMenuServiceServer(s, eRes.MenuHandler)

	return &Server{
		listener: l,
		grpc:     s,
		conf:     c,
	}, nil
}

// Listen is responsible for starting the HTTP server.
func (s *Server) Listen() error {
	log.Info(fmt.Sprintf("%s server started listening on port <%d>", s.conf.Protocol, s.conf.Port))
	return s.grpc.Serve(s.listener)
}
