package config

import(
	"database/sql"
	"fmt"
	"gopkg.in/gorp.v1"
	_ "github.com/go-sql-driver/mysql"
	"net/url"
	"os"
)

func convert_datasource(ds string) (result string) {
	url, _ := url.Parse(ds)
	result = fmt.Sprintf("%s@tcp(%s:3306)%s", url.User.String(), url.Host, url.Path)
	return
}

func InitDb() *gorp.DbMap {
	var datasource string

	if os.Getenv("CLEARDB_DATABASE_URL") != "" {
		datasource = convert_datasource(os.Getenv("CLEARDB_DATABASE_URL"))
	} else {
		datasource = DbUser + ":" + DbPwd + "@" + DbHost + "/" + DbTable + "?charset=utf8"
	}

	db, err := sql.Open("mysql", datasource)
	if err != nil {
		panic(err) // TODO error handle
	}

	dialect := gorp.MySQLDialect{"InnoDB", "UTF8"}
	dbmap := &gorp.DbMap{Db: db, Dialect: dialect}
	return dbmap
}
