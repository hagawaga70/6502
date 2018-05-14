package gfxElemente

import "encoding/hex"
import . "gfx"
import . "strconv"

//import "fmt"
import "encoding/binary"

import (
	"bufio"
	"flag"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/bmp"
	//"golang.org/x/image/font"
	"image"
	"image/draw"
	"io/ioutil"
	"log"
	"os"
)

type impl struct {
	seite    [255]byte
	register byte
}

func NewGfxElement() *impl {
	var gfxElement *impl
	gfxElement = new(impl)
	return gfxElement
}

var (
	dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
	fontfile = flag.String("fontfile", "./font/LiberationMono-Regular.ttf", "filename of the ttf font")
	hinting  = flag.String("hinting", "none", "none | full")
	size     = flag.Float64("size", 20, "font size in points")
	spacing  = flag.Float64("spacing", 1, "line spacing (e.g. 2 means double spaced)")
	wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
)

func (gfxElement *impl) AbbildSpeicherseite(x1, y1 uint16, seite uint, seiteninhalt []byte) {
	var counter uint
	var byteBuffer byte
	var startAdresseDerSeite uint16 = uint16(seite) * 256            // Erste Adresse der Seite n
	var stopAdresseDerSeite uint16 = ((uint16(seite) + 1) * 256) - 1 // Letzte Adresse der Seite n
	var labelZeile string
	var zeilenNummer byte = 0x00
	var bmpDateiName string = "seite" + Itoa(int(seite)) + ".bmp"
	var pfad string = "./seitenBMPs/"
	flag.Parse()                         // ???
	b, err := ioutil.ReadFile(*fontfile) // |
	if err != nil {                      // |
		log.Println(err) // |
		return           // |
	} // |
	f, err := truetype.Parse(b) // |
	if err != nil {             // |
		log.Println(err) // |
		return           // ^
	}

	bg, fg := image.Black, image.White                     // BMP-Speicherseite: Festlegen der Bildparameter
	rgba := image.NewRGBA(image.Rect(0, 0, 400, 900))      // |
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src) // |
	c := freetype.NewContext()                             // |
	c.SetDPI(*dpi)                                         // |
	c.SetFont(f)                                           // |
	c.SetFontSize(*size)                                   // |
	c.SetClip(rgba.Bounds())                               // |
	c.SetDst(rgba)                                         // |
	c.SetSrc(fg)                                           // ^-----

	/*
		switch *hinting {
		default:
			c.SetHinting(font.HintingNone)
		case "full":
			c.SetHinting(font.HintingFull)
		}
	*/

	// Truetype stuff
	//opts := truetype.Options{}
	//opts.Size = 12.0
	////face := truetype.NewFace(f, &opts)
	pt := freetype.Pt(10, 30)                       // Position des Textes -> Seite: n
	c.DrawString("Seite:    "+Itoa(int(seite)), pt) // Schreiben des Textes -> Seite: n
	// <-----

	adresseStartDez := uint16(startAdresseDerSeite)               // Die dezimale Startadresse wir in ein Byte-Array
	adresseStartByte := make([]byte, 2)                           // konvertieret
	binary.BigEndian.PutUint16(adresseStartByte, adresseStartDez) // <-----

	adresseStopDez := uint16(stopAdresseDerSeite)               // Die dezimale Stopadresse wir in ein Byte-Array
	adresseStopByte := make([]byte, 2)                          // konvertiert
	binary.BigEndian.PutUint16(adresseStopByte, adresseStopDez) // <-----

	pt = freetype.Pt(10, 60)
	c.DrawString("Adressb.: ["+hex.EncodeToString(adresseStartByte)+"] - "+"["+hex.EncodeToString(adresseStopByte)+"]", pt)
	for i := 0; i <= 255; i++ {
		byteBuffer, seiteninhalt = seiteninhalt[0], seiteninhalt[1:]
		zeilenNummer++
		labelZeile = labelZeile + hex.EncodeToString([]byte{byteBuffer}) + " "
		counter++

		if counter == 8 {
			if seite == 0 {
				labelZeile = "[" + hex.EncodeToString([]byte{byte(0), byte(i)}) + "] " + labelZeile
			} else {
				adresseDez := uint16(int(seite)*256 + i)
				adresseByte := make([]byte, 2)
				binary.BigEndian.PutUint16(adresseByte, adresseDez)
				//labelZeile = "[" + hex.EncodeToString([]byte{Itoa(int(seite)*256 + i)}) + "] " + labelZeile
				labelZeile = "[" + hex.EncodeToString([]byte(adresseByte)) + "] " + labelZeile
			}

			pt := freetype.Pt(10, i*3+85)
			c.DrawString(labelZeile, pt)
			labelZeile = ""
			counter = 0
		}
	}

	// Save that RGBA image to disk.
	outFile, err := os.Create(pfad + bmpDateiName)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer outFile.Close()
	bf := bufio.NewWriter(outFile)
	err = bmp.Encode(bf, rgba)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = bf.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	LadeBild(x1, y1, pfad+bmpDateiName)
}

func (gfxElement *impl) AbbildSpeicherseite1(x1, y1 uint16, seite uint, seiteninhalt []byte) {
	var counter uint
	var byteBuffer byte
	var startAdresseDerSeite uint16 = uint16(seite) * 256            // Erste Adresse der Seite n
	var stopAdresseDerSeite uint16 = ((uint16(seite) + 1) * 256) - 1 // Letzte Adresse der Seite n
	var labelZeile string
	var zeilenNummer byte = 0x00
	Stiftfarbe(152, 13, 8)
	Vollrechteck(x1, y1, 400, 900)
	Stiftfarbe(217, 222, 226)
	Vollrechteck(x1+2, y1+2, 396, 896)
	Stiftfarbe(0, 5, 2)
	SetzeFont("./font/LiberationMono-Regular.ttf", 20)
	SchreibeFont(19+x1, 30, "Seite:    "+Itoa(int(seite))) // Schreiben des Textes -> Seite: n
	// <-----

	adresseStartDez := uint16(startAdresseDerSeite)               // Die dezimale Startadresse wir in ein Byte-Array
	adresseStartByte := make([]byte, 2)                           // konvertieret
	binary.BigEndian.PutUint16(adresseStartByte, adresseStartDez) // <-----

	adresseStopDez := uint16(stopAdresseDerSeite)               // Die dezimale Stopadresse wir in ein Byte-Array
	adresseStopByte := make([]byte, 2)                          // konvertiert
	binary.BigEndian.PutUint16(adresseStopByte, adresseStopDez) // <-----

	SchreibeFont(19+x1, 60, "Adressb.: ["+hex.EncodeToString(adresseStartByte)+"] - "+"["+hex.EncodeToString(adresseStopByte)+"]")
	for i := 0; i <= 255; i++ {
		byteBuffer, seiteninhalt = seiteninhalt[0], seiteninhalt[1:]
		zeilenNummer++
		labelZeile = labelZeile + hex.EncodeToString([]byte{byteBuffer}) + " "
		counter++

		if counter == 8 {
			if seite == 0 {
				labelZeile = "[" + hex.EncodeToString([]byte{byte(0), byte(i)}) + "] " + labelZeile
			} else {
				adresseDez := uint16(int(seite)*256 + i)
				adresseByte := make([]byte, 2)
				binary.BigEndian.PutUint16(adresseByte, adresseDez)
				//labelZeile = "[" + hex.EncodeToString([]byte{Itoa(int(seite)*256 + i)}) + "] " + labelZeile
				labelZeile = "[" + hex.EncodeToString([]byte(adresseByte)) + "] " + labelZeile
			}

			SchreibeFont(19+x1, uint16(i*3+85), labelZeile)
			labelZeile = ""
			counter = 0
		}
	}

}

func (gfxElement *impl) AbbildRegister(x1, y1 uint16, name string, registerInhalt byte, registerInhaltAlt byte) {
	var label string
	var labelOffset uint16
	Stiftfarbe(0, 81, 47)
	SetzeFont("./font/LiberationMono-Regular.ttf", 24)
	label = labelanpassung(name)
	SchreibeFont(x1, y1, label)
	Stiftfarbe(0, 0, 0)
	labelOffset = 12 * 16
	SchreibeFont(x1+labelOffset, y1, hex.EncodeToString([]byte{registerInhalt}))

}

func labelanpassung(label string) string {
	var labelLaengeSoll int = 12
	var labelLaengeISt int = len(label)

	if labelLaengeSoll-labelLaengeISt < 0 {
		panic("gfxElementimp::Error::01: Das Label ist zu lang!")
	}
	for i := 0; i < (labelLaengeSoll - labelLaengeISt); i++ {
		label = label + " "
	}
	label = label + ": "

	return label

}
