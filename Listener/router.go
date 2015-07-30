package Listener

import (
	"appengine"
	"appengine/urlfetch"
	"encoding/json"
	"github.com/carbocation/go-instagram/instagram"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"util"

	"database/controllers"
)

var titime int64

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
		var tempTime int64
		for _, v := range res {
			if v.CreatedTime > titime {
				if tempTime == 0 || tempTime < v.CreatedTime {
					tempTime = v.CreatedTime
				}
				url := v.Images.StandardResolution.URL
				rgb, err := util.PixColor(url, c)
				if err != nil {
					json.NewEncoder(w).Encode(err.Error())
				}
				err = controllers.AddImage(url, rgb, c)
				if err != nil {
					json.NewEncoder(w).Encode(err.Error())
				}
			}
		}
		titime = tempTime
	} else {
		json.NewEncoder(w).Encode(err.Error())
	}
}

func init() {
	r := mux.NewRouter()
	r.HandleFunc("/getpicture", NewPic)
	http.Handle("/", r)
}
