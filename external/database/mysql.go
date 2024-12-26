package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yudhisrana/boilerplate-goyam/internal/config"
	"github.com/yudhisrana/boilerplate-goyam/internal/handlerErrors"
)

func ConnectDB(cfg config.DBConfig) (conn *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, handlerErrors.ErrDataSourceName
	}

	db.SetMaxIdleConns(cfg.DBConnection.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.DBConnection.MaxOpenConnections)
	db.SetConnMaxIdleTime(time.Second * time.Duration(cfg.DBConnection.ConnMaxIdleTime))
	db.SetConnMaxLifetime(time.Second * time.Duration(cfg.DBConnection.ConnMaxLifetime))

	if err := db.Ping(); err != nil {
		return nil, handlerErrors.ErrPingDatabase
	}

	fmt.Printf("success connect to database\n")
	return db, nil
}
