package main

import (
	"flag"
	_ "github.com/lib/pq"
	"github.com/nomikz/training/internal/platform/database"
	"github.com/nomikz/training/internal/schema"
	"log"
)

func main() {

	// =========================================================================
	// Set up dependencies

	db, err := database.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Println(err)
	}


	flag.Parse()

	switch flag.Arg(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			log.Fatal("applying migrations error: ", err)
		}
		log.Println("Migrations completed")
		return
	case "seed":
		if err := schema.Seed(db); err != nil {
			log.Fatal("applying seeding error: ", err)
		}
		log.Println("Seeding  completed")
		return
	}
}


