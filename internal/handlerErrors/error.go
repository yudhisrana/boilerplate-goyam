package handlerErrors

import "errors"

var (
	// handle config error
	ErrReadFileConfig  = errors.New("failed to read config! file yaml not found")
	ErrUnmarshalConfig = errors.New("failed to unmarshal config file")

	// handle database error
	ErrDataSourceName = errors.New("data source name invalid")
	ErrPingDatabase   = errors.New("not connected, failed to ping database")
)
