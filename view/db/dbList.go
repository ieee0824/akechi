package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/yosssi/ace"
)

var client = &http.Client{}

func ViewDBList(w http.ResponseWriter, r *http.Request, conf map[string]interface{}) {
	host := r.FormValue("DBHost")
	value := url.Values{}
	value.Add("DBHost", host)
	upConf, ok := conf["upConf"].(map[string]string)
	if !ok {
		http.Error(w, "Incorrect setting.", http.StatusInternalServerError)
		return
	}

	resp, err := http.PostForm(fmt.Sprintf("http://localhost:%s/api/getDBList", upConf["port"]), value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dbList []string

	err = json.Unmarshal(bin, &dbList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"DBs": dbList,
	}

	tpl, err := ace.Load("template/db/dbList", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
