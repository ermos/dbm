package credentials

import (
	"encoding/json"
	"errors"
	"github.com/ermos/dbm/internal/pkg/auth"
	"github.com/ermos/dbm/internal/pkg/db"
	"github.com/ermos/dbm/internal/pkg/goliath"
	"os"
	"path/filepath"
)

var config = &Config{
	Credentials: make(map[string]db.Config),
}

type Config struct {
	Path        string `json:"-"`
	Credentials map[string]db.Config
}

func Get() *Config {
	return config
}

func (c *Config) Load(path string) (err error) {
	c.Path = filepath.Join(path, "credentials.json")
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

func (c *Config) Get(alias string) (dbConfig db.Config, err error) {
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
