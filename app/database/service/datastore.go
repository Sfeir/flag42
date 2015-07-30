package service

import (
	"appengine"
	"appengine/datastore"

	"database/models"
)

func AddData(img models.Image, c appengine.Context) (string, error) {
	//create a new key for our data
	key := datastore.NewIncompleteKey(c, "image", nil)
	//store in the datastore
	a, err := datastore.Put(c, key, &img)
	return a.String(), err
}

func GetData(c appengine.Context, id int, count int) ([]models.Image, error) {
	var listImg []models.Image

	//create a query for the datastore
	q := datastore.NewQuery("image")
	//filter by choosen color
	q = q.Filter("Color =", id)
	//order by most recently added image
	q = q.Order("-Date")
	//limit to number of needed images (if -1 get all images)
	if count != -1 {
		q = q.Limit(count)
	}
	//get data from the datastore
	if _, err := q.GetAll(c, &listImg); err != nil {
		return nil, err
	} else {
		return listImg, nil
	}
}
