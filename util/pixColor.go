//package main
package util

import(
"fmt"
"image"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"
"image/color"
"log"
"net/http"
"strconv"
)
// returns the dominant color of a palette
func dominantColor(palette map[string]int) string{
	var max = 0
	var result string
	for couleur, nombre:= range palette{
		if nombre>max{
			max=nombre
			result=couleur
			fmt.Printf("color=%s, number=%d\n", couleur, nombre)
		}
	}
	return result
}
// converts a 8 bits integer to hexadecimal in string format
func int2hex(i uint8) string{
	hex:=fmt.Sprintf("%02x", i)
	return string(hex)
}
// converts a 2 characters hexadecimal to a 8 bits integer
func alamain(hex string) uint8{
	result:=0
	for i:=0; i<2; i++{
		switch hex[i]{
			case 'a':
				result+=10
				break
			case 'b':
				result+=11
				break
			case 'c':
				result+=12
				break
			case 'd':
				result+=13
				break
			case 'e':
				result+=14
				break
			case 'f':
				result+=15
				break
			default:
				tmp,_:=strconv.Atoi(string(hex[i]))
				result+=tmp
				break
		}
		if i==0{
			result*=16
		}
	}
	return uint8(result)
}
// converts the whole hexadecimal of 24 bits to 3 integers of 8 bits each
// defining the red, green and blue color
func hex2int(h string) (uint8,uint8,uint8){
	sr:=""
	sg:=""
	sb:=""
	for i:=0; i<len(h); i++{
		if 2<=i && i<=3{
			sr+=string(h[i])
		}else if 4<=i && i<=5{
			sg+=string(h[i])
		}else if 6<=i && i<=7{
			sb+=string(h[i])
		}
	}
	r:=alamain(sr)
	g:=alamain(sg)
	b:=alamain(sb)
	fmt.Printf("r=%d\n", r)
	fmt.Printf("g=%d\n", g)
	fmt.Printf("b=%d\n", b)
	return uint8(r), uint8(g), uint8(b)
}
// add a hexadecimal color in a map
// the map represents the number of each colors
func addTab(tab map[string]int, hex string){
	inside:=false
	for name,_:=range tab{
		if name==hex{
			inside=true
			tab[name]++
			break
		}
	}
	if inside==false{
		tab[hex]=1
	}
}
// takes an image url and returns the dominant color
// of type color.RGBA
func PixColor(url string) color.Color{
	resp, errh := http.Get(url)
	if errh!=nil{
		log.Fatal("get error:")
        log.Fatal(errh)
	}
	defer resp.Body.Close()
	m,_,err:=image.Decode(resp.Body)
	if err != nil {
		log.Fatal("decode error:")
        log.Fatal(err)
    }
	blocksize:=5
	tableau := make(map[string]int)
	rgb:=color.RGBA{R:0, G:0, B:0, A:1}
	bounds:=m.Bounds()
	for j:=bounds.Min.X; j < bounds.Max.X; j+=blocksize{
    	for i:=bounds.Min.Y; i < bounds.Max.Y; i+=blocksize{
			couleur:=m.At(j, i)
			r,g,b,_:=couleur.RGBA()
			hexa:="0x"+int2hex(uint8(r))+int2hex(uint8(g))+int2hex(uint8(b))
			addTab(tableau, hexa)
    	}
    }
	result:=dominantColor(tableau)
	rgb.R, rgb.G, rgb.B=hex2int(result)
	return rgb
}/*
func main(){
	couleur:=PixColor("http://i.imgur.com/Peq1U1u.jpg")
	//couleur:=PixColor("https://pmcmovieline.files.wordpress.com/2012/02/alextuis_bruceleespidey__120202000223.jpg?w=550&h=850")
	fmt.Printf("{r,g,b}=%d\n", couleur)
}*/
