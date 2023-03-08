package dbm

import (
	"encoding/json"
	"github.com/ermos/dbm/internal/pkg/goliath"
	"os"
	"path/filepath"
)

var config = &Config{}

type Config struct {
	Path           string `json:"-"`
	EncryptChecker string `json:"EncryptChecker"`
}

func Get() *Config {
	return config
}

func (c *Config) Load(path string) (err error) {
	c.Path = filepath.Join(path, "dbm.json")
	return c.Reload()
}

func (c *Config) Reload() (err error) {
	var content []byte

	if content, err = os.ReadFile(c.Path); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return
	}

	return json.Unmarshal(content, &c)
}

func (c *Config) Save() (err error) {
	content, err := json.Marshal(&c)
	if err != nil {
		return
	}

	return os.WriteFile(c.Path, content, 0755)
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
