package database

import "github.com/go-pg/pg/orm"

var (
	dbc = dbConnection{
		/*		username: "ashka",
				password: "a124578",*/
		username: "postgres",
		password: "mbsoli1743399413",
		/*username: "postgres",
		password: "s1234567",*/
	}

	options = &orm.CreateTableOptions{
		IfNotExists: true,
	}
)

type dbConnection struct {
	username string
	password string
}
