package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"github.com/vfuntikov/simple_bank/api"
	db "github.com/vfuntikov/simple_bank/db/sqlc"
	"github.com/vfuntikov/simple_bank/util"
)

func main() {
	var err error

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	if err = db.MigrateDB(config.MigrationURL, config.DBSource); err != nil {
		log.Fatal("cannot migrate database:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
