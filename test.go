package main

//test git
import . "./speicher"
import . "./register"
import . "gfx"
import . "./gfxElemente"
import "fmt"
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
	//var takte int
	var speicher1 []byte
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
	var flags []string = []string{"C","Z","I","D","B","-","V","N"}
	var flagStatusBOOL bool
	var flagStatusINT int 
	var flagStatusAlt []int =[]int{0,0,0,0,0,0,0,0}

	var buffer  []byte
	var counter int
	var programmPath string = os.Args[1] // Übergabe des Programmpfades


	// Öffnen und lesen der Bytes des auszuführenden Programms
	dateiInhalt := dateien.Oeffnen(programmPath,'l')
	for !dateiInhalt.Ende(){
		buffer = append(buffer,dateiInhalt.Lesen())
		counter++
	}
	dateiInhalt.Schliessen()


	// Umwandeln der Bytes in EINE Zeichenkette
	hagCode 	:= 	string(buffer[:counter])

	// Umwandeln des Assemblerprogramms in eine opcodeListe. Die Zeilen des Assemblerprogramms
	// und eine Liste der Pseudobefehle wird auch ausgegeben.
	opcodeList,assenblerCodeListe,pseudoCodeListe	:=	opcode.GetOpcodeList(hagCode) 
	fmt.Println(opcodeList)
	fmt.Println(assenblerCodeListe)
	fmt.Println(pseudoCodeListe)
	// Die Key des Hashes "opcodeList" sind die Zeilennummern des Assemblerprogramms diese Keys
	// werden hier sortiert
    var keys []int
    for k := range opcodeList {
        keys = append(keys, k)
    }

    sort.Ints(keys)

	// Deklaration vom Variablen
	var switcher 			bool		//
	var singleOpcode		[]byte
	var opcodeFragment 		byte
	var stopAdresse 		uint16
	var startAdresse 		uint16
	var startAdressePcHEX 	string = pseudoCodeListe["$t@rt@dre$$e"][0]
	//fmt.Println("startAdressePcHEX",startAdressePcHEX)
	var pcHigh 		byte
	var pcLow 		byte
	var pcHighOld 	byte
	var pcLowOld 	byte

	opcodeRegister 		:=  map[string]int{} 		// Enhält als Schlüssel die hexadezimale Nummer der Opcodes 
	showOpcode 			:=  map[string][]string{} 	// Enthält als Schlüssel die Adresse des Opcodes und als Werte den gesammten Opcode 
	showAssemblerCode 	:=  map[string][]string{} 	// Enthält als Schlüssel die Adresse des Opcodes und als Werte den dazugehörigen Assemblercode

	// Die Opcodes werden hier in den Speicher geladen
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
			//fmt.Println(value)
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
	fmt.Println("Register",opcodeRegister)
	fmt.Println(assenblerCodeListe)
	fmt.Println(pseudoCodeListe)
	var getOpcode []byte
	var opcodeHeadAdresse uint16
	var switchStartAdresse bool = true
	var anzahlOpcodeElemente int
	Fenster(1920, 1200)
	var pseudoCodeContentSwitch bool 
	for i := 0; i < 8; i++ {
		
		pseudoCodeContentSwitch  = true // Wird für die Anzeige der Pseudocodes benötigt

		startAdressePcUINT64, err := strconv.ParseUint(startAdressePcHEX, 16, 16) // Die hexadezimale Adresse wird in eine uint64 
																				  // konvertiert 
		//fmt.Println("startAdressePcUINT64",startAdressePcUINT64)
		if err!=nil{
			panic(err)
		}

		// Einmaliges schreiben der Startadressen in den Programmzähler
		if switchStartAdresse{
			startAdressePcUINT16 := uint16(startAdressePcUINT64)
			startAdresseByte := make([]byte, 2)						// Ein Slice mit zwei Byte wird erstellt
			binary.BigEndian.PutUint16(startAdresseByte, startAdressePcUINT16)
			switchStartAdresse = false

			_ = programmZaehlerHigh.SchreibenByte(startAdresseByte[0])
			_ = programmZaehlerLow.SchreibenByte(startAdresseByte[1])
		}

		// Auslesen des Programmzählers
		pcHigh, _	= programmZaehlerHigh.LesenByte()
		pcLow, _	= programmZaehlerLow.LesenByte()

		// Konvertiere byte in uint16
		opcodeHeadAdresse = binary.BigEndian.Uint16([]byte{pcHigh,pcLow})
			//fmt.Println("Ausgelsen opcodeHeadAdresse",opcodeHeadAdresse)
		// Lesen der im Programmzähler angegebenen Adresse in Byte
		getOpcode,	_ = speicher64k.Lesen([]uint16{opcodeHeadAdresse, opcodeHeadAdresse})  // 
			//fmt.Println("Ausgelesene OpcodeHEAD",hex.EncodeToString(getOpcode))
		// Nachschlagen, wie viele weitere Bytes zum Opcode gehören
		anzahlOpcodeElemente =  opcodeRegister[hex.EncodeToString(getOpcode)]
			//fmt.Println("Anzahl", anzahlOpcodeElemente)
			//fmt.Println("Opcode", hex.EncodeToString(getOpcode))
		// Auslesen des gesamten Opcodes
		getOpcode,	_ = speicher64k.Lesen([]uint16{opcodeHeadAdresse, opcodeHeadAdresse+uint16(anzahlOpcodeElemente)-1})  // 
		opcodeHeadAdresse = opcodeHeadAdresse+uint16(anzahlOpcodeElemente )
		opcode.ExecuteOpcode ( 	getOpcode,
								speicher64k, 
								x_register,
								y_register,
								programmZaehlerHigh,
								programmZaehlerLow,
								stapelzeiger, 
								akku, 
								statusbits)

		Stiftfarbe(220, 222, 217)      // Hintergrundfarbe des gesamten Bildschirms
		Vollrechteck(0, 0, 1920, 1200) // Bildschirmhintergrund

		speicher1,	_ = speicher64k.Lesen([]uint16{0		, 255})
		speicher2,	_ = speicher64k.Lesen([]uint16{256		, 511})
		speicher3,	_ = speicher64k.Lesen([]uint16{512		, 767})
		speicher4,	_ = speicher64k.Lesen([]uint16{768		, 1023})
		//speicher4, 	_ = speicher64k.Lesen([]uint16{65280	, 65535})
		UpdateAus()


		// Label AssemblerCode-------------------------------------------------------------------------
		fmt.Println("-->",opcodeHeadAdresse-uint16(anzahlOpcodeElemente))
		fmt.Println("ac",showAssemblerCode)
		gfxElement01.AbbildLabel(10,950,"Assembler-Code",24,0,0,255)
		gfxElement01.AbbildLabel(10,980,strings.Join(showAssemblerCode[strconv.FormatUint(uint64(opcodeHeadAdresse-uint16(anzahlOpcodeElemente)),16)][:]," "),24,255,0,0)

		// Label Opcode-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(10,1010,"Opcode",24,0,0,255)
		gfxElement01.AbbildLabel(10,1040,strings.Join(showOpcode[strconv.FormatUint(uint64(opcodeHeadAdresse-uint16(anzahlOpcodeElemente)),16)][:]," "),24,255,0,0)



		// Label Pseudocode-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(400,950,"Pseudocode",24,0,0,255)
		var counter int
		var y		uint16 = 980
		var x		uint16 = 200
		var rowElements int =7
		for key,list := range pseudoCodeListe{
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
				fmt.Println("x",x)
				fmt.Println("y",y)
				gfxElement01.AbbildLabel(400,y,key+"="+list[0],24,0,0,0)
				x=400
				counter=0
			}else{
				x=x+200
				fmt.Println("x",x)
				fmt.Println("y",y)
				gfxElement01.AbbildLabel(x,y,key+"="+list[0],24,0,0,0)
				counter++
			}
		}


		gfxElement01.AbbildSpeicherseite1(10	, 10	, 0		, speicher1		)	
		gfxElement01.AbbildSpeicherseite1(411	, 10	, 1		, speicher2		)
		gfxElement01.AbbildSpeicherseite1(812	, 10	, 2		, speicher3		)
		gfxElement01.AbbildSpeicherseite1(1213	, 10	, 3		, speicher4		)
		//time.Sleep(1000000000)


		// Label Register-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,10,"Register",24,0,0,255)
		// Anzeigen des X-Registers -------------------------------------------------------------------------
		registerXdaten, _ = x_register.LesenByte()
		gfxElement01.AbbildRegister(1630, 40, "X-Register", registerXdaten, registerXdatenAlt)
		registerXdatenAlt = registerXdaten
		// <<---------------------------------------------------------------------------------------------------

		// Anzeigen des Y-Registers ----------------------------------------------------------------------------
		registerYdaten, _ = y_register.LesenByte()
		gfxElement01.AbbildRegister(1630, 70, "Y-Register", registerYdaten, registerYdatenAlt)
		registerYdatenAlt = registerYdaten
		// <<---------------------------------------------------------------------------------------------------

		// Anzeigen des Akkus ----------------------------------------------------------------------------------
		akkuDaten, _ = akku.LesenByte()
		gfxElement01.AbbildRegister(1630, 100, "Akku", akkuDaten, akkuDatenAlt)
		akkuDatenAlt = akkuDaten
		// <<---------------------------------------------------------------------------------------------------
		
		// Label Stapelzeiger-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,130,"Stapelzeiger",24,0,0,255)
		// Anzeigen des Stapelzeigers --------------------------------------------------------------------------
		stapelzeigerDaten, _ = stapelzeiger.LesenByte()
		gfxElement01.AbbildRegister(1630, 160, "SZ", stapelzeigerDaten, stapelzeigerDatenAlt)
		stapelzeigerDatenAlt = stapelzeigerDaten
		// <<---------------------------------------------------------------------------------------------------

		// Label ProgrammzählerFlags-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,190,"Progammzähler",24,0,0,255)

		// Anzeigen des Programmzählers ----------------------------------------------------------------------------
		pcHigh, _ 	= programmZaehlerHigh.LesenByte()
		pcLow, _ 	= programmZaehlerLow.LesenByte()

		gfxElement01.AbbildRegister(1630, 220, "High", pcHigh, pcHighOld)
		pcHighOld = pcHigh

		gfxElement01.AbbildRegister(1630, 250, "Low", pcLow, pcLowOld)
		pcLowOld = pcLow
		// <<---------------------------------------------------------------------------------------------------
		// <<---------------------------------------------------------------------------------------------------


		// Label Flags-------------------------------------------------------------------------
		gfxElement01.AbbildLabel(1630,280,"Flags",24,0,0,255)

		// Anzeigen der Flags  ---------------------------------------------------------------------------------
		//fmt.Println("-------------------------------------------------------------------------------------")

		for  index,flag := range(flags){
			flagStatusBOOL, _ =statusbits.LeseBit( uint(index))
			//fmt.Println(flagStatusBOOL)	
			//fmt.Println(flag)	
			if flagStatusBOOL {
				flagStatusINT = 1
			}else{
				flagStatusINT = 0
			}
			gfxElement01.AbbildFlag(1630, uint16(310+index*30), flag, flagStatusINT, flagStatusAlt[index])
			flagStatusAlt[index]= flagStatusINT
		}
		// <<---------------------------------------------------------------------------------------------------

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
