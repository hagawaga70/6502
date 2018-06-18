package main



import "./dateien"
import "fmt"
import "regexp"
import "strings"
//import "reflect"
//import ."./assembler"

var buffer 		[]byte
var counter 	int
func main(){
	// Liste der Assemblerbefehle
	befehleListe := map[string][]string{
	"ADC":{},"AND":{},"ASL":{},"BCC":{},"BCS":{},"BEQ":{},
	"BIT":{},"BMI":{},"BNE":{},"BPL":{},"BRK":{},"BVC":{},
	"BVS":{},"CLC":{},"CLD":{},"CLI":{},"CMP":{},"CLV":{},
	"CPX":{},"CPY":{},"DEC":{},"DEX":{},"DEY":{},"EOR":{},
	"INC":{},"INX":{},"INY":{},"JMP":{},"JSR":{},"LDA":{},
	"LDX":{},"LDY":{},"LSR":{},"NOP":{},"ORA":{},"PHA":{},
	"PHP":{},"PLA":{},"PLP":{},"ROL":{},"ROR":{},"RTI":{},
	"RTS":{},"SBC":{},"SEC":{},"SED":{},"SEI":{},"STA":{},
	"STX":{},"STY":{},"TAX":{},"TAY":{},"TSX":{},"TXS":{},
	"TXA":{},"TYA":{}}
	// Öffenen der Programmdatei
	dateiInhalt := dateien.Oeffnen("./programm.hag",'l')
	for !dateiInhalt.Ende(){
		buffer = append(buffer,dateiInhalt.Lesen())
		counter++
	}
	dateiInhalt.Schliessen()

	// Umwandeln der Bytes in EINE Zeichenkette
	hagCode := string(buffer[:counter])

	// Umwandeln der EINEN Zeichenkette in ein Slice: Das Trennzeichen ist \n (NEWLINE)
	// Das Array hat am Ende ein Leerzeile mehr als die Programmdatei (Kein Problem)
	hagCodeArray := strings.Split(hagCode,"\n")

	// Finde führende und abschliessende Leerzeichen
	regexLeerzeichen := regexp.MustCompile(`^\s*(.*)\s*$`)

	// Finde  reine Kommentarzeilen
	jumpComment := regexp.MustCompile(`^\s*;.+$`)

	// Finde Kommentare
	deleteComment := regexp.MustCompile(`^(.*);.*$`)

	// Finde leere Zeilen
	ueberspringeLeereZeilen := regexp.MustCompile(`^\s*$`)


	// Finde Zeilen mit n  Semikolons 
	ueberspringeSemikolonZeilen := regexp.MustCompile(`^\s*;+\s*$`)

	// Finde Pseudobefehle 
	pseudoBefehle 		:= regexp.MustCompile(`^\s*[A-Za-z0-9]+\s*=\s*(#\$|\$)[A-Fa-f0-9]+\s*$`)
	pseudoBefehleExakt 	:= regexp.MustCompile(`^\s*([A-Z0-9]+)\s*=\s*((#\$|\$)[0-9A-Fa-f]{1,4})\s*$`)

	pseudoBefehleHASH := map[string][]string{}



	//var assemble Assembler = NewAssembler() 
	var codeLine string
	for _,codeLine = range(hagCodeArray){
			fmt.Println(codeLine)

		// Überspringe leere Zeilen
		if ueberspringeLeereZeilen.MatchString(codeLine){
			continue
		}

		// Überspringe Semikolonzeilen 
		if ueberspringeSemikolonZeilen.MatchString(codeLine){
			continue
		}

		// Überspringe reine Kommentarzeilen
		if jumpComment.MatchString(codeLine){
			continue
		}
		// Löschen der Leezeichen
		if regexLeerzeichen.MatchString(codeLine){
			codeLine = regexLeerzeichen.ReplaceAllString(codeLine, `$1`)
		}

		// Löschen der Kommentare
		if  deleteComment.MatchString(codeLine){
			codeLine = deleteComment.ReplaceAllString(codeLine, `$1`)
		}

		if  pseudoBefehle.MatchString(codeLine){
			if 	pseudoBefehleExakt.MatchString(codeLine){

				pseudoBefehlName	:= pseudoBefehleExakt.ReplaceAllString(codeLine, `$1`)
				pseudoBefehlInhalt	:= pseudoBefehleExakt.ReplaceAllString(codeLine, `$2`)
				pseudoBefehleHASH[pseudoBefehlName] = []string{pseudoBefehlInhalt}
				continue
			}else{
				panic("001:Die Zuweisung des Pseudocodes entspricht nicht den Anforderungen: z.B ADR1=$1111")
			}
		}
		fmt.Println(codeLine)
		array:=strings.Fields(codeLine)
		if _, ok := befehleListe[array[0]]; ok {
			fmt.Println("value: ", array[0])
		} else {
			fmt.Println("key not found")
		}
		fmt.Println(array)



/*

		for _,line := range(codeArray){	
			opcode,takte := assemble.TranslateLDA(codeArray)
			for _,value := range(opcode){
				fmt.Println(value)
			}
			fmt.Println(takte)
		}
*/
	}
        fmt.Println(pseudoBefehleHASH)
}

/*
// Create map of string slices.
    m := map[string][]string{
						// 101b bb01								
        "LDA": {"--",	// 			 IMPLIZIT
				"--",	// 			 AKKUMULATOR	
				"AD",	// 1010 1101 ABSOLUT		<-realisert
				"A5",	// 1010 0101 SEITE 0 		<-realisert
				"A9",	// 1010 1001 UNMITTELBAR	<-realisert
				"BD",	// 1011 1101 ABS.X
				"B9",	// 1011 1001 ABS.Y
				"A1",	// 1010 0001 (IND,X)
				"B1",	// 1011 0001 (IND,Y)
				"B5",	// 1011 0101 SEITE 0,X
				"--",	// 			 SEITE 0,Y
				"--",	// 			 RELATIV
				"--",},	// 			 INDIREKT
        "AND": {},
    }

// Add a string at the dog key.
    // ... Append returns the new string slice.
    res := append(m["dog"], "brown")

    // Add a key for fish.
    m["fish"] = []string{"orange", "red"}

    // Print slice at key.
    fmt.Println(m["fish"])

    // Loop over string slice at key.
    for i := range m["fish"] {
        fmt.Println(i, m["fish"][i])
    }


func testBefehl(befehl string) bool{
	var befehle []string  =  	[]string{		","AND","ASL","BCC","BCS","BEQ",
												"BIT","BMI","BNE","BPL","BRK","BVC",
												"BVS","CLC","CLD","CLI","CMP","CLV",
												"CPX","CPY","DEC","DEX","DEY","EOR",
												"INC","INX","INY","JMP","JSR","LDA",
												"LDX","LDY","LSR","NOP","ORA","PHA",
												"PHP","PLA","PLP","ROL","ROR","RTI",
												"RTS","SBC","SEC","SED","SEI","STA",
												"STX","STY","TAX","TAY","TSX","TXS",
												"TXA","TYA"}
	for _,value := range(befehle){
		if befehl == value{
			return true
		}
	}		
	
	return false

*/




