package credentials

import (
	"errors"
	"github.com/ermos/dbm/internal/pkg/auth"
	"github.com/ermos/dbm/internal/pkg/config/stores"
	"github.com/ermos/dbm/internal/pkg/db"
	"github.com/ermos/dbm/internal/pkg/goliath"
	"time"
)

var config = &Config{
	BaseStore:   &stores.BaseStore{},
	Credentials: make(map[string]db.Config),
}

type Config struct {
	*stores.BaseStore
	Credentials map[string]db.Config
}

func Get() *Config {
	return config
}

func (c *Config) Load(path string) error {
	return c.BaseStore.Load(&c, path, "credentials")
}

func (c *Config) Reload() error {
	return c.BaseStore.Reload(&c)
}

func (c *Config) Save() error {
	return c.BaseStore.Save(&c)
}

func (c *Config) Add(dbConfig db.Config) (err error) {
	if dbConfig.PlainPassword != "" {
		dbConfig.EncryptedPassword, err = goliath.EncryptData([]byte(dbConfig.PlainPassword), auth.String())
		if err != nil {
			return
		}
	}

	c.Credentials[dbConfig.Alias] = dbConfig

	return c.Save()
}

func (c *Config) GetAlias(alias string) (dbConfig db.Config, err error) {
	dbConfig = c.Credentials[alias]
	if dbConfig.Alias == "" {
		err = errors.New("alias not found")
		return
	}

	if dbConfig.EncryptedPassword != "" {
		var plain []byte

		plain, err = goliath.DecryptData(dbConfig.EncryptedPassword, auth.String())
		if err != nil {
			return
		}

		dbConfig.PlainPassword = string(plain)
	}

	return
}

func (c *Config) RemoveAlias(alias string) error {
	dbConfig := c.Credentials[alias]
	if dbConfig.Alias == "" {
		return errors.New("alias not found")
	}

	delete(c.Credentials, alias)

	return c.Save()
}

func (c *Config) UpdateLastConnection(alias string) error {
	dbConfig := c.Credentials[alias]
	if dbConfig.Alias == "" {
		return errors.New("alias not found")
	}

	dbConfig.LastConnectionAt = time.Now()
	c.Credentials[alias] = dbConfig

	return c.Save()
}
