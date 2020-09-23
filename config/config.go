package config


import	"database/sql"
import _"ujian/github.com/go-sql-driver/mysql"
import "log"

func GetMySQLDB()(db *sql.DB, err error)  {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""		
	dbName := "mahasiswa_enigma"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		log.Fatal(err)
	}
	return
}