package main

import (
	"database/sql"
	_"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/obasootom/langtranslator/translator/api"
	"github.com/obasootom/langtranslator/translator/config"
	db "github.com/obasootom/langtranslator/translator/db/sqlc"
)

func main() {
	config, err := config.LoadConfigTranslator("..")
	if err != nil {
		return
	}
	conn, err := sql.Open(config.DB_DRIVER, "")
	if err != nil {
		log.Panic("cannot open database")
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Panic("cannot create http server")
	}
	server.Startl(config.HTTP_ADDRESS_TRANS)
}


