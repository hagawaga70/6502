package assembler

import "regexp"
//import "fmt"
import "encoding/hex"
//import "encoding/binary"
type impl struct {
}

func NewAssembler() *impl {
	var r *impl
	r = new(impl)
	return r
}
func checkAdresse(adresse string,pseudoBefehle map[string][]string{})(hex []string, hit bool){
	adresse1Byte1Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{1})$`)	// - Initialisiere Suchmuster 1 Bit einstellig
	adresse1Byte2Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 1 Bit zweistellig
	adresse2Byte3Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{1})([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 2 Bit dreistellig
	adresse2Byte4Stellig	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster 2 Bit vierstellig
	adresse2Byte	:= regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{3,4})$`)	// - Initialisiere Suchmuster 2 Bit
	adressePCByte	:= regexp.MustCompile(`^([ABCDEFabcdef0-9]+$`)			// - Initialisiere Suchmuster PseudoCode

	if hit= adressePCByte.MatchString(adresse);hit{							// - Ließt den Wert des PseudoCodes aus
		adresse = pseudoBefehle[adresse]									//
	}

	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 1-stellig (die führende Null wird davorgesetzt)
	if hit= adresse1Byte1Stellig.MatchString(adresse);hit{
		adresse = adresse1Byte1Stellig.ReplaceAllString(adresse,`$1`)
		adresse = "0"+adresse
		hex		= append(hex,adresse)

	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 2-stellig 
	}else if hit= adresse1Byte2Stellig.MatchString(adresse);hit{
		adresse = adresse1Byte1Stellig.ReplaceAllString(adresse,`$1`)
		hex		= append(hex,adresse)

	// - Hexadezimaler Wert > 255 ohne $-Zeichen (die führende Null wird davorgesetzt) 
	}else if hit= adresse2Byte3Stellig.MatchString(adresse);hit{
		hex		= append(hex,"0"+adresse2Byte3Stellig.ReplaceAllString(adresse,`$1`))
		hex		= append(hex,adresse2Byte3Stellig.ReplaceAllString(adresse,`$2`))

	// - Hexadezimaler Wert > 255 ohne $-Zeichen 
	}else if hit= adresse2Byte4Stellig.MatchString(adresse);hit{
		hex		= append(hex,adresse2Byte3Stellig.ReplaceAllString(adresse,`$1`))	
		hex		= append(hex,adresse2Byte3Stellig.ReplaceAllString(adresse,`$2`))	// - Hexadezimaler Wert <= 255 ohne $-Zeichen 

	}else{
		panic("003: Falsche Angabe der Adresse")
	}


	return
/*
																		//   Rückgabewert "true" 
	hit = adresse1Byte.MatchString(adresse)								// - Wendet Suchmuster an. Wenn es passt ist der 
																		//   Rückgabewert "true" 
*/
}

func check8BitWert(adresse string)(hex string, hit bool){
	adresse1Byte := regexp.MustCompile(`^\#\$([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster
	hit = adresse1Byte.MatchString(adresse)								// - Wendet Suchmuster an. Wenn es passt ist der 
	adresse = adresse1Byte.ReplaceAllString(adresse,`$1`)				// - Hexadezimaler Wert ohne # und $-Zeichen 
	hex = adresse
	return
}

func (r *impl) TranslateLDA	(assemblerCode []string,pseudoBefehle map[string][]string{}, aktuelleAdresse string) 
							( optcode []string, takte int, naechsteAdresse string) {

	if hexa,hit := checkAdresse(assemblerCode[1]);hit == true{	
		// Seite 0
		if len(hex)==1{
			optcode = append(optcode, "A5") 			// 101bbb01 1010 0101
			optcode = append(optcode, hexa) 			// Hexadezimale Zahl ohne Dollarzeichen
			takte = 3									// Anzahl der benötigten Takte

		// Absolut	
		else if len(hex)==2{
			optcode = append(optcode, "AD")				// 101bbb01 1010 1101
			optcode = append(optcode, hexa)				// Hexadezimale Zahl ohne Dollarzeichen
			takte = 4									// Anzahl der benötigten Takte

			//optcode = append(optcode, "Fehler bei der Übersetzung des Befehls LDA")
			//optcode = append(optcode, "001: Das zweite Byte entspricht nicht der Anforderungen z.B. $0A oder $0a" )
			//takte = -1
		}else{
			panic("Fehler 010")
		}
	else if hexa,hit := checkAdresse(assemblerCode[1]);hit == true{	
	}else if hexa1,hit1,hexa2,hit2 := checkWerteTwice(assemblerCode[1],assemblerCode[2]);    //
			hit1 == true && hit2 == true{
						// Fehlermeldung
		}else{
			optcode = append(optcode, "Fehler bei der Übersetzung des Befehls LDA")
			optcode = append(optcode, "002: Das 2. oder 3. Byte entspricht nicht der Anforderungen z.B. $0A oder $0a" )
			takte = -1
		// Adressierungsart: unmittelbar
		}else if hexa,hit := check8BitWert(assemblerCode[1]);hit == true{	
			optcode = append(optcode, "A9") 			// 101bbb01 1010 1001
			optcode = append(optcode, hexa) 			// Hexadezimale Zahl ohne Dollarzeichen
			takte = 2									// Anzahl der benötigten Takte
		// Fehlermeldung
}
	return optcode,takte
}

func check8BitwerteTwice(adress1 string,adress2 string)(hexa1 string,hit1 bool, hexa2 string, hit2 bool){
	hexa1,hit1 = check8BitAdresse(adress1)
	hexa2,hit2 = check8BitAdresse(adress2)
	return
}


