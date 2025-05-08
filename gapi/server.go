package gapi

import (
	"fmt"

	db "github.com/vfuntikov/simple_bank/db/sqlc"
	"github.com/vfuntikov/simple_bank/pb"
	"github.com/vfuntikov/simple_bank/token"
	"github.com/vfuntikov/simple_bank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, db db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      db,
		config:     config,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
