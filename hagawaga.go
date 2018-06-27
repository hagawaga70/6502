package main

import . "./speicher"
import . "./register"
import . "gfx"
import . "./gfxElemente"
import "./dateien"
import "sort"
import "./opcode"
import "strconv"
import "os"
import "encoding/binary"
import "encoding/hex"
import "strings"

func main() {
	var speicher64k 		Speicher 	= NewSpeicher()
	var x_register 			Register 	= NewRegister()
	var y_register 			Register 	= NewRegister()
	var programmZaehlerHigh	Register 	= NewRegister()
	var programmZaehlerLow	Register 	= NewRegister()
	var stapelzeiger 		Register	= NewRegister()
	var akku 				Register	= NewRegister()
	var statusbits 			Register	= NewRegister()

	var gfxElement01 GfxElement = NewGfxElement()
	var speicher1 		[]byte 
	var speicher1old	[]byte
	var speicher2	 	[]byte
	var speicher2old 	[]byte
	var speicher3 		[]byte
	var speicher3old 	[]byte
	var speicher4 		[]byte
	var speicher4old 	[]byte
	var registerXdaten byte
	var registerYdaten byte
	var registerXdatenAlt byte
	var registerYdatenAlt byte
	var stapelzeigerDaten byte
	var stapelzeigerDatenAlt byte
	var akkuDaten byte
	var akkuDatenAlt byte
	var flags []string = []string{"C","Z","I","D","B","-","V","N"}
	var flagStatusBOOL bool
	var flagStatusINT int 
	var flagStatusAlt []int =[]int{0,0,0,0,0,0,0,0}

	var buffer  []byte
	var counter int
	var programmPath string = os.Args[1] 	// Übergabe des Programmpfades
	var mainSwitch = true 					// Hauptschalter: Hat er den Wert false, wird das Programm beendet
	speicher1old,	_ = speicher64k.Lesen([]uint16{0		, 255})		// Die speicherXold-variablen werden zum Vergleich
	speicher2old,	_ = speicher64k.Lesen([]uint16{256		, 511})		// benötigt, damit geänderte Speicherzeile rot
	speicher3old,	_ = speicher64k.Lesen([]uint16{512		, 767})     // markiert werden könnnen							
	speicher4old,	_ = speicher64k.Lesen([]uint16{768		, 1023})	//
	
	// Öffnen und lesen der Bytes des auszuführenden Programms
	dateiInhalt := dateien.Oeffnen(programmPath,'l')
	for !dateiInhalt.Ende(){
		buffer = append(buffer,dateiInhalt.Lesen())
		counter++
	}
	dateiInhalt.Schliessen()


	// Umwandeln der Bytes in EINE Zeichenkette
	hagCode 	:= 	string(buffer[:counter])

	// Umwandeln des Assemblerprogramms in eine opcodeListe. Eine Liste des Assemblerprogramms
	// und eine Liste der Pseudobefehle wird auch ausgegeben.
	opcodeList,assenblerCodeListe,pseudoCodeListe	:=	opcode.GetOpcodeList(hagCode) 

	// Die Key des Hashes "opcodeList" sind die Zeilennummern des um Leerzeilen, Kommentaren, Startadressen
	// und Pseudocode BEREINGTEN Assemblerprogramms. Diese Keys werden hier sortiert
    var keys []int
    for k := range opcodeList {
        keys = append(keys, k)
    }

    sort.Ints(keys)

	// Deklaration vom Variablen
	var switcher 			bool		//
	var singleOpcode		[]byte
	var opcodeFragment 		byte
	var stopAdresse 		uint16	// Start- und Stopadresse zum lesen oder schreiben von Speicheradressen
	var startAdresse 		uint16  // |
	var startAdressePcHEX 	string = pseudoCodeListe["$t@rt@dre$$e"][0]		// Startadresse für den Programmzählr
	var pcHigh 		byte	// Höheres Byte des Programmzähler
	var pcLow 		byte	// Niedriges Byte des Programmzählers

	var pcHighOld 	byte	// Speichert das vorherige Byte des Programmzählers. Wird benötigt zur roten Markierung
	var pcLowOld 	byte	// geänderter Werte

	opcodeRegister 		:=  map[string]int{} 		// Enhält als Schlüssel die hexadezimale Nummer der Opcodes 
	showOpcode 			:=  map[string][]string{} 	// Enthält als Schlüssel die Adresse des Opcodes und als Werte den gesammten Opcode 
	showAssemblerCode 	:=  map[string][]string{} 	// Enthält als Schlüssel die Adresse des Opcodes und als Werte den dazugehörigen Assemblercode


	// Die Opcodes werden hier in den Speicher geladen.
    for _, k := range keys { 
		// Im opcodeRegister wird der befehl-Opcode und seine jeweilige Länge gespeichert
		opcodeRegister[opcodeList[k][1]] = len(opcodeList[k])-1

 		showOpcode[opcodeList[k][0]] 		= opcodeList[k][1:]				// siehe oben
 		showAssemblerCode[opcodeList[k][0]] = assenblerCodeListe[k][0:]		// siehe oben

		singleOpcode = []byte{}												   // Zuweisung eines leeren Slice
		startAdresseUINT64, err := strconv.ParseUint(opcodeList[k][0], 16, 16) // Die hexadezimale Adresse wird in eine uint16 konvertiert 
		startAdresse = uint16(startAdresseUINT64)
		stopAdresse = startAdresse
		if err!=nil{
			panic(err)
		}
		switcher = true
		// Auslesen der einzelnen Elemente eines einzelnen Opcodes
		for _,value := range opcodeList[k]{
			if  switcher{			// Das erste Element speichert die Startadresse und wird daher übersprungen
				switcher = false
				continue
			}
			opcodeFragmentUINT64, err := strconv.ParseUint(value, 16, 8) // Die hexadezimale Adresse wird in eine uint16 konvertiert 
			if err!=nil{
				panic(err)
			}

			opcodeFragment = byte(uint8(opcodeFragmentUINT64))
			singleOpcode = append(singleOpcode,opcodeFragment)
			stopAdresse++
		}
		// Abspeichern der einzelnen Opcodeelemente(Byte) pro Opcode
		_ = speicher64k.Schreiben([]uint16{startAdresse,stopAdresse-1}, singleOpcode)
    }


	var getOpcode []byte				// Enthält die einzelnen Opcodeelemente eines Opcodes
	var opcodeHeadAdresse uint16		// Enthält die Adresse des ersten Elements des Opcodes 
	var switchStartAdresse bool = true	//
	var anzahlOpcodeElemente int		//
	var pseudoCodeContentSwitch bool 	// Beim ersten Schleifendurchlauf muss die Startadresse in den
										// Programmzähler geladen werden. Danach nicht mehr. Daher der
										// Switch

	Fenster(1920, 1200)

	for {
		// Unterbrich die Endlosschleife, wenn der Opcode f2 (ENDE) gelesen wurde
		if !mainSwitch {
			break
		}

		// Wird für die Anzeige der Pseudocodes benötigt
		pseudoCodeContentSwitch  = true 

		// Die hexadezimale Adresse wird in eine uint64 konvertiert
		startAdressePcUINT64, err := strconv.ParseUint(startAdressePcHEX, 16, 16) 

		if err!=nil{
			panic(err)
		}

		// Einmaliges schreiben der Startadressen in den Programmzähler
		if switchStartAdresse{
			startAdressePcUINT16 := uint16(startAdressePcUINT64)
			startAdresseByte := make([]byte, 2)						// Ein Slice mit zwei Byte wird erstellt
			binary.BigEndian.PutUint16(startAdresseByte, startAdressePcUINT16) // Konvertierung von uint16 -> byte
			switchStartAdresse = false

			_ = programmZaehlerHigh.SchreibenByte(startAdresseByte[0])		// Schreiben der Startadresse in den Programm-
			_ = programmZaehlerLow.SchreibenByte(startAdresseByte[1])       // zähler. Als Rückgabewert war angedacht die
																			// benötigten Takte des Microprozessors zu liefern
																			// Dieses Feature habe ich nicht umgesetzt
		}

		// Auslesen des Programmzählers
		pcHigh, _	= programmZaehlerHigh.LesenByte()
		pcLow, _	= programmZaehlerLow.LesenByte()

		// Konvertiere byte in uint16
		opcodeHeadAdresse = binary.BigEndian.Uint16([]byte{pcHigh,pcLow})

		// Lesen der im Programmzähler angegebenen Adresse in Byte
		getOpcode,	_ = speicher64k.Lesen([]uint16{opcodeHeadAdresse, opcodeHeadAdresse})  // 

		// Der opcode f2 gibt an, dass das Programm zu Ende ist. Hat mainSwitch den Wert false wird beim nächsten
		// Schleifendurchlauf ganz am Anfang die Endlosschleife durch ein break abgebrochen
		if hex.EncodeToString(getOpcode)== "f2"{
			mainSwitch=false
		}

		// Nachschlagen, wie viele weitere Bytes zum Opcode gehören
		anzahlOpcodeElemente =  opcodeRegister[hex.EncodeToString(getOpcode)]

		// Auslesen aller OpcodeElemente zu einem Befehl
		getOpcode,	_ = speicher64k.Lesen([]uint16{opcodeHeadAdresse, opcodeHeadAdresse+uint16(anzahlOpcodeElemente)-1})  // 

		// Berechnet die nächste OpcodeHeadAdresse
		opcodeHeadAdresse = opcodeHeadAdresse+uint16(anzahlOpcodeElemente )


		// Ausführen des Opcodes
		_ =opcode.ExecuteOpcode ( 	getOpcode,
									speicher64k, 
									x_register,
									y_register,
									programmZaehlerHigh,
									programmZaehlerLow,
									stapelzeiger, 
									akku, 
									statusbits)

		// Hintergrundfarbe des gesamten Bildschirmshintergrunds
		Stiftfarbe(220, 222, 217)
		Vollrechteck(0, 0, 1920, 1200)

		// Auslesen des Speichers zum Anzeigen von vier Speicherseiten
		speicher1,	_ = speicher64k.Lesen([]uint16{0		, 255})		// Seite 0: Wird mit einem Byte adressiert
		speicher2,	_ = speicher64k.Lesen([]uint16{256		, 511})		// Seite 1: Wird für den Stack benutzt
		speicher3,	_ = speicher64k.Lesen([]uint16{512		, 767})		// Seite 2
		speicher4,	_ = speicher64k.Lesen([]uint16{768		, 1023})	// Seite 3
		UpdateAus()


		// Label AssemblerCode-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,670,"Assembler-Code",24,0,0,255) // alt 10,950
		gfxElement01.AbbildLabel(1630,700,strings.Join(showAssemblerCode[strconv.FormatUint(uint64(opcodeHeadAdresse-uint16(anzahlOpcodeElemente)),16)][:]," "),24,255,0,0)

		// Label Opcode-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,760,"Opcode",24,0,0,255)
		gfxElement01.AbbildLabel(1630,790,strings.Join(showOpcode[strconv.FormatUint(uint64(opcodeHeadAdresse-uint16(anzahlOpcodeElemente)),16)][:]," "),24,255,0,0)

		
		// Bild Hagawaga-------------------------------------------------------------------------
		gfxElement01.AbbildBild(10,910,"./images/hagawaga2.bmp")

		// Label Pseudocode-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(400,950,"Pseudocode",24,0,0,255)
		var counter int
		var y		uint16 = 980
		var x		uint16 = 200
		var rowElements int =7
		
		// Die Pseudocodes und die dazugehörigen Werte werden im Grafikfenster angezeigt
		keys := make([]string, 0, len(pseudoCodeListe))
		for key := range pseudoCodeListe{
			keys = append(keys, key)
		}
		sort.Strings(keys) //sort by key
		for _, key := range keys {

			list := pseudoCodeListe[key]
			if  key == "$t@rt@dre$$e"{
				continue
			}
			if pseudoCodeContentSwitch && counter==7{
				rowElements = 7
				pseudoCodeContentSwitch=false
			}else if !pseudoCodeContentSwitch{
				rowElements = 6
			}
			if counter == rowElements{
				y=y+30
				gfxElement01.AbbildLabel(400,y,key+"="+list[0],24,0,0,0)
				x=400
				counter=0
			}else{
				x=x+200
				gfxElement01.AbbildLabel(x,y,key+"="+list[0],24,0,0,0)
				counter++
			}
		}

		// Anzeige der vier Speicherseiten im Grafikfenster
		gfxElement01.AbbildSpeicherseite1(10	, 10	, 0		, speicher1 ,speicher1old	)
		gfxElement01.AbbildSpeicherseite1(411	, 10	, 1		, speicher2	,speicher2old	)
		gfxElement01.AbbildSpeicherseite1(812	, 10	, 2		, speicher3	,speicher3old	)
		gfxElement01.AbbildSpeicherseite1(1213	, 10	, 3		, speicher4	,speicher4old	)


		speicher1old = speicher1	// Die speicherXold-Variablen werden zum Vergleich
		speicher2old = speicher2	// benötigt, damit geänderte Speicherzeile rot
		speicher3old = speicher3	// markiert werden könnnen							
		speicher4old = speicher4



		// Label Register-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,10,"Register",24,0,0,255)

		// Anzeigen des X-Registers -------------------------------------------------------------------------
		registerXdaten, _ = x_register.LesenByte()
		gfxElement01.AbbildRegister(1630, 40, "X-Register", registerXdaten, registerXdatenAlt)
		registerXdatenAlt = registerXdaten

		// Anzeigen des Y-Registers ----------------------------------------------------------------------------
		registerYdaten, _ = y_register.LesenByte()
		gfxElement01.AbbildRegister(1630, 70, "Y-Register", registerYdaten, registerYdatenAlt)
		registerYdatenAlt = registerYdaten

		// Anzeigen des Akkus ----------------------------------------------------------------------------------
		akkuDaten, _ = akku.LesenByte()
		gfxElement01.AbbildRegister(1630, 100, "Akku", akkuDaten, akkuDatenAlt)
		akkuDatenAlt = akkuDaten

		// Label Stapelzeiger-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,160,"Stapelzeiger",24,0,0,255)

		// Anzeigen des Stapelzeigers --------------------------------------------------------------------------
		stapelzeigerDaten, _ = stapelzeiger.LesenByte()
		gfxElement01.AbbildRegister(1630, 190, "SZ", stapelzeigerDaten, stapelzeigerDatenAlt)
		stapelzeigerDatenAlt = stapelzeigerDaten

		// Label Programmzähler-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,250,"Progammzähler",24,0,0,255)

		// Anzeigen des Programmzählers ----------------------------------------------------------------------------
		pcHigh, _ 	= programmZaehlerHigh.LesenByte()
		pcLow, _ 	= programmZaehlerLow.LesenByte()

		gfxElement01.AbbildRegister(1630, 280, "High", pcHigh, pcHighOld)
		pcHighOld = pcHigh

		gfxElement01.AbbildRegister(1630, 310, "Low", pcLow, pcLowOld)
		pcLowOld = pcLow

		// Label Flags-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,370,"Flags",24,0,0,255)

		// Anzeigen der Flags  ---------------------------------------------------------------------------------

		for  index,flag := range(flags){
			flagStatusBOOL, _ =statusbits.LeseBit( uint(index))

			if flagStatusBOOL {
				flagStatusINT = 1
			}else{
				flagStatusINT = 0
			}
			gfxElement01.AbbildFlag(1630, uint16(400+index*30), flag, flagStatusINT, flagStatusAlt[index])
			flagStatusAlt[index]= flagStatusINT
		}

		// Der Programmzähler wird auf die Adresse des nächsten Opcode-HEAD gesetzt
		adressenBytes := make([]byte, 2)						// Ein Slice mit zwei Byte wird erstellt
		binary.BigEndian.PutUint16(adressenBytes, opcodeHeadAdresse)

		_ = programmZaehlerHigh.SchreibenByte(adressenBytes[0])
		_ = programmZaehlerLow.SchreibenByte(adressenBytes[1])


		UpdateAn()
		TastaturLesen1()
		Cls()
	}
}
