package main

import (

	"database/sql"
	_"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/obasootom/langtranslator/api"
	"github.com/obasootom/langtranslator/config"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

func main() {
	config, err := config.LoadConfigClient(".")
	if err != nil {
		return
	}
	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Panic("cannot open database")
	}

   
	store := db.NewStore(conn)
	server,err := api.NewServer(store,config)
	if err != nil {
		log.Panic("cannot create http server")
	}
	
   server.Start(config.HTTP_ADDRESS_CLIENT)

}