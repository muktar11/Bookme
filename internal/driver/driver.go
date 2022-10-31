package driver

import (
	"database/sql"
	"fmt"

	"log"
	"time"
	_"github.com/jackc/pgconn"
	_"github.com/jackc/pgx/v4/stdlib"
	_"github.com/jackc/pgconn"
)

//DB holds the databse connection information
type DB struct{
	SQL *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 25 
const maxIdleDbConn = 25 
const maxDbLifetime = 5 *time.Minute 

//Connect Postgres creates databse pool for postgres 
func ConnectPostgres(dsn string) (*DB, error){
	d, err := sql.Open("pgx", dsn)
	if err != nil{
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxIdleTime(maxDbLifetime)
	dbConn.SQL = d 

	err = testDB(err, d)
	return dbConn, err
}

//testDB pings datbase 
func testDB(err error, d *sql.DB) error {
	err = d.Ping()
	if err != nil{
		fmt.Println("Error!", err)
	} else{
		log.Println("***Pinged postgres database successfully!")
	}
	return err 
}
