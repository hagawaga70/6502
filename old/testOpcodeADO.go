package main


import "./dateien"
import ."fmt"
//import "regexp"
//import "strings"
import "sort"
//import ."./assembler"
import "./opcode"




func main() {
	
	// Öffnen der Programmdatei
	var buffer  []byte
	var counter int

	dateiInhalt := dateien.Oeffnen("./programm.hag",'l')
	for !dateiInhalt.Ende(){
		buffer = append(buffer,dateiInhalt.Lesen())
		counter++
	}
	dateiInhalt.Schliessen()

	// Umwandeln der Bytes in EINE Zeichenkette
	hagCode 	:= 	string(buffer[:counter])
	opcodeList,assenblerCodeListe,pseudoCodeListe	:=	opcode.GetOpcodeList(hagCode) 

	// To store the keys in slice in sorted order
    var keys []int
    for k := range opcodeList {
        keys = append(keys, k)
    }
    sort.Ints(keys)

    // To perform the opertion you want
    for _, k := range keys {
        Println("Key:", k)
        Println("Value:", opcodeList[k])
    }
	Println(assenblerCodeListe)
	Println(pseudoCodeListe)
}
