package Listener

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"appengine/urlfetch"
	"database/controllers"
	"database/models"
	"encoding/json"
	"github.com/carbocation/go-instagram/instagram"
	"github.com/gorilla/mux"
	"image/color"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"util"
)

func NewPic(w http.ResponseWriter, req *http.Request) {
	m, _ := url.ParseQuery(req.URL.RawQuery)
	c := appengine.NewContext(req)
	if m["hub.challenge"] != nil {
		io.WriteString(w, m["hub.challenge"][0])
		return
	}
	bod, _ := ioutil.ReadAll(req.Body)
	var resp []instagram.RealtimeResponse
	err := json.Unmarshal(bod, &resp)
	//	w.Header().Set("content-type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	cli := urlfetch.Client(c)
	client := instagram.NewClient(cli)
	client.ClientID = "3f2ec4223de64513b58f6a2764083a66"
	client.ClientSecret = "fb8a88e4743144b39b011a402a4a66c2"
	res, _, err := client.Tags.RecentMedia(resp[0].ObjectID, nil)
	if err == nil {
		i := 0
		for true {
			v := res[i]
			i++
			item, err := memcache.Get(c, "last")
			url := v.Images.StandardResolution.URL
			if err == nil && string(item.Value) == url {
				c.Infof("Cache %d", i)
				break
			}
			q := datastore.NewQuery("image").Filter("Link =", url)
			var link []models.Image
			q.GetAll(c, &link)
			if len(link) != 0 {
				c.Infof("Data %d", i)
				break
			}
			rgb, err := util.PixColor(url, c)
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			}
			err = controllers.AddImage(url, rgb, c)
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			}
		}
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}
	cache := &memcache.Item{
		Key:   "last",
		Value: []byte(res[0].Images.StandardResolution.URL),
	}
	if err := memcache.Add(c, cache); err != nil {
		memcache.Set(c, cache)
	}
}

func sendLinks(w http.ResponseWriter, req *http.Request) {
	m, _ := url.ParseQuery(req.URL.RawQuery)
	col, _ := strconv.Atoi(m["col"][0])
	count, _ := strconv.Atoi(m["count"][0])
	str := controllers.GetImages(col, count, appengine.NewContext(req))
	json, _ := json.Marshal(str)
	c := appengine.NewContext(req)
	c.Infof("%d", len(str))
	io.WriteString(w, string(json))
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/getpicture", NewPic)
	r.HandleFunc("/sendlinks", sendLinks)
	http.Handle("/", r)
}
