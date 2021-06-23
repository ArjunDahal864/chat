package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"github.com/arjun/chat/backend/util"
)

func ConnectDatabase() (*sql.DB, error) {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return conn, nil

}
