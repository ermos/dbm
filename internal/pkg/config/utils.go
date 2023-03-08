package config

import (
	"github.com/ermos/dbman/internal/pkg/config/stores"
	"github.com/mitchellh/go-homedir"
	"os"
	"path/filepath"
)

func (c *Config) InitDir() (err error) {
	home, err := homedir.Dir()
	if err != nil {
		return
	}

	c.Path = filepath.Join(home, dirName)

	return os.MkdirAll(config.Path, 0755)
}

func (c *Config) InitStore(store stores.Store) {
	if err := store.Load(c.Path); err != nil {
		return
	}
	if err := store.Save(); err != nil {
		return
	}
}
