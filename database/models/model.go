package models

import (
	"time"
)

type Image struct {
	Date  time.Time
	Color int
	Link  string
}
