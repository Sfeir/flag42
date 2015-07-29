package controllers

import (
	"appengine"
	"image/color"
	"image/color/palette"
	"net/http"
	"time"

	"models"
	"service"
)

var img models.Image

func GetImages(col color.Color, count int, req *http.Request) []string {
	c := appengine.NewContext(req)
	var links []string
	var p color.Palette

	p = palette.WebSafe
	id := p.Index(col)
	listImg, _ := service.GetData(c, id, count)
	if listImg != nil {
		for _, v := range listImg {
			links = append(links, v.Link)
		}
	}
	return links
}

func AddImage(link string, col color.Color, req *http.Request) {
	c := appengine.NewContext(req)
	var p color.Palette

	p = palette.WebSafe
	id := p.Index(col)
	img := models.Image{
		Date:  time.Now(),
		Color: id,
		Link:  link,
	}
	service.AddData(img, c)
}
