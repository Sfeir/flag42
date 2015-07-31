package Listener

import (
	//"appengine"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"

	"storage"
)

func storeImages(w http.ResponseWriter, req *http.Request) {
	m, _ := url.ParseQuery(req.URL.RawQuery)
	dir := m["dir"][0]
	store := m["storage"][0]
	storage.ComputeToStorage(dir, store)
	// c := appengine.NewContext(req)
	// c.Infof("%d", len(str))
	// io.WriteString(w, string(json))
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/store", storeImages)
	r.HandleFunc("/compose", util/Compose)
	http.Handle("/", r)
}
