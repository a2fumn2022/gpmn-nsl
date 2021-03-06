// The psql package handles all interactions with the database.
package psql

import (
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"github.com/jackc/pgx"
)

// DB holds a live connection to the database.
// DB functions use the connection to query the database.
type DB struct {
	connection *pgx.Conn
}

// Init establishes a reusable connection.
func (db *DB) Init(config datastore.DatabaseConfig) error {
	connConfig := pgx.ConnConfig{
		User:              config.User,
		Password:          config.Password,
		Host:              config.Host,
		Port:              uint16(config.Port),
		Database:          config.Dbname,
		TLSConfig:         nil,
		UseFallbackTLS:    false,
		FallbackTLSConfig: nil,
	}
	conn, err := pgx.Connect(connConfig)
	if err != nil {
		return err
	}
	db.connection = conn
	return nil
}
