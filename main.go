package main

import (
	"flag"
	"tirupatiBE/config"
	"tirupatiBE/dal/dbSql"
	"tirupatiBE/router"
)

func main() {
	configFileName := flag.String("config", "config.json", "config file")
	flag.Parse()

	if *configFileName == "" {
		panic("Something goes wrong")

	}

	conf, errFile := config.ReadConfig(configFileName)
	if errFile != nil {
		panic("Error in to read the file")
	}

	db, err := dbSql.SqlConnection(conf)

	if err != nil {
		panic("Error in the  connection to make DB ")
	}

	rtr := router.CreateServer(db, router.HttpServer(conf.HttpServer))
	rtr.Begin()

}
