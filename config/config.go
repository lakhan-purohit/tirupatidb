package config

import (
	"encoding/json"
	"io/ioutil"
	"tirupatiBE/dal/dbModel"
	"tirupatiBE/router"
)

type Config struct {
	Db         dbModel.DB
	HttpServer router.HttpServer
}

func ReadConfig(fileName *string) (*Config, error) {

	var conf Config
	fileValue, err := ioutil.ReadFile(*fileName)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(fileValue, &conf)

	return &conf, nil
}
