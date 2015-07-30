package resize

import (
	"github.com/nfnt/resize"

	"image"
	"net/http"
)

func ResizeImage(link string, size uint) image.Image {
	//getting the image from the url
	resp, _ := http.Get(link)
	defer resp.Body.Close()
	//decode the image to get an image.Image data
	img, _, _ := image.Decode(resp.Body)
	//resizing the image
	img = resize.Resize(size, size, img, resize.Bilinear)
	return img
}
