package conn

import (
	"fmt"
	"log"

	"github.com/imrancluster/goboiler/config"
	"github.com/jinzhu/gorm"

	// postgres conn
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PostgresClient : a postgres connection instance type
type PostgresClient struct {
	*gorm.DB
}

// DbConn : postgres connection instance
var DbConn PostgresClient

// ConnectDB : exportable db connection for app bootstrap
func ConnectDB() error {
	cfg := config.DB()
	cfgApp := config.App()
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.DatabaseName,
		cfg.Password,
	)
	con, err := gorm.Open("postgres", connString)
	if err != nil {
		log.Fatalf("cannot connect to database. reason: %s", err.Error())
		return err
	}

	if cfgApp.Env != "prod" {
		con.LogMode(true)
	}

	DbConn = PostgresClient{
		DB: con,
	}

	return nil
}

// PostgresDB : expotable db query object
func PostgresDB() *PostgresClient {
	return &DbConn
}
