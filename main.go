package main

import (
	"flag"
	"time"

	"github.com/nidhey27/project-service/api"
	"github.com/nidhey27/project-service/dbs"
	"github.com/nidhey27/project-service/seeds"
	log "github.com/sirupsen/logrus"
	
)

var (
	laddr           = flag.String("addr", ":8880", "Local address for the HTTP API")
	loglevel        = flag.String("loglevel", "INFO", "Log-level (ERROR|WARN|INFO|DEBUG|TRACE)")
	initialSeedFile = flag.String("initialSeedFile", "", "Run one-time seeds passing path to a valid JSON seed file")
)

func configureLogging() error {
	l, err := log.ParseLevel(*loglevel)
	if err != nil {
		return err
	}
	log.SetLevel(l)
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat:   time.RFC3339Nano,
		DisableHTMLEscape: true,
	})
	return nil
}

func main() {
	flag.Parse()
	if err := configureLogging(); err != nil {
		log.Fatal(err)
	}

	err := dbs.InitializeDatabaseLayer()
	if err != nil {
		log.Fatal(err)
	}

	if initialSeedFile != nil && *initialSeedFile != "" {
		if err = seeds.RunSeeds(*initialSeedFile); err != nil {
			log.Fatal(err)
		}
	}

	api := api.NewRESTApiV1()
	if err = api.Serve(*laddr); err != nil {
		log.Fatal(err)
	}
}
