package assembler

import "regexp" 
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
	adresse1Byte := regexp.MustCompile(`^\$([ABCDEFabcdef0-9]{2})$`)
	hit = adresse1Byte.MatchString(adresse)
	hex = adresse1Byte.ReplaceAllString(hex,`$1`)
	return
}
func (r *impl) TranslateLDA(assemblerCode []string) ( optcode []string, takte int) {

	if len(assemblerCode) == 2{				//Seite 0 ; noch nicht umgesetzt-> unmittelbar ;(ind,x) ;(ind,y) ; Seite 0,x
		if hexa,hit := check8BitAdresse(assemblerCode[1]);hit == true{	// Adressierungsart: Seite 0
			optcode = append(optcode, "C5")
			optcode = append(optcode, hexa)
			optcode[1] = hexa 
			takte = 3
		}else{
			optcode = append(optcode, "Fehler bei der Übersetzung des Befehls LDA") 
			optcode = append(optcode, "001: Das zweite Byte entspricht nicht der Anforderungen z.B. $0A oder $0a" )
			takte = -1
		}
	}
	return optcode,takte
}


