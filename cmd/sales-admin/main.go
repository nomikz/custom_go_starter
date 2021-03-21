package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nomikz/training/internal/platform/conf"
	"github.com/nomikz/training/internal/platform/database"
	"github.com/nomikz/training/internal/schema"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	// =========================================================================
	// Get Configuration

	var cfg struct {
		DB struct {
			User       string `conf:"default:postgres"`
			Password   string `conf:"default:postgres,noprint"`
			Host       string `conf:"default:localhost:5433"`
			Name       string `conf:"default:postgres"`
			DisableTLS bool   `conf:"default:true"`
		}
		Args conf.Args
	}

	if err := conf.Parse(os.Args[1:], "sales", &cfg); err != nil {
		if err == conf.ErrHelpWanted {
			usage, err := conf.Usage("SALES", &cfg)
			if err != nil {
				return errors.Wrap(err, "generating config usage")
			}
			fmt.Println(usage)
			return nil
		}
		return errors.Wrap(err, "parsing config")
	}

	out, err := conf.String(&cfg)
	if err != nil {
		return errors.Wrap(err, "generating config for output")
	}
	log.Printf("main : Config :\n%v\n", out)

	// =========================================================================
	// Set up dependencies

	db, err := database.Open(database.Config{
		Host:       cfg.DB.Host,
		User:       cfg.DB.User,
		Password:   cfg.DB.Password,
		Name:       cfg.DB.Name,
		DisableTLS: cfg.DB.DisableTLS,
	})
	if err != nil {
		return err
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Println(err)
	}

	switch cfg.Args.Num(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			log.Fatal("applying migrations error: ", err)
		}
		log.Println("Migrations completed")
		return nil
	case "seed":
		if err := schema.Seed(db); err != nil {
			return errors.Wrap(err, "applying seeding")
		}
		log.Println("Seeding  completed")
		return nil
	}

	return nil
}
