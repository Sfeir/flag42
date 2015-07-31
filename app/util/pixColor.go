//package main
package util

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"strconv"
	//"github.com/nfnt/resize"
	"log"
)

// converts a 8 bits integer to hexadecimal in string format
func int2hex(i uint8) string {
	return string(fmt.Sprintf("%02x", i))
}

// converts a 2 characters hexadecimal to a 8 bits integer
func hex28bit(hex string) uint8 {
	result := 0
	for i := 0; i < 2; i++ {
		switch hex[i] {
		case 'a':
			result += 10
			break
		case 'b':
			result += 11
			break
		case 'c':
			result += 12
			break
		case 'd':
			result += 13
			break
		case 'e':
			result += 14
			break
		case 'f':
			result += 15
			break
		default:
			tmp, _ := strconv.Atoi(string(hex[i]))
			result += tmp
			break
		}
		if i == 0 {
			result *= 16
		}
	}
	return uint8(result)
}

// converts the whole hexadecimal of 24 bits to 3 integers of 8 bits each
// defining the red, green and blue color
func hex2int(h string) (uint8, uint8, uint8) {
	sr := string(h[2])+string(h[3])
	sg := string(h[4])+string(h[5])
	sb := string(h[6])+string(h[7])
	/*for i := 0; i < len(h); i++ {
		if 2 <= i && i <= 3 {
			sr += string(h[i])
		} else if 4 <= i && i <= 5 {
			sg += string(h[i])
		} else if 6 <= i && i <= 7 {
			sb += string(h[i])
		}
	}*/
	r := hex28bit(sr)
	g := hex28bit(sg)
	b := hex28bit(sb)
	return uint8(r), uint8(g), uint8(b)
}

// takes an image url and returns the dominant color
// of type color.RGBA
func PixColor(url string, c appengine.Context) (color.Color, error) {
	client := urlfetch.Client(c)
	resp, errh := client.Get(url)
	if errh != nil {
		return nil, errh
	}
	m,_,err := image.Decode(resp.Body)
	if err != nil {
		log.Printf("StatusCode=%d",resp.StatusCode)
		return nil, err
	}
	blocksize := 4
	tableau := make(map[string]int)
	rgb := color.RGBA{R: 0, G: 0, B: 0, A: 0xff}
	//m:=resize.Resize(64,64,r, resize.Bilinear)
	bounds := m.Bounds()
	dominantValue:=0
	dominantColor:=""
	for j := bounds.Min.X; j < bounds.Max.X; j += blocksize {
		for i := bounds.Min.Y; i < bounds.Max.Y; i += blocksize {
			couleur := m.At(j, i)
			r, g, b, _ := couleur.RGBA()
			hexa := "0x" + int2hex(uint8(r)) + int2hex(uint8(g)) + int2hex(uint8(b))
			tableau[hexa]++
			// saves the dominant color if it is one
			if tableau[hexa]>dominantValue{
				dominantValue=tableau[hexa]
				dominantColor=hexa
			}
		}
	}
	result:=dominantColor
	rgb.R, rgb.G, rgb.B = hex2int(result)
	log.Printf("%#v\n", rgb)
	return rgb, nil
}
