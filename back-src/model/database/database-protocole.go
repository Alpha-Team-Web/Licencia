package database

import "github.com/go-pg/pg/orm"

var (
	dbc = dbConnection{
		username: "postgres",
		password: "s1234567",
		/*		username: "postgres",
				password: "mbsoli1743399413",*/
	}

	options = &orm.CreateTableOptions{
		IfNotExists: true,
	}
)

type dbConnection struct {
	username string
	password string
}
