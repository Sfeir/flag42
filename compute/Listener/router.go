package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/url"

	"storage"
	"util"
)

func storeImages(w http.ResponseWriter, req *http.Request) {
	m, _ := url.ParseQuery(req.URL.RawQuery)
	dir := m["dir"][0]
	store := m["storage"][0]
	storage.ComputeToStorage(dir, store)
}

func ComposeFunc(w http.ResponseWriter, req *http.Request) {
	util.Compose("https://igcdn-photos-c-a.akamaihd.net/hphotos-ak-xaf1/t51.2885-15/11380762_1476796635966754_1332771621_n.jpg")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/store", storeImages)
	r.HandleFunc("/compose", ComposeFunc)
	http.Handle("/", r)
}
