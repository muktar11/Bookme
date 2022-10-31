package dbrepo

import (
	"database/sql"
	"Bookme/internal/config"
	"Bookme/internal/repository"
)

//set  up database repository pattern it would be easier to update between databases
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

type testDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}
//super packe for postgresDBRepo and other alternative DB connnections
//pass our connection pool and AppConfig return repositroy.DatabaseRepo
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
//if we want to create MYSQL Repo 
//type MysqlRepo struct{

//}
//func NewMysqlRepo(conn *sql,DB, A *CONFIG.AppCOnfig){
//return &&NewMySqlRepo{}
//}

func NewTestingsRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}