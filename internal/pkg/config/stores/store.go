package stores

type Store interface {
	Load(path string) (err error)
	Reload() (err error)
	Save() (err error)
}
