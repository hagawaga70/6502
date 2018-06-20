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
	//adresse = adresse1Byte.ReplaceAllString(adresse,`$1`)				// - Hexadezimaler Wert ohne # und $-Zeichen 
	//hex = append(hex,adresse)
	return
}

func (r *impl) TranslateLDA	(assemblerCode []string,pseudoBefehle map[string][]string, aktuelleAdresse string) ( optcode []string, takte string, naechsteAdresse string) {
	// Entspricht die Syntax der absoluten Addressierung oder der Addressierug der Seite 0
	var adressOffset int
	if hexa,hit := checkAdresse(assemblerCode[1],pseudoBefehle);hit == true{	 
		// Seite 0
		if len(hexa)==1{
			optcode = append(optcode, "A5") 			// 101bbb01 1010 0101
			optcode = append(optcode, hexa[0]) 			// Hexadezimale Zahl ohne Dollarzeichen
			takte = "3"									// Anzahl der benötigten Takte
			adressOffset=3								// Zwei Byte bis zur nächsten freien Adresse

		// Absolut	
		}else if len(hexa)==2{
			optcode = append(optcode, "AD")				// 101bbb01 1010 1101
			optcode = append(optcode, hexa[0])			// Hexadezimale Zahl ohne Dollarzeichen
			optcode = append(optcode, hexa[1])			// Hexadezimale Zahl ohne Dollarzeichen
			takte = "4"									// Anzahl der benötigten Takte
			adressOffset=4								// Zwei Byte bis zur nächsten freien Adresse

			//optcode = append(optcode, "Fehler bei der Übersetzung des Befehls LDA")
			//optcode = append(optcode, "001: Das zweite Byte entspricht nicht der Anforderungen z.B. $0A oder $0a" )
			//takte = -1
		}else{

			optcode = append(optcode, "LDA: Fehler der absoluten Addressierung oder der Addressierug der Seite 0 >>"+assemblerCode[1]+"<")	 // 101bbb01 1010 1101
			takte = "-1"			// Wenn takte = -1  -> FEHLER
			adressOffset=0
		}

	// Adressierungsart: unmittelbar
	}else if hexa,hit := check8BitWert(assemblerCode[1],pseudoBefehle);hit == true{	
		optcode = append(optcode, "A9") 			// 101bbb01 1010 1001
		optcode = append(optcode, hexa[0]) 			// Hexadezimale Zahl ohne Dollarzeichen
		takte 	= "2"									// Anzahl der benötigten Takte
		adressOffset=3								// Ein Byte bis zur nächsten freien Adresse

	// Fehlermeldung
	}else{

		optcode = append(optcode, "LDA: Fehler der unmittelbaren Addressierung  >>"+assemblerCode[1]+"<")	 // 101bbb01 1010 1101
		takte = "-1"			// Wenn takte = -1  -> FEHLER
		adressOffset=0
	}

	aktuelleAdresseINT, err := strconv.ParseInt(aktuelleAdresse, 16, 0)

	if err != nil {

		optcode = append(optcode, "LDA: Fehler bei der Konvertierung HEX -> INT")	 // 101bbb01 1010 1101
		takte = "-1"			// Wenn takte = -1  -> FEHLER
		adressOffset=0
	}


	naechsteAdresseINT	:= int(aktuelleAdresseINT) + adressOffset
	naechsteAdresse		 = strconv.FormatInt(int64(naechsteAdresseINT), 16)

	return optcode, takte, naechsteAdresse
}



