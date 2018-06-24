package main

import (
    "flag"
    "fmt"
    "io/ioutil"
    "log"
    "image"
    "bufio"
    "image/draw"
    //"image/png"
    //"image/png"
    "golang.org/x/image/bmp"
    //"image/color"
    "github.com/golang/freetype/truetype"
    "golang.org/x/image/font"
    "github.com/golang/freetype"
    "os"
)

var (
    dpi      = flag.Float64("dpi", 72, "screen resolution in Dots Per Inch")
    fontfile = flag.String("fontfile", "LiberationMono-Regular.ttf", "filename of the ttf font")
    hinting  = flag.String("hinting", "none", "none | full")
    size     = flag.Float64("size", 16, "font size in points")
    spacing  = flag.Float64("spacing", 2, "line spacing (e.g. 2 means double spaced)")
    wonb     = flag.Bool("whiteonblack", false, "white text on a black background")
    text     = string("JOJO")
)


func main() {
    flag.Parse()
    fmt.Printf("Loading fontfile %q\n", *fontfile)
    b, err := ioutil.ReadFile(*fontfile)
    if err != nil {
        log.Println(err)
        return
    }
    f, err := truetype.Parse(b)
    if err != nil {
        log.Println(err)
        return
    }

    // Freetype context
    fg, bg := image.Black, image.White
    rgba := image.NewRGBA(image.Rect(0, 0, 1000, 200))
    draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
    c := freetype.NewContext()
    c.SetDPI(*dpi)
    c.SetFont(f)
    c.SetFontSize(*size)
    c.SetClip(rgba.Bounds())
    c.SetDst(rgba)
    c.SetSrc(fg)
    switch *hinting {
    default:
        c.SetHinting(font.HintingNone)
    case "full":
        c.SetHinting(font.HintingFull)
    }

    // Truetype stuff
    opts := truetype.Options{}
    opts.Size = 125.0
    face := truetype.NewFace(f, &opts)


    // Calculate the widths and print to image
    var counter uint =10
    for i, x := range(text) {
        awidth, ok := face.GlyphAdvance(rune(x))
        if ok != true {
            log.Println(err)
            return
        }
        iwidthf := int(float64(awidth) / 64)
        fmt.Printf("%+v\n", iwidthf)

        pt := freetype.Pt(10,i*20)
        c.DrawString(string(x), pt)
	   counter=counter+10
    }


    // Save that RGBA image to disk.
    outFile, err := os.Create("out.bmp")
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
    fmt.Println("Wrote out.bmp OK.")


}
