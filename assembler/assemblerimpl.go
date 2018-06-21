package assembler

import "regexp"
import ."fmt"
//import "encoding/hex"
//import "encoding/binary"
import "strconv"


type impl struct {
}

func NewAssembler() *impl {
	var r *impl
	r = new(impl)
	return r
}
/*
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
func checkAdresse(adresse string,pseudoBefehle map[string][]string)(hex []string, hit bool){
	var debug bool = true
	adresse1Byte1Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{1})$`)	// - Initialisiere Suchmuster 1 Bit einstellig
	adresse1Byte2Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 1 Bit zweistellig
	adresse2Byte3Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{1})([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 2 Bit dreistellig
	adresse2Byte4Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 2 Bit vierstellig
	//adresse2Byte	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{3,4})$`)	// - Initialisiere Suchmuster 2 Bit
	adressePCByte	:= regexp.MustCompile(`^([A-Za-z0-9])+$`)			// - Initialisiere Suchmuster PseudoCode

	if debug{

			Println("---------------------------------------------------------------------")
			Println("assemblerimpl -pseudoBefehle -adresse")
			Println(pseudoBefehle)
			Println(adresse)
			Println("°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")

	}


	// Es wurde ein Pseudocode übergeben
	if hit= adressePCByte.MatchString(adresse);hit{							// - Ließt den Wert des PseudoCodes aus
		adresse = pseudoBefehle[adresse][0]									//
		if debug{Println("assemblerimpl -machtadresse"); Println(adresse);Println("END")}
	}


	if debug{

			Println("---------------------------------------------------------------------")
			Println("assemblerimpl -adresse")
			Println(adresse)
			Println("°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")

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
/*
																		//   Rückgabewert "true" 
	hit = adresse1Byte.MatchString(adresse)								// - Wendet Suchmuster an. Wenn es passt ist der 
																		//   Rückgabewert "true" 
*/
}

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
	// Entspricht die Syntax der absoluten Addressierung oder der Addressierug der Seite 0
	var adressOffset int

	// Folgende Parameter werden in der Befehlsliste pro Befehl abgelegt
	// IMPLIZIT AKKUMULATOR ABSOLUT SEITE0 UNMITTELBAR ABS.X ABS.Y (IND,X) (IND,Y) SEITE0.Y RELATIV INDIREKT
	befehleListe := map[string][]string{
		"ADC":{"--","--","6D","65","69","7D","79","61","71","75","--","--","--"}, 	//Umgesetzt: --11100000---
		"LDA":{"--","--","AD","A5","A9","BD","B9","A1","B1","B5","--","--","--"}, 	//Umgesetzt: --11100000---
		"LDX":{"--","--","AE","A6","A2","--","BE","--","--","--","B6","--","--"},	//Umgesetzt: --111-0---0--
		"LDY":{"--","--","AC","A4","A0","BC","--","--","--","B4","--","--","--"},	//Umgesetzt: --1110---0---
		"STA":{"--","--","8D","85","--","9D","99","81","91","95","--","--","--"}, 	//Umgesetzt: --11-00000---
		"STX":{"--","--","8E","86","--","--","--","--","--","--","96","--","--"}, 	//Umgesetzt: --111----0---
		"STY":{"--","--","8C","84","--","--","--","--","--","--","94","--","--"}} 	//Umgesetzt: --111----0---


	if hexa,hit := checkAdresse(assemblerCode[1],pseudoBefehle);hit == true{	 
		// Seite 0
		if len(hexa)==1{
			optcode = append(optcode, befehleListe[assemblerCode[0]][3]) 			// 
			optcode = append(optcode, hexa[0]) 										// Hexadezimale Zahl ohne Dollarzeichen
			adressOffset=len(optcode)												// x Byte bis zur nächsten freien Adresse

		// Absolut	
		}else if len(hexa)==2{
			optcode = append(optcode, befehleListe[assemblerCode[0]][2]) 			// 
			optcode = append(optcode, hexa[0])										// Hexadezimale Zahl ohne Dollarzeichen
			optcode = append(optcode, hexa[1])										// Hexadezimale Zahl ohne Dollarzeichen
			adressOffset=len(optcode)												// x Byte bis zur nächsten freien Adresse

			//optcode = append(optcode, "Fehler bei der Übersetzung des Befehls LDA")
			//optcode = append(optcode, "001: Das zweite Byte entspricht nicht der Anforderungen z.B. $0A oder $0a" )
			//takte = -1
		}else{

			optcode = append(optcode, "LDA: Fehler der absoluten Addressierung oder der Addressierug der Seite 0 >>"+assemblerCode[1]+"<")	 // 101bbb01 1010 1101
			err = true
			adressOffset=0
		}

	// Adressierungsart: unmittelbar
	}else if hexa,hit := check8BitWert(assemblerCode[1],pseudoBefehle);hit == true{	
		optcode = append(optcode, befehleListe[assemblerCode[0]][4]) 			// 
		optcode = append(optcode, hexa[0]) 										// Hexadezimale Zahl ohne Dollarzeichen
		adressOffset=len(optcode)												// x Byte bis zur nächsten freien Adresse

	// Fehlermeldung
	}else{

		optcode = append(optcode, "LDA: Fehler der unmittelbaren Addressierung  >>"+assemblerCode[1]+"<")	 // 101bbb01 1010 1101
		err = true
		adressOffset=0
	}

	aktuelleAdresseINT, erro := strconv.ParseInt(aktuelleAdresse, 16, 0)

	if erro != nil {

		optcode = append(optcode, assemblerCode[0]+": Fehler bei der Konvertierung HEX -> INT")	 // 101bbb01 1010 1101
		err = true
		adressOffset=0
	}


	naechsteAdresseINT	:= int(aktuelleAdresseINT) + adressOffset
	naechsteAdresse		 = strconv.FormatInt(int64(naechsteAdresseINT), 16)

	return optcode, err, naechsteAdresse
}


func (r *impl) TranslateModifyFlags(assemblerCode []string, aktuelleAdresse string)(optcode []string, err bool,naechsteAdresse string){

	var adressOffset int
	befehleListe := map[string][]string{
		"CLC":{"18"},
		"CLD":{"D8"}}
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

