package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ieee0824/akechi/api/db"
)

var conf = map[string]interface{}{}
var dbConf = map[string]string{
	"user":     "root",
	"password": "",
	"port":     "",
	"host":     "localhost",
}

func init() {
	conf[dbConf["host"]] = dbConf
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It works."))
	})
	http.HandleFunc("/api/getDBList", func(w http.ResponseWriter, r *http.Request) {
		db.APIGetDBList(w, r, conf)
	})

	http.ListenAndServe(":10000", nil)
}
