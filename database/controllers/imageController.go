package controllers

import (
	"appengine"
	"image/color"
	"image/color/palette"
	"time"

	"database/models"
	"database/service"
)

var img models.Image

func GetImages(col color.Color, count int, c appengine.Context) []string {
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

func AddImage(link string, col color.Color, c appengine.Context) error {
	var p color.Palette

	p = palette.WebSafe
	id := p.Index(col)
	img := models.Image{
		Date:  time.Now(),
		Color: id,
		Link:  link,
	}
	_, err := service.AddData(img, c)
	return err
}
