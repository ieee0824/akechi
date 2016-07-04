package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	apidb "github.com/ieee0824/akechi/api/db"
	"github.com/ieee0824/akechi/view/db"
	"github.com/ieee0824/akechi/view/hosts"
)

var conf = map[string]interface{}{}
var dbConf = map[string]string{
	"user":     "root",
	"password": "",
	"port":     "",
	"host":     "localhost",
}

var databaseHosts = []string{"localhost"}

var upConf = map[string]string{
	"port": "10000",
}

func init() {
	conf[dbConf["host"]] = dbConf
	conf["upConf"] = upConf
	conf["databaseHosts"] = databaseHosts
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It works."))
	})
	http.HandleFunc("/api/getDBList", func(w http.ResponseWriter, r *http.Request) {
		apidb.APIGetDBList(w, r, conf)
	})
	http.HandleFunc("/hostsList", func(w http.ResponseWriter, r *http.Request) {
		hosts.ViewHostsList(w, r, conf)
	})
	http.HandleFunc("/dbList", func(w http.ResponseWriter, r *http.Request) {
		db.ViewDBList(w, r, conf)
	})

	http.ListenAndServe(fmt.Sprintf(":%s", upConf["port"]), nil)
}
