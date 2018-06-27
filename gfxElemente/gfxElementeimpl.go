package gfxElemente

import "encoding/hex"
import . "gfx"
import . "strconv"

import "encoding/binary"

type impl struct {
	seite    [255]byte
	register byte
}

func NewGfxElement() *impl {
	var gfxElement *impl
	gfxElement = new(impl)
	return gfxElement
}

// Positioniert und zeigt ein unter dem pfad angegebenes Bild im Grafikfenster
func (gfxElement *impl) AbbildBild(x1, y1 uint16,pfad string) {
	LadeBild(x1, y1, pfad)
}


// Positioniert und zeigt eine Seite des Hauptspeichers. Gibt die Seite den Speicherbereich und die letzte Adresse
// der visualisierten Speicherzeile an
func (gfxElement *impl) AbbildSpeicherseite1(x1, y1 uint16, seite uint, seitenInhalt []byte,seitenInhaltAlt []byte) {
	var counter uint
	var byteBuffer byte
	var startAdresseDerSeite uint16 = uint16(seite) * 256            // Erste Adresse der Seite n
	var stopAdresseDerSeite uint16 = ((uint16(seite) + 1) * 256) - 1 // Letzte Adresse der Seite n
	var labelZeile string
	var zeilenNummer byte = 0x00

	// Erzeugung eines Seitenfensters mit einem Rand
	Stiftfarbe(152, 13, 8)
	Vollrechteck(x1, y1, 400, 900)
	Stiftfarbe(217, 222, 226)
	Vollrechteck(x1+2, y1+2, 396, 896)
	Stiftfarbe(0, 5, 2)

	// Schreibt die Seitenanzahl in das Seitenfenster
	SetzeFont("./font/LiberationMono-Regular.ttf", 20)
	SchreibeFont(19+x1, 30, "Seite:    "+Itoa(int(seite))) // Schreiben des Textes -> Seite: n

	adresseStartDez := uint16(startAdresseDerSeite)               // Die dezimale Startadresse wir in ein Byte-Array
	adresseStartByte := make([]byte, 2)                           // konvertieret
	binary.BigEndian.PutUint16(adresseStartByte, adresseStartDez) // <-----

	adresseStopDez := uint16(stopAdresseDerSeite)               // Die dezimale Stopadresse wir in ein Byte-Array
	adresseStopByte := make([]byte, 2)                          // konvertiert
	binary.BigEndian.PutUint16(adresseStopByte, adresseStopDez) // <-----

	// Schreibt den Adressbereich der Speicherseite in das Seitenfenster
	SchreibeFont(19+x1, 60, "Adressb.: ["+hex.EncodeToString(adresseStartByte)+"] - "+"["+hex.EncodeToString(adresseStopByte)+"]")
	var labelZeileAlt string
	var byteBufferAlt byte

	// Schreibt die Bytes des Speicherbereicht hexadezimal in 8-stellige Speicherzeilen 
	for i := 0; i <= 255; i++ {
		byteBuffer, seitenInhalt = seitenInhalt[0], seitenInhalt[1:]

		// Wird benötigt zur rötlichen Markierung veränderter Seitenzeilen
		byteBufferAlt, seitenInhaltAlt = seitenInhaltAlt[0], seitenInhaltAlt[1:]

		zeilenNummer++

		// Erstellen einer Seitenzeie. Die enthält 8 Bytes. Der Wert des Bytes wird hexadezimal angezeigt
		labelZeile = labelZeile + hex.EncodeToString([]byte{byteBuffer}) + " "

		// Wird benötigt zur rötlichen Markierung veränderter Seitenzeilen
		labelZeileAlt = labelZeileAlt + hex.EncodeToString([]byte{byteBufferAlt}) + " "

		counter++

		if counter == 8 {
			// Die Adresse zur Seite 0 besteht aus einem Byte. Daher muss ein der Funktion
			// hex.EncodeToString ein weiteres  Byte byte(0)  mit dem Wert 0 übergeben werden,
			// Damit die Funktion funktioniert
			if seite == 0 { 
				labelZeile = "[" + hex.EncodeToString([]byte{byte(0), byte(i)}) + "] " + labelZeile
				labelZeileAlt = "[" + hex.EncodeToString([]byte{byte(0), byte(i)}) + "] " + labelZeileAlt
			} else {
				adresseDez := uint16(int(seite)*256 + i)
				adresseByte := make([]byte, 2)
				binary.BigEndian.PutUint16(adresseByte, adresseDez)

				// Erstellt Seitenzeile inklusive [letzte Adresse der Zeile]
				labelZeile = "[" + hex.EncodeToString([]byte(adresseByte)) + "] " + labelZeile

				// Wird benötigt, um eine veränderte Zeile rötlich zu markieren
				labelZeileAlt = "[" + hex.EncodeToString([]byte(adresseByte)) + "] " + labelZeileAlt
			}

			// Stiftfarbe wird auf rot gesetzt, um veränderte Seitenzeilen zu markieren. Gibt es 
			// keine Seitenzeilenveränderung bleibt die Stiftfarbe schwarz
			if labelZeileAlt != labelZeile{
				Stiftfarbe(255,0,0)
			}else{
				Stiftfarbe(0, 5, 2)
			}

			// Die Seitenzeile wird im Grafikfenster ausgegeben. Dabei wird die Position abhängig
			// von den n-ten Seitenzeile neuberechnet
			SchreibeFont(19+x1, uint16(i*3+85), labelZeile)
			labelZeile = ""
			labelZeileAlt = ""
			counter = 0
		}
	}

}

// Positioniert und schreibt ein Register ins Grafikfenster
func (gfxElement *impl) AbbildRegister(x1, y1 uint16, name string, registerInhalt byte, registerInhaltAlt byte) {
	var label string
	var labelOffset uint16

	Stiftfarbe(0, 81, 47)
	SetzeFont("./font/LiberationMono-Regular.ttf", 24)

	// Die übergebene Zeichkette wird durch Leerzeichen auf eine einheitliche Länge gebracht. Am Ende wird 
	// eine Doppelpunkt angehängt.
	label = labelanpassung(name)
	SchreibeFont(x1, y1, label)
	if registerInhaltAlt != registerInhalt{
		Stiftfarbe(255, 0, 0)
	}else{
		Stiftfarbe(0, 0, 0)
	}
	labelOffset = 12 * 16
	SchreibeFont(x1+labelOffset, y1, hex.EncodeToString([]byte{registerInhalt}))

}


// Positioniert und schreibt ein Label ins Grafikfenster
func (gfxElement *impl) AbbildLabel(x1, y1 uint16, label string, schriftgroesse int, r uint8, g uint8, b uint8) {
	
	Stiftfarbe(r,g,b)
	SetzeFont("./font/LiberationMono-Regular.ttf", schriftgroesse)
	SchreibeFont(x1, y1, label)

}

// Positioniert und schreibt den Namen und den Wert eines Statusbits ins Grafikfenster
func (gfxElement *impl) AbbildFlag(x1, y1 uint16, label string, flagStatus int, flagStatusAlt int) {
	var labelOffset uint16

	Stiftfarbe(0, 81, 47)
	SetzeFont("./font/LiberationMono-Regular.ttf", 24)

	// Schreiben des Namen des Statusbits
	SchreibeFont(x1, y1, label + " -flag: ")

	// Hat sich der Wert des Statusbit verändert wird er rot geschrieben 
	if flagStatusAlt != flagStatus{
		Stiftfarbe(255, 0, 0)
	}else{
		Stiftfarbe(0, 0, 0)
	}

	// Schreiben des Wertes des Statusbits
	labelOffset = 12 * 16
	SchreibeFont(x1+labelOffset, y1, Itoa(flagStatus))

}

// Hängt Leerzeichen an einen übergebenen String. Somit kann alles in Reih und Glied im Grafikfenster 
// angezeigt werden :-)
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

