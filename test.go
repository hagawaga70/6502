package main

import . "./speicher"
import . "./register"
import . "gfx"

//import "time"
import . "./gfxElemente"

//import "strconv"

import "fmt"

func main() {
	var speicher64k Speicher = NewSpeicher()
	var x_register Register = NewRegister()
	var y_register Register = NewRegister()
	var stapelzeiger = NewRegister()
	var akku = NewRegister()
	var gfxElement01 GfxElement = NewGfxElement()
	var takte int
	var speicher []byte
	var speicher2 []byte
	var speicher3 []byte
	var speicher4 []byte
	var registerXdaten byte
	var registerYdaten byte
	var registerXdatenAlt byte
	var registerYdatenAlt byte
	var stapelzeigerDaten byte
	var stapelzeigerDatenAlt byte
	var akkuDaten byte
	var akkuDatenAlt byte

	Fenster(1900, 1200)
	for i := 0; i < 3; i++ {
		Stiftfarbe(220, 222, 217)      // Hintergrundfarbe des gesamten Bildschirms
		Vollrechteck(0, 0, 1920, 1200) // Bildschirmhintergrund

		takte = speicher64k.Schreiben([]int16{256, 259}, []byte{byte(10), byte(3), byte(4), byte(5)})
		speicher, takte = speicher64k.Lesen([]int16{256, 511})
		speicher2, takte = speicher64k.Lesen([]int16{0, 255})
		speicher3, takte = speicher64k.Lesen([]int16{0, 255})
		speicher4, takte = speicher64k.Lesen([]int16{0, 255})
		fmt.Println(speicher)
		UpdateAus()
		gfxElement01.AbbildSpeicherseite1(10, 10, 0, speicher2)
		gfxElement01.AbbildSpeicherseite1(411, 10, 1, speicher)
		gfxElement01.AbbildSpeicherseite1(812, 10, 2, speicher3)
		gfxElement01.AbbildSpeicherseite1(1213, 10, 255, speicher4)
		//time.Sleep(1000000000)
		fmt.Println(takte)
		fmt.Println(speicher)
		// Anzeigen des X-Registers -------------------------------------------------------------------------
		takte = x_register.SchreibenByte(byte(5 + i))
		registerXdaten, takte = x_register.LesenByte()
		gfxElement01.AbbildRegister(1630, 10, "X-Register", registerXdaten, registerXdatenAlt)
		registerXdatenAlt = registerXdaten
		// <<---------------------------------------------------------------------------------------------------

		// Anzeigen des Y-Registers -------------------------------------------------------------------------
		takte = y_register.SchreibenByte(byte(12 + i))
		registerYdaten, takte = y_register.LesenByte()
		gfxElement01.AbbildRegister(1630, 40, "Y-Register", registerYdaten, registerYdatenAlt)
		registerYdatenAlt = registerYdaten
		// <<---------------------------------------------------------------------------------------------------

		// Anzeigen des Akkus -------------------------------------------------------------------------
		takte = akku.SchreibenByte(byte(2 + i))
		akkuDaten, takte = akku.LesenByte()
		gfxElement01.AbbildRegister(300, 950, "Akku", akkuDaten, akkuDatenAlt)
		akkuDatenAlt = akkuDaten
		// <<---------------------------------------------------------------------------------------------------

		// Anzeigen des Stapelzeigers -------------------------------------------------------------------------
		takte = stapelzeiger.SchreibenByte(byte(15 + i))
		stapelzeigerDaten, takte = stapelzeiger.LesenByte()
		gfxElement01.AbbildRegister(300, 1000, "Stapelzeiger", stapelzeigerDaten, stapelzeigerDatenAlt)
		stapelzeigerDatenAlt = stapelzeigerDaten
		// <<---------------------------------------------------------------------------------------------------
		UpdateAn()
		TastaturLesen1()
		Cls()
	}
}
