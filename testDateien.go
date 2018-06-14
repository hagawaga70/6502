package main



import "./dateien"
import "fmt"
import "regexp"
import "strings"

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


}
