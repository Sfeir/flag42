package service

import (
	"appengine"
	"appengine/datastore"

	"database/models"
)

func AddData(img models.Image, c appengine.Context) (string, error) {
	key := datastore.NewIncompleteKey(c, "image", nil)
	a, err := datastore.Put(c, key, &img)
	return a.String(), err
}

func GetData(c appengine.Context, id int, count int) ([]models.Image, error) {
	var listImg []models.Image

	q := datastore.NewQuery("image")
	q = q.Filter("Color =", id)
	q = q.Order("-Date")
	if count != -1 {
		q = q.Limit(count)
	}
	if _, err := q.GetAll(c, &listImg); err != nil {
		return nil, err
	} else {
		return listImg, nil
	}
}
