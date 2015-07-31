package util

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	//"code.google.com/p/graphics-go/graphics"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	//	"math"
	"strconv"

	"github.com/nfnt/resize"
	"net/http"
	"os"
	"strings"
)

type imgLib struct {
	Col uint8
	Tab []image.Image
}

// send image resize request with url in string and size in uint
// receive image type
func cropIn(filename string, paneau draw.Image, location image.Point) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	reader := strings.NewReader(string(data))
	img, _, _ := image.Decode(reader)
	//bounds := paneau.Bounds()
	sr := img.Bounds()
	r := image.Rectangle{location, location.Add(sr.Size())}
	draw.Draw(paneau, r, img, sr.Min, draw.Src)
	log.Println("cropped")
}

/*
func shuffle(uri []string) string{
	//
}
*/
// count = le nombre d'images
func myGetImages(panneau image.Image, location image.Point, level uint8) []string {
	var p color.Palette
	p = palette.Plan9
	col := panneau.At(location.X, location.Y)
	//return GetImages(col, -1, c)
	//return (ResizeImage(shuffle(myUri), math.Exp2(level)))
	id := p.Index(col)
	resp, _ := http.Get("https://flag-42.appspot.com/sendlinks?col=" + strconv.Itoa(id) + "&count=" + string(-1))
	data, _ := ioutil.ReadAll(resp.Body)
	var v []string
	json.Unmarshal(data, &v)
	log.Print("%#v", v)
	return v
}

func getColor(panneau image.Image, location image.Point) color.Color {
	return (panneau.At(location.X, location.Y))
}

/*
func addImg(lib []imgLib, img image.Image, color.Color){
	if lib[]
}
*/
func ResizeImage(link string, size uint) image.Image {
	//getting the image from the url
	resp, _ := http.Get(link)
	//defer resp.Body.Close()
	//decode the image to get an image.Image data
	img, _, _ := image.Decode(resp.Body)
	//resizing the image
	img = resize.Resize(size, size, img, resize.Bilinear)
	return img
}

func Compose(url string) {
	level := 0
	paneau := ResizeImage(url, 640)
	bounds := paneau.Bounds()
	var table [256]uint
	var p color.Palette
	p = palette.Plan9
	for level <= 6 {
		os.Mkdir(strconv.Itoa(int(level)), 0777)
		os.Mkdir("colors/"+strconv.Itoa(int(level)), 0777)
		//step:=uint8(math.Exp2(float64(level)))
		if level == 0 {
			m := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Min.X+64, bounds.Min.Y+64))
			//p:=image.Point{bounds.Min.X+64,bounds.Min.Y+64}
			/*if table[id]==0{
				SaveImg(m, uint(x), uint(y), level)
			}else{
				CopynSave(uint8(id), uint(x), uint(y), level)
			}
			table[id]++*/
			sr := m.Bounds()
			location := image.Point{bounds.Min.X, bounds.Min.Y}
			r2 := image.Rectangle{location, location.Add(sr.Size())}
			//r:=bounds.Sub(bounds.Min).Add(p)
			draw.Draw(m, r2, paneau, bounds.Min, draw.Src)
			SaveImg(m, 0, 0, uint8(level))
			return
		} /*else if level>=1 && level <=3{
			nbimg:=uint(math.Exp2(math.Exp2(float64(level))))
			log.Printf("nbimg=%d\n",nbimg)
			m := image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Min.X+64, bounds.Min.Y+64))
			for x:=0;uint8(x)<step;x++{
				for y:=0;uint8(y)<step;y++{
					p:=image.Point{bounds.Min.X+64/int(step),bounds.Min.Y+64/int(step)}
					liens:=myGetImages(paneau, p, level)
					img:=ResizeImage(liens[0], uint(64/step))

				}
			}
		}else{
			for i:=0; i<256; i++{
				table[i]=0
			}
		}*/
		level++
	}
	log.Printf("%v%v%v\n", p, table)
	/*	for y:=bounds.Min.Y; y<bounds.Max.Y; y++{
			//if y>10*int(step){break}
			for x:=bounds.Min.X; x<bounds.Max.X; x++{
				//if x>10*int(step){break}
				col:=paneau.At(x, y)
				id := p.Index(col)
				if table[id]==0{
					GetnSave(uint8(id), uint(x), uint(y), level)
				}else{
					CopynSave(uint8(id), uint(x), uint(y), level)
				}
				table[id]++
				//m.Set(x, y, paneau.At(x, y))
				//point:=image.Point{x,y}
				//myGetImages(paneau, point, level, uri)
				//img:=ResizeImage("https://igcdn-photos-c-a.akamaihd.net/hphotos-ak-xaf1/t51.2885-15/11380762_1476796635966754_1332771621_n.jpg", uint(step))
				//cropIn("pixels/level"+strconv.Itoa(int(level))+"/"+strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+".jpg", m, point)
				//addImg(lib, img, paneau.At(x,y))
			}
		}
	*/
}

func CopynSave(id uint8, x, y uint, level uint8) {
	data, err := ioutil.ReadFile("colors/" + strconv.Itoa(int(level)) + "/" + strconv.Itoa(int(id)) + ".jpg")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(strconv.Itoa(int(level))+"/"+strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+".jpg", data, 0644)
}

func SaveImg(img image.Image, x, y uint, level uint8) {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, img, nil)
	send := buf.Bytes()
	ioutil.WriteFile(strconv.Itoa(int(level))+"/"+strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+".jpg", send, 0644)
}

// func GetnSave(id uint8, x, y uint, level uint8) {
// 	//resp,_:=http.Get("https://flag-42.appspot.com/sendlinks?col="+strconv.Itoa(int(id))+"&count="+string(-1))
// 	resp, _ := http.Get("https://flag-42.appspot.com/sendlinks?r=10&g=10&b=255&count=2")
// 	data, _ := ioutil.ReadAll(resp.Body)
// 	var v []string
// 	json.Unmarshal(data, &v)
// 	if v == nil {
// 		log.Fatal("urls empty")
// 	}
// 		pixel := ResizeImage(v[0], uint(math.Exp2(float64(level))))
// 	buf := new(bytes.Buffer)
// 	jpeg.Encode(buf, pixel, nil)
// 	send := buf.Bytes()
// 	ioutil.WriteFile("colors/"+strconv.Itoa(int(level))+"/"+strconv.Itoa(int(id))+".jpg", send, 0644)
// 	ioutil.WriteFile(strconv.Itoa(int(level))+"/"+strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+".jpg", send, 0644)
// }

/*
func main(){
	Compose("https://igcdn-photos-c-a.akamaihd.net/hphotos-ak-xaf1/t51.2885-15/11380762_1476796635966754_1332771621_n.jpg")
}*/
