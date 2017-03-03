package models

import (
	"database/sql"
	"fmt"

	"github.com/deep/GOMVC/models/mafiaDB"
	"github.com/go-gorp/gorp"
)

//lol
//lol
func InitializeGoDB() Dbmap {
	db, err := sql.Open("mymysql", "tcp:localhost:3306*mafiDB/root/Anno2070")

	if err != nil {
		fmt.Println("lol")
	}
	// construct a gorp DbMap
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	dbmap.AddTable(mafiaDB.User{}).SetKeys(true, "Id")
	dbmap.AddTable(mafiaDB.CharacterType{}).SetKeys(true, "Id")
	dbmap.AddTable(mafiaDB.HostedGames{}).SetKeys(true, "Id")
	dbmap.AddTable(mafiaDB.Round{}).SetKeys(true, "Id")
	dbmap.AddTable(mafiaDB.JoinedPlayers{}).SetKeys(true, "Id")
	dbmap.AddTable(mafiaDB.GameData{}).SetKeys(true, "Id")

	return dbmap
}
