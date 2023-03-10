package db

import (
	"errors"
)

var UnsupportedProtocolErr = errors.New("unsupported protocol")

func Run(dbConfig Config) (err error) {
	switch dbConfig.Protocol {
	case ProtocolMySQL:
		err = RunLinuxMySQL(dbConfig)
	case ProtocolRedis:
		err = RunLinuxRedis(dbConfig)
	case ProtocolMongoDB:
		err = RunLinuxMongoDB(dbConfig)
	default:
		return UnsupportedProtocolErr
	}

	return err
}
