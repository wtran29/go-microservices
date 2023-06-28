package main

import (
	"database/sql"

	"github.com/wtran29/go-services/auth/data"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
