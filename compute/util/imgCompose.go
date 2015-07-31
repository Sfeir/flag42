package main

import (
	"image"
	"image/color"
	"image/color/palette"
	_ "image/gif"
	"image/jpeg"
	_ "image/png"
	"image/draw"
	//"code.google.com/p/graphics-go/graphics"
	"strconv"
	"log"
	"encoding/json"
	"math"
	"bytes"
	"io/ioutil"
	
	"github.com/nfnt/resize"
	"net/http"
	"os"
	"strings"
)
type imgLib struct{
	Col uint8
	Tab []image.Image
}

// send image resize request with url in string and size in uint
// receive image type
func cropIn(filename string, paneau draw.Image, location image.Point){
	data,err:=ioutil.ReadFile(filename)
	if err!=nil{
		log.Fatal(err)
	}
	reader := strings.NewReader(string(data))
	img,_,_:=image.Decode(reader)
	//bounds := paneau.Bounds()
	sr:=img.Bounds()
	r:=image.Rectangle{location, location.Add(sr.Size())}
	draw.Draw(paneau, r, img, sr.Min, draw.Src)
	log.Println("cropped")
}
/*
func shuffle(uri []string) string{
	// 
}
*/
// count = le nombre d'images
func myGetImages(panneau image.Image, location image.Point, level uint8) []string{
	var p color.Palette
	p = palette.Plan9
	col:=panneau.At(location.X, location.Y)
	//return GetImages(col, -1, c)
	//return (ResizeImage(shuffle(myUri), math.Exp2(level)))
	id := p.Index(col)
	resp,_:=http.Get("https://flag-42.appspot.com/sendlinks?col="+strconv.Itoa(id)+"&count="+string(-1))
	data, _ := ioutil.ReadAll(resp.Body)
	var v []string
	json.Unmarshal(data, &v)
	log.Print("%#v", v)
	return v
}

func getColor(panneau image.Image, location image.Point) color.Color{
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

func Compose(paneau image.Image, level uint8) draw.Image{
	//var lib []imgLib
	step:=uint8(math.Exp2(float64(level)))
	bounds:=paneau.Bounds()
	paneau=resize.Resize(uint(bounds.Size().X)*uint(step),uint(bounds.Size().Y)*uint(step), paneau, resize.Bilinear)
	bounds=paneau.Bounds()
	//m:=image.NewRGBA(image.Rect(0,0,640,640))
	//draw.Draw(m,m.Bounds(), paneau, image.ZP, draw.Src)
	m:=image.NewRGBA(image.Rect(bounds.Min.X, bounds.Min.Y, bounds.Max.X, bounds.Max.Y))
	var table [256]uint
	for i:=0; i<256; i++{
		table[i]=0
	}
	var p color.Palette
	p = palette.Plan9
	for y:=bounds.Min.Y; y<bounds.Max.Y; y+=int(step){
		if y>40*int(step){break}
		for x:=bounds.Min.X; x<bounds.Max.X; x+=int(step){
			if x>40*int(step){break}
			col:=paneau.At(x, y)
			id := p.Index(col)
			if table[id]==0{
				GetnSave(uint8(id), uint(x), uint(y), level)
			}else{
				CopynSave(uint8(id), uint(x), uint(y), level)
			}
			table[id]++
			//m.Set(x, y, paneau.At(x, y))
			point:=image.Point{x,y}
			//myGetImages(paneau, point, level, uri)
			//img:=ResizeImage("https://igcdn-photos-c-a.akamaihd.net/hphotos-ak-xaf1/t51.2885-15/11380762_1476796635966754_1332771621_n.jpg", uint(step))
			cropIn("pixels/level"+strconv.Itoa(int(level))+"/"+strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+".jpg", m, point)
			//addImg(lib, img, paneau.At(x,y))
		}
	}
	return m
}

func CopynSave(id uint8, x, y uint, level uint8){
	data,err:=ioutil.ReadFile("pixels/colors/level"+strconv.Itoa(int(level))+"/"+strconv.Itoa(int(id))+".jpg")
	if err!=nil{
		log.Fatal(err)
	}
	ioutil.WriteFile("pixels/level"+strconv.Itoa(int(level))+"/"+strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+".jpg", data, 0644)
}

func GetnSave(id uint8, x, y uint, level uint8){
	//resp,_:=http.Get("https://flag-42.appspot.com/sendlinks?col="+strconv.Itoa(int(id))+"&count="+string(-1))
	resp,_:=http.Get("https://flag-42.appspot.com/sendlinks?r=10&g=10&b=255&count=2")
	data, _ := ioutil.ReadAll(resp.Body)
	var v []string
	json.Unmarshal(data, &v)
	if v==nil{
		log.Fatal("urls empty")
	}
	pixel:=ResizeImage(v[0], uint(math.Exp2(float64(level))))
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, pixel, nil)
	send:=buf.Bytes()
	ioutil.WriteFile("pixels/colors/level"+strconv.Itoa(int(level))+"/"+strconv.Itoa(int(id))+".jpg", send, 0644)
	ioutil.WriteFile("pixels/level"+strconv.Itoa(int(level))+"/"+strconv.Itoa(int(x))+"_"+strconv.Itoa(int(y))+".jpg", send, 0644)
}

func initiate(){
	level:=0
	//step:=uint8(math.Exp2(float64(level)))
	os.Mkdir("pixels", 0777)
	os.Mkdir("pixels/colors", 0777)
	img:=ResizeImage("https://igcdn-photos-c-a.akamaihd.net/hphotos-ak-xaf1/t51.2885-15/11380762_1476796635966754_1332771621_n.jpg", 640)
	var paneau draw.Image
	for level <= 6 {
		os.Mkdir("pixels/level"+strconv.Itoa(level), 0777)
		os.Mkdir("pixels/colors/level"+strconv.Itoa(int(level)), 0777)
		paneau=Compose(img, uint8(level))
		/*
		img:=ResizeImage("https://igcdn-photos-c-a.akamaihd.net/hphotos-ak-xaf1/t51.2885-15/11380762_1476796635966754_1332771621_n.jpg", 640*uint(step))
		paneau:=Compose(img, uint8(level), uri)*/
		buf := new(bytes.Buffer)
		jpeg.Encode(buf, paneau, nil)
		send:=buf.Bytes()
		ioutil.WriteFile("/home/dnguyen/Images/panel"+strconv.Itoa(level)+".jpg", send, 0644)
		level++
		//step=uint8(math.Exp2(float64(level)))
	}
	log.Printf("%#v", paneau)
}
func main(){
	initiate()
}