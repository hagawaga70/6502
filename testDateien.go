package main



import "./dateien"
import "fmt"
import "regexp"
import "strings"
import "reflect"


var buffer 		[]byte
var counter 	int
func main(){
		dateiInhalt := dateien.Oeffnen("./programm.hag",'l')
	for !dateiInhalt.Ende(){
		buffer = append(buffer,dateiInhalt.Lesen())
		counter++
	}
	dateiInhalt.Schliessen()
 	hagCode:= string(buffer[:counter])
	fmt.Println(hagCode)

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
	fmt.Println(reflect.TypeOf(m))
    // Add a string at the dog key.
    // ... Append returns the new string slice.
    res := append(m["dog"], "brown")
    fmt.Println(res)

    // Add a key for fish.
    m["fish"] = []string{"orange", "red"}

    // Print slice at key.
    fmt.Println(m["fish"])

    // Loop over string slice at key.
    for i := range m["fish"] {
        fmt.Println(i, m["fish"][i])
    }



codeLine := "   		ADC #$11 ;Hier steht ein Text   "
//re := regexp.MustCompile(`(^|[^_])\bproducts\b([^_]|$)`)
re := regexp.MustCompile(`^\s*(.*)\s*$`)			// Führende und abschliessende Leerzeichen, Tabs etc. werden gelöscht

codeLine = re.ReplaceAllString(codeLine, `$1`)
checkKommentar := regexp.MustCompile(`^.*;.*$`)	


re = regexp.MustCompile(`(.*);.*$`)			// Führende und abschliessende Leerzeichen, Tabs etc. werden gelöscht

if checkKommentar.MatchString(codeLine){				// Kommentartest 
	codeLine = re.ReplaceAllString(codeLine, `$1`)		// Löschen des Kommentars
	fmt.Println(codeLine)
	codeArray := strings.Fields(codeLine)
	fmt.Println(codeArray)

}else{
	fmt.Println(codeLine)
}

var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)

	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))

fmt.Println(testBefehl("KKK"))
}

func testBefehl(befehl string) bool{
	var befehle []string  =  	[]string{		"ADC","AND","ASL","BCC","BCS","BEQ",
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



}



