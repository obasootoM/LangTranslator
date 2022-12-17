package main

import (
	"database/sql"
	"fmt"

	"log"

	_ "github.com/lib/pq"
	"github.com/obasootom/langtranslator/api"
	"github.com/obasootom/langtranslator/config"
	db "github.com/obasootom/langtranslator/db/sqlc"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		return
	}
	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE)
	if err != nil {
		log.Panic("cannot open database")
	}
	store := db.NewStore(conn)
	http := RunHttp(*store, config)
	fmt.Println()
	https := RunHttps(*store, config)
	select {
	case err := <-http:
		log.Panic("cannot load http", err)
	case err := <-https:
		log.Panic("cannot load https", err)
	}

}

func RunHttps(store db.Store, config config.Config) chan error {
	server, err := api.NewServer(&store)
	if err != nil {
		log.Panic("cannot create https server", err)
	}
	runhttps := make(chan error)
	go func() {
		runhttps <- server.StartTls(":8000")
	}()
	return runhttps
}

func RunHttp(store db.Store, config config.Config) chan error {
	server, err := api.NewServer(&store)
	if err != nil {
		log.Panic("cannot create http server", err)
	}
	runhttp := make(chan error)
	go func() {
		runhttp <- server.Start(config.HTTP_ADDRESS)
	}()
	return runhttp
}
