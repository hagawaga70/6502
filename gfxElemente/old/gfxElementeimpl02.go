package gfxElemente

//import "strconv"
import "encoding/hex"
import . "gfx"
import "golang.org/x/image/font"

import "golang.org/x/image/font/basicfont"

//import "golang.org/x/image/font/inconsolata"
//import "golang.org/x/image/font/gofont/gosmallcaps"
import "golang.org/x/image/math/fixed"
import "golang.org/x/image/bmp"
import "image"
import "image/color"
import "os"
import "fmt"

type impl struct {
	seite    [255]byte
	register byte
}

func NewGfxElement() *impl {
	var gfxElement *impl
	gfxElement = new(impl)
	return gfxElement
}
func (gfxElement *impl) AbbildSpeicherseite(x1, y1 uint16, seite uint, seiteninhalt []byte) {
	var counter uint
	//var x1Buffer = x1
	var byteBuffer byte
	var startAdresseDerSeite uint16
	var stopAdresseDerSeite uint16
	var labelZeile string
	var zeilenNummer byte = 0x00
	startAdresseDerSeite = uint16(seite) * 256
	stopAdresseDerSeite = ((uint16(seite) + 1) * 256) - 1

	img := image.NewRGBA(image.Rect(0, 0, 250, 800))

	addLabel(img, 5, 30, " Seite   : "+string(seite))
	addLabel(img, 5, 40, " Adressb.: "+string(startAdresseDerSeite)+" "+string(stopAdresseDerSeite))
	for i := 0; i < 256; i++ {
		byteBuffer, seiteninhalt = seiteninhalt[0], seiteninhalt[1:]
		zeilenNummer++
		labelZeile = labelZeile + hex.EncodeToString([]byte{byteBuffer}) + " "
		counter++
		if counter == 8 {
			counter = 0
			labelZeile = "[" + hex.EncodeToString([]byte{zeilenNummer}) + "] " + labelZeile
			addLabel(img, 20, 50+(i*2), labelZeile)
			labelZeile = ""
		}
	}
	f, err := os.Create("hello-gq9.bmp")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := bmp.Encode(f, img); err != nil {
		panic(err)
	}
	LadeBild(x1, y1, "hello-gq9.bmp")
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{255, 0, 0, 125}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	//ascent := face.Metrics().Ascent.Ceil()
	f := font.Metrics{
		Height: fixed.Int26_6(40),
	}
	fmt.Println(f)

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		//Face: inconsolata.Regular8x16,
		//Face: gosmallcaps.TTF,
		Dot: point,
		//Dot: fixed.P(0, ascent),
	}
	d.DrawString(label)
}
