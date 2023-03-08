package config

import (
	"github.com/ermos/dbm/internal/pkg/config/stores/credentials"
	"github.com/ermos/dbm/internal/pkg/config/stores/dbm"
)

var config = &Config{}

var dirName = ".dbm"

type Config struct {
	Path string
}

func Init() (err error) {
	if err = config.InitDir(); err != nil {
		return
	}

	config.InitStore(dbm.Get())
	config.InitStore(credentials.Get())

	return nil
}
