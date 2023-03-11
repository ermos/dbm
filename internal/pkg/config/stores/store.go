package stores

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Store interface {
	Load(path string) (err error)
	Reload() (err error)
	Save() (err error)
}

type BaseStore struct {
	Path string `json:"-"`
}

func (c *BaseStore) Load(store any, path, name string) (err error) {
	c.Path = filepath.Join(path, fmt.Sprintf("%s.json", name))
	return c.Reload(&store)
}

func (c *BaseStore) Reload(store any) (err error) {
	var content []byte

	if content, err = os.ReadFile(c.Path); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return
	}

	return json.Unmarshal(content, &store)
}

func (c *BaseStore) Save(store any) (err error) {
	content, err := json.Marshal(&store)

	return os.WriteFile(c.Path, content, 0755)
}
