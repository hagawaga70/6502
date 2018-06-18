package assembler

import "regexp" 
//import "fmt"
//import "encoding/hex"
//import "encoding/binary"
type impl struct {
}

func NewAssembler() *impl {
	var r *impl
	r = new(impl)
	return r
}
func check8BitAdresse(adresse string)(hex string, hit bool){
	adresse1Byte := regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster
	hit = adresse1Byte.MatchString(adresse)								// - Wendet Suchmuster an. Wenn es passt ist der 
																		//   Rückgabewert "true" 
	adresse = adresse1Byte.ReplaceAllString(adresse,`$1`)				// - Hexadezimaler Wert ohne $-Zeichen 
	hex = adresse
	return
}

func check8BitWert(adresse string)(hex string, hit bool){
	adresse1Byte := regexp.MustCompile(`^\#\$([ABCDEFabcdef0-9]{2})$`)	// - Initialisiere Suchmuster
	hit = adresse1Byte.MatchString(adresse)								// - Wendet Suchmuster an. Wenn es passt ist der 
																		//   Rückgabewert "true" 
	adresse = adresse1Byte.ReplaceAllString(adresse,`$1`)				// - Hexadezimaler Wert ohne # und $-Zeichen 
	hex = adresse
	return
}

func (r *impl) TranslateLDA(assemblerCode []string) ( optcode []string, takte int) {

	if len(assemblerCode) == 2{				//Seite 0 ; noch nicht umgesetzt-> unmittelbar ;(ind,x) ;(ind,y) ; Seite 0,x
		// Adressierungsart: Seite 0
		if hexa,hit := check8BitAdresse(assemblerCode[1]);hit == true{	
			optcode = append(optcode, "A5") 			// 101bbb01 1010 0101
			optcode = append(optcode, hexa) 			// Hexadezimale Zahl ohne Dollarzeichen
			takte = 3									// Anzahl der benötigten Takte
		// Adressierungsart: unmittelbar
		}else if hexa,hit := check8BitWert(assemblerCode[1]);hit == true{	
			optcode = append(optcode, "A9") 			// 101bbb01 1010 1001
			optcode = append(optcode, hexa) 			// Hexadezimale Zahl ohne Dollarzeichen
			takte = 2									// Anzahl der benötigten Takte
		// Fehlermeldung
		}else{
			optcode = append(optcode, "Fehler bei der Übersetzung des Befehls LDA")
			optcode = append(optcode, "001: Das zweite Byte entspricht nicht der Anforderungen z.B. $0A oder $0a" )
			takte = -1
		}
	}else if len(assemblerCode) == 3{
		// Absolut	
		if	hexa1,hit1,hexa2,hit2 := check8BitwerteTwice(assemblerCode[1],assemblerCode[2]);    //
			hit1 == true && hit2 == true{
				optcode = append(optcode, "AD")			// 101bbb01 1010 1101
				optcode = append(optcode, hexa1)		// Hexadezimale Zahl ohne Dollarzeichen
				optcode = append(optcode, hexa2)		// Hexadezimale Zahl ohne Dollarzeichen
				takte = 4								// Anzahl der benötigten Takte
		// Fehlermeldung
		}else{
			optcode = append(optcode, "Fehler bei der Übersetzung des Befehls LDA")
			optcode = append(optcode, "001: Das 2. oder 3. Byte entspricht nicht der Anforderungen z.B. $0A oder $0a" )
			takte = -1
		}
	}
	return optcode,takte
}

func check8BitwerteTwice(adress1 string,adress2 string)(hexa1 string,hit1 bool, hexa2 string, hit2 bool){
	hexa1,hit1 = check8BitAdresse(adress1)
	hexa2,hit2 = check8BitAdresse(adress2)
	return
}


