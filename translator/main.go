package main

import (
	"database/sql"
	"fmt"
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
	conn, err := sql.Open(config.DB_DRIVER, config.DB_SOURCE_TRANSLATOR)
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
	server, err := api.NewServer(&store, config)
	if err != nil {
		log.Panic("cannot create https server", err)
	}
	runhttps := make(chan error)
	go func() {
		runhttps <- server.StartTls(config.HTTPS_ADDRESS_TRANSLATOR)
	}()
	return runhttps
}

func RunHttp(store db.Store, config config.Config) chan error {
	server, err := api.NewServer(&store, config)
	if err != nil {
		log.Panic("cannot create http server", err)
	}
	runhttp := make(chan error)
	go func() {
		runhttp <- server.Start(config.HTTP_ADDRESS_TRANSLATOR)
	}()
	return runhttp
}
