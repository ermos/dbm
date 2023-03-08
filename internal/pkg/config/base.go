package config

import (
	"github.com/ermos/dbman/internal/pkg/config/stores/credentials"
	"github.com/ermos/dbman/internal/pkg/config/stores/dbman"
)

var config = &Config{}

var dirName = ".dbman"

type Config struct {
	Path string
}

func Init() (err error) {
	if err = config.InitDir(); err != nil {
		return
	}

	config.InitStore(dbman.Get())
	config.InitStore(credentials.Get())

	return nil
}
