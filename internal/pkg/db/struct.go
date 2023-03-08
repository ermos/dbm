package db

import "time"

type Config struct {
	Alias             string
	Protocol          string
	Host              string
	Port              string
	Username          string
	PlainPassword     string `json:"-"`
	EncryptedPassword string
	DefaultDatabase   string
	LastConnectionAt  time.Time
}

const (
	ProtocolMySQL = "mysql"
	ProtocolRedis = "redis"
)

func GetProtocols() []string {
	return []string{
		ProtocolMySQL,
		ProtocolRedis,
	}
}
