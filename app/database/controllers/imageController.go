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

func GetImages(id int, count int, c appengine.Context) []string {
	var links []string

	//get images corresponding to our color
	listImg, _ := service.GetData(c, id, count)
	//create a table of corresponding links
	if listImg != nil {
		for _, v := range listImg {
			links = append(links, v.Link)
		}
	}
	return links
}

func AddImage(link string, col color.Color, c appengine.Context) error {
	var p color.Palette

	//get the index of the nearest color in a our color palette (256 colors)
	p = palette.Plan9
	id := p.Index(col)
	//initalize the struct to be stored
	img := models.Image{
		Date:  time.Now(),
		Color: id,
		Link:  link,
	}
	//add the struct to the database
	_, err := service.AddData(img, c)
	return err
}
