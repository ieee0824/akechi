package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ieee0824/akechi/api"
)

func getConnectionString(userName, password, host, port, dbname string) string {
	if port == "" {
		port = "3306"
	}
	if host == "" {
		host = "localhost"
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		userName, password, host, port, dbname,
	)
}

//GetDBList -> Listup database list
// ignores
// * information_schema
// * mysql
// * performance_schema
// GetDBList is returned json
func getDBList(userName, password, host, port string) (api.JSON, error) {
	ignores := map[string]bool{
		"information_schema": true,
		"mysql":              true,
		"performance_schema": true,
	}
	connectionString := getConnectionString(userName, password, host, port, "")
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("show databases")
	if err != nil {
		return nil, err
	}
	var dbName []string
	for rows.Next() {
		var name string
		rows.Scan(&name)
		if !ignores[name] {
			dbName = append(dbName, name)
		}
	}

	return json.Marshal(dbName)
}

func APIGetDBList(w http.ResponseWriter, r *http.Request, conf map[string]interface{}) {
	host := r.FormValue("DBHost")

	dbConf, ok := conf[host].(map[string]string)
	if !ok {
		w.WriteHeader(503)
		w.Write([]byte(fmt.Sprintf("\"%s\" not found DB configure.\n", host)))
		return
	}

	user := dbConf["user"]
	password := dbConf["password"]
	port := dbConf["port"]

	list, err := getDBList(user, password, host, port)
	if err != nil {
		w.WriteHeader(503)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(list)
}
