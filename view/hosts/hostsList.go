package hosts

import (
	"fmt"
	"net/http"

	"github.com/yosssi/ace"
)

func ViewHostsList(w http.ResponseWriter, r *http.Request, conf map[string]interface{}) {
	hosts, ok := conf["databaseHosts"].([]string)
	if !ok {
		http.Error(w, "no hosts", http.StatusInternalServerError)
		return
	}
	fmt.Println(hosts)
	data := map[string]interface{}{
		"DBHosts": hosts,
	}
	tpl, err := ace.Load("template/hosts/hostsList", "", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
