package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang_tutorial/internal/pkg/config/envconfig"
)

type (
	Config struct {
		Driver   string
		User     string
		Pass     string
		Database string
	}
)

func LoadConfigFromEnv() Config {
	var conf Config
	envconfig.Load(&conf)
	return conf
}

func Connect() *sql.DB {
	conf := LoadConfigFromEnv()
	db, err := sql.Open(conf.Driver, conf.User+":"+conf.Pass+"@/"+conf.Database)

	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
