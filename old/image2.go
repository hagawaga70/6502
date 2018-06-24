package main


import (
    "golang.org/x/image/font"
    "golang.org/x/image/font/basicfont"
    "golang.org/x/image/math/fixed"
	"golang.org/x/image/bmp"
    "image"
    "image/color"
    //"image/png"
    "os"
)

func addLabel(img *image.RGBA, x, y int, label string) {
    //col := color.RGBA{200, 100, 0, 255}
    col := color.RGBA{50, 205, 50, 10}
    point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

    d := &font.Drawer{
        Dst:  img,
        Src:  image.NewUniform(col),
        Face: basicfont.Face7x13,
        Dot:  point,
    }
    d.DrawString(label)
}

func main() {
    img := image.NewRGBA(image.Rect(0, 0, 200, 360))
    for i:=0;i<32;i++{
    		addLabel(img, 20, 30+(i*10), "AA FF 00 99 DD 34 55 12")
	}	
    	f, err := os.Create("hello-gq9.bmp")
    if err != nil {
        panic(err)
    }
	//err:=Encode(f,img)
    defer f.Close()
    if err := bmp.Encode(f, img); err != nil {
        panic(err)
    }
}
