package dbm

import (
	"github.com/ermos/dbm/internal/pkg/config/stores"
	"github.com/ermos/dbm/internal/pkg/goliath"
)

var config = &Config{
	BaseStore: &stores.BaseStore{},
}

type Config struct {
	*stores.BaseStore
	EncryptChecker string `json:"EncryptChecker"`
}

func Get() *Config {
	return config
}

func (c *Config) Load(path string) error {
	return c.BaseStore.Load(&c, path, "dbm")
}

func (c *Config) Reload() error {
	return c.BaseStore.Reload(&c)
}

func (c *Config) Save() error {
	return c.BaseStore.Save(&c)
}

func (c *Config) GenerateEncryptChecker(password string) error {
	result, err := goliath.EncryptData([]byte("dbm"), password)
	if err != nil {
		return err
	}

	c.EncryptChecker = result

	return c.Save()
}

func (c *Config) IsValidMasterPassword(password string) bool {
	result, err := goliath.DecryptData(c.EncryptChecker, password)
	return err == nil && string(result) == "dbm"
}
