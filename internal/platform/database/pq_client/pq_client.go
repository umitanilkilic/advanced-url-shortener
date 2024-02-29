package postgresql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToPostgres(pqAddress string, pqPort string, pqPassword string, pqUser string, pqDatabase string, pqSSLMode string) error {
	connectionTxt := fmt.Sprintf("user=%v dbname=%v sslmode=%v password=%v host=%v", pqUser, pqDatabase, pqSSLMode, pqPassword, pqAddress+":"+pqPort)
	db, err := sqlx.Connect("postgres", connectionTxt)
	if err != nil {
		return fmt.Errorf("could not connect to postgres: %v", err)
	}
	db.Close()
	return nil
}
