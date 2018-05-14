package gfxElemente

//import "strconv"
import "encoding/hex"
import . "gfx"

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
	var x1Buffer = x1
	//var y1Buffer = y1
	var byteBuffer byte
	Stiftfarbe(0xFF, 0, 0)
	SetzeFont("./font/OpenSans-CondLight.ttf", int(groesse))
	for i := 0; i < 256; i++ {
		byteBuffer, seiteninhalt = seiteninhalt[0], seiteninhalt[1:]

		if counter == 8 {
			counter = 0
			y1 = y1 + uint16(groesse)
			x1 = x1Buffer
		}

		SchreibeFont(x1, y1, hex.EncodeToString([]byte{byteBuffer}))
		//SchreibeFont(x1, y1, string(byteBuffer))
		x1 = x1 + uint16(groesse)
		counter++
	}
}
