package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SQL struct {
	Db             *sqlx.DB
	DataSourceName string
}

func (s *SQL) Connect() {
	db, err := sqlx.Connect("mysql", s.DataSourceName)
	if err != nil {
		log.Fatal("Fail connect to MySQL")
		return
	}
	s.Db = db
	if err := s.Db.Ping(); err != nil {
		log.Fatal("Fail connect to MySQL")
		return
	}
	log.Println("Connect Success")
}

func (s *SQL) Close() {
	s.Db.Close()
}
