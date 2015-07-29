package Listener

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
	"appengine"
	"io"
	"io/ioutil"
)


func NewPic(w http.ResponseWriter, req *http.Request){
	m, _ := url.ParseQuery(req.URL.RawQuery)
	c := appengine.NewContext(req)
	if m["hub.challenge"] != nil {
		io.WriteString(w, m["hub.challenge"][0])
		return
	}
	c.Infof("photo get")
	bod, _ := ioutil.ReadAll(req.Body)
	c.Infof(string(bod))
	io.WriteString(w, "Get photo")
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/getpicture", NewPic)
	http.Handle("/", r)
}
