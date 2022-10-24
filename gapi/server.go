package gapi

import (
	"fmt"

	db "github.com/imperiutx/nan_forum/db/sqlc"
	"github.com/imperiutx/nan_forum/pb"
	"github.com/imperiutx/nan_forum/token"
	"github.com/imperiutx/nan_forum/utils"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedNanForumServer
	config     utils.Config
	store      db.Store
	tokenMaker token.PasetoMaker
}

// NewServer creates a new gRPC server.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: *tokenMaker,
	}

	return server, nil
}
