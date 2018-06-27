package assembler

import "regexp"
import ."fmt"
import "strconv"


type impl struct {
}

func NewAssembler() *impl {
	var r *impl
	r = new(impl)
	return r
}
/*
Noch umzusetzende Befehle: 
befehleListe := map[string][]string{
		"AND":{},
		"ASL":{},
		"BCC":{},
		"BCS":{},
		"BEQ":{},
		"BIT":{},
		"BMI":{},
		"BNE":{},
		"BPL":{},
		"BRK":{},
		"BVC":{},
		"BVS":{},
		"CLI":{},
		"CMP":{},
		"CLV":{},
		"CPX":{},
		"CPY":{},
		"DEC":{},
		"DEX":{},
		"DEY":{},
		"EOR":{},
		"INC":{},
		"INX":{},
		"INY":{},
		"JMP":{},
		"JSR":{},
		"LSR":{},
		"NOP":{},
		"ORA":{},
		"PHA":{},
		"PHP":{},
		"PLA":{},
		"PLP":{},
		"ROL":{},
		"ROR":{},
		"RTI":{},
		"RTS":{},
		"SBC":{},
		"SEC":{},
		"SED":{},
		"SEI":{},
		"TAX":{},
		"TAY":{},
		"TSX":{},
		"TXS":{},
		"TXA":{},
		"TYA":{}}
*/
// Vor.:  -
// Erg.: Ein Array mit einem oder zwei Elementen ist zurückgegeben. Dabei handelt es sich um Strings mit hexadezimalen
// 		 Zahlen. Das Array kann ein oder zwei Elemente haben. Zurückgegeben werden 1-bytige oder 2-bytige Adressen. Die
//		 Inhalte von Pseudobefehlen/Sprungmarken sind aufgelöst.

func checkAdresse(adresse string,pseudoBefehle map[string][]string)(hex []string, hit bool){
	adresse1Byte1Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{1})$`)	// - Initialisiere Suchmuster 1 Bit einstellig
	adresse1Byte2Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 1 Bit zweistellig
	adresse2Byte3Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{1})([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 2 Byte dreistellig
	adresse2Byte4Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 2 Byte vierstellig
	//adresse2Byte	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{3,4})$`)	// - Initialisiere Suchmuster 2 Bit
	adressePCByte	:= regexp.MustCompile(`^([A-Za-z0-9])+$`)			// - Initialisiere Suchmuster PseudoCode


	// Es wurde ein Pseudocode übergeben
	if hit= adressePCByte.MatchString(adresse);hit{							// - Ließt den Wert des PseudoCodes aus
		adresse = pseudoBefehle[adresse][0]									//
	}

	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 1-stellig (die führende Null wird davorgesetzt)
	if hit= adresse1Byte1Stellig.MatchString(adresse);hit{
		adresse = adresse1Byte1Stellig.ReplaceAllString(adresse,`$1`)
		adresse = "0"+adresse
		hex		= append(hex,adresse)

	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 2-stellig 
	}else if hit= adresse1Byte2Stellig.MatchString(adresse);hit{
		adresse = adresse1Byte2Stellig.ReplaceAllString(adresse,`$1`)
		hex		= append(hex,adresse)

	// - Hexadezimaler Wert > 255 ohne $-Zeichen (die führende Null wird davorgesetzt) 
	}else if hit= adresse2Byte3Stellig.MatchString(adresse);hit{
		hex		= append(hex,"0"+adresse2Byte3Stellig.ReplaceAllString(adresse,`$1`))
		hex		= append(hex,adresse2Byte3Stellig.ReplaceAllString(adresse,`$2`))

	// - Hexadezimaler Wert > 255 ohne $-Zeichen 
	}else if hit= adresse2Byte4Stellig.MatchString(adresse);hit{
		hex		= append(hex,adresse2Byte3Stellig.ReplaceAllString(adresse,`$1`))	
		hex		= append(hex,adresse2Byte3Stellig.ReplaceAllString(adresse,`$2`))	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 
	}

	return

}
// Vor.:  -
// Erg.: Ein Array mit einem  Elementen ist zurückgegeben. Dabei handelt es sich um einen String mit einer hexadezimalen
// 		 Zahl. Passt der reguläre Ausdruck nicht zum String in der Variablen Adresse ist ein als boolscher Wert ein 
// 		 false (hit) zurückgegeben. Die
//		 Inhalte von Pseudobefehlen/Sprungmarken sind aufgelöst.


func check8BitWert(adresse string,pseudoBefehle map[string][]string)(hex []string, hit bool){

	adresse1Byte1Stellig := regexp.MustCompile(`^\#\$([ABCDEFabcdef0-9]{1})$`)		// - Initialisiere Suchmuster
	adresse1Byte2Stellig := regexp.MustCompile(`^\#\$([ABCDEFabcdef0-9]{2})$`)		// - Initialisiere Suchmuster
	adressePCByte	:= regexp.MustCompile(`^([A-Za-z0-9])+$`)			// - Initialisiere Suchmuster PseudoCode

	if hit= adressePCByte.MatchString(adresse);hit{							// - Ließt den Wert des PseudoCodes aus
		adresse = pseudoBefehle[adresse][0]									//
	}

	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 1-stellig 
	if hit= adresse1Byte1Stellig.MatchString(adresse);hit{
		adresse = adresse1Byte1Stellig.ReplaceAllString(adresse,`$1`)
		adresse = "0"+adresse
		hex		= append(hex,adresse)

	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 2-stellig 
	}else if hit= adresse1Byte2Stellig.MatchString(adresse);hit{
		adresse = adresse1Byte2Stellig.ReplaceAllString(adresse,`$1`)
		hex		= append(hex,adresse)
	}

	//hit = adresse1Byte.MatchString(adresse)								// - Wendet Suchmuster an. Wenn es passt ist der 
	//adresse = adresse1Byte.ReplaceAllString(adresse,`$1`)					// - Hexadezimaler Wert ohne # und $-Zeichen 
	//hex = append(hex,adresse)
	return
}



func (r *impl) TranslateXXX(assemblerCode []string,pseudoBefehle map[string][]string, aktuelleAdresse string) ( optcode []string, err bool, naechsteAdresse string) {

	var adressOffset int

	// <ADRL> Folgende Parameter werden in der Befehlsliste pro Befehl abgelegt
	// IMPLIZIT AKKUMULATOR ABSOLUT SEITE0 UNMITTELBAR ABS.X ABS.Y (IND,X) (IND,Y) SEITE0.Y RELATIV INDIREKT

	befehleListe := map[string][]string{
		"#€°":{	"IMPLIZIT"		,
				"AKKUMULATOR"	,
				"ABSOLUT"		,
				"SEITE0"		,
				"UNMITTELBAR"	,
				"ABS.X"			,
				"ABS.Y"			,
				"(IND,X)"		,
				"(IND,Y)"		,
				"SEITE0.X"		,
				"SEITE0.Y"		,
				"RELATIV"		,
				"INDIREKT"}, 	//Umgesetzt: --11100000---
		"ADC":{"--","--","6d","65","69","7d","79","61","71","75","--","--","--"}, 	//Umgesetzt: --11100000---
		"JMP":{"--","--","4c","--","--","--","--","--","--","--","--","--","6c"}, 	//Umgesetzt: --1---------0
		"LDA":{"--","--","ad","a5","a9","bd","b9","a1","b1","b5","--","--","--"}, 	//Umgesetzt: --11100000---
		"LDX":{"--","--","ae","a6","a2","--","be","--","--","--","b6","--","--"},	//Umgesetzt: --111-0---0--
		"LDY":{"--","--","ac","a4","a0","bc","--","--","--","b4","--","--","--"},	//Umgesetzt: --1110---0---
		"STA":{"--","--","8d","85","--","9d","99","81","91","95","--","--","--"}, 	//Umgesetzt: --11-00000---
		"STX":{"--","--","8e","86","--","--","--","--","--","--","96","--","--"}, 	//Umgesetzt: --111----0---
		"STY":{"--","--","8c","84","--","--","--","--","--","--","94","--","--"}} 	//Umgesetzt: --111----0---

	// Überprüft die zum Assemblercode übergebene Konstante oder Adresse. hexa ist ein Slice mit ein oder
	// zwei Elementen (1 oder zwei Bytes)
	if hexa,hit := checkAdresse(assemblerCode[1],pseudoBefehle);hit == true{	 
		// Seite 0
		if len(hexa)==1{
			// Bei einer falschen Addressierung wird das Programm beendet
			if befehleListe[assemblerCode[0]][3] == "--"{
				panic(assemblerCode[0]+" kann nicht "+befehleListe["#€°"][3]+" addressiert werden!!!")
			}
			optcode = append(optcode, befehleListe[assemblerCode[0]][3]) 			// 
			optcode = append(optcode, hexa[0]) 										// Hexadezimale Zahl ohne Dollarzeichen
			adressOffset=len(optcode)												// x Byte bis zur nächsten freien Adresse

		// Absolut	
		}else if len(hexa)==2{

			// Bei einer falschen Addressierung wird das Programm beendet
			if befehleListe[assemblerCode[0]][2] == "--"{
				panic(assemblerCode[0]+" kann nicht "+befehleListe["#€°"][2]+" addressiert werden!!!")
			}
			optcode = append(optcode, befehleListe[assemblerCode[0]][2]) 			// 
			optcode = append(optcode, hexa[0])										// Hexadezimale Zahl ohne Dollarzeichen
			optcode = append(optcode, hexa[1])										// Hexadezimale Zahl ohne Dollarzeichen
			adressOffset=len(optcode)												// x Byte bis zur nächsten freien Adresse
		}else{

			optcode = append(optcode, "LDA: Fehler der absoluten Addressierung oder der Addressierug der Seite 0 >>"+assemblerCode[1]+"<")	 // 101bbb01 1010 1101
			err = true
			adressOffset=0
		}

	// Adressierungsart: unmittelbar
	}else if hexa,hit := check8BitWert(assemblerCode[1],pseudoBefehle);hit == true{	

		// Bei einer falschen Addressierung wird das Programm beendet
		if befehleListe[assemblerCode[0]][4] == "--"{
			panic(assemblerCode[0]+" kann nicht "+befehleListe["#€°"][4]+" addressiert werden!!!")
		}
		optcode = append(optcode, befehleListe[assemblerCode[0]][4]) 			// 
		optcode = append(optcode, hexa[0]) 										// Hexadezimale Zahl ohne Dollarzeichen
		adressOffset=len(optcode)												// x Byte bis zur nächsten freien Adresse

	// Fehlermeldung wenn der Adress- bzw Werteil nicht ausgewertet werden konnte. z.B bei falschen Zeichen!!
	}else{

		optcode = append(optcode, "Fehler bei der  Addressierung  >>"+assemblerCode[1]+"<")	 // 101bbb01 1010 1101
		err = true
		adressOffset=0
	}

	aktuelleAdresseINT, erro := strconv.ParseInt(aktuelleAdresse, 16, 0)

	if erro != nil {

		optcode = append(optcode, assemblerCode[0]+": Fehler bei der Konvertierung HEX -> INT")	 // 101bbb01 1010 1101
		err = true
		adressOffset=0
	}

	// Berechnung der nächsten Adresse, die ein opcodeHEAD (erstes Element des opcodes) enthält
	naechsteAdresseINT	:= int(aktuelleAdresseINT) + adressOffset
	naechsteAdresse		 = strconv.FormatInt(int64(naechsteAdresseINT), 16)

	return optcode, err, naechsteAdresse
}


func (r *impl) TranslateModifyFlags(assemblerCode []string, aktuelleAdresse string)(optcode []string, err bool,naechsteAdresse string){

	var adressOffset int
	befehleListe := map[string][]string{
		"CLC":{"18"},
		"CLD":{"d8"}}

	// Der opcode besteht nur aus dem OpcodeHEAD
	optcode 	 = append(optcode, befehleListe[assemblerCode[0]][0]) 			// 
	adressOffset = len(optcode)													// x Byte bis zur nächsten freien Adresse

	aktuelleAdresseINT, erro := strconv.ParseInt(aktuelleAdresse, 16, 0)
	if erro != nil {

		optcode = append(optcode, "LDA: Fehler bei der Konvertierung HEX -> INT")	 // 101bbb01 1010 1101
		err = true
		adressOffset=0
	}

	naechsteAdresseINT	:= int(aktuelleAdresseINT) + adressOffset
	naechsteAdresse		 = strconv.FormatInt(int64(naechsteAdresseINT), 16)
	return optcode, err, naechsteAdresse
}

func (r *impl) TranslateEnd(assemblerCode []string, aktuelleAdresse string)(optcode []string, err bool,naechsteAdresse string){

	var adressOffset int
	befehleListe := map[string][]string{"END":{"f2"}}
	// Das erste Element des Slice assemblerCode enthält den Befehlname
	Println("ac ",assemblerCode)	
	optcode 	 = append(optcode, befehleListe[assemblerCode[0]][0]) 			// 
	adressOffset = len(optcode)													// x Byte bis zur nächsten freien Adresse

	aktuelleAdresseINT, erro := strconv.ParseInt(aktuelleAdresse, 16, 0)
	if erro != nil {

		optcode = append(optcode, "LDA: Fehler bei der Konvertierung HEX -> INT")	 // 101bbb01 1010 1101
		err = true
		adressOffset=0
	}

	naechsteAdresseINT	:= int(aktuelleAdresseINT) + adressOffset
	naechsteAdresse		 = strconv.FormatInt(int64(naechsteAdresseINT), 16)
	return optcode, err, naechsteAdresse
}

