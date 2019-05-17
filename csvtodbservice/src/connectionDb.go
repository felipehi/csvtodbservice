package src

import (
	"fmt"

	_ "github.com/lib/pq" //driver Postgres
)

const (
	host     = "localhost"
	port     = 1040
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

// GetStringConnection é responsável por fazer a conexão com o banco de dados.
func GetStringConnection() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}
