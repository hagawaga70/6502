package main


import ."./assembler"
import  "fmt"



func  main(){
	var ta Assembler = NewAssembler() 
	var code []string = []string{"LDA","$FF"}
	//code[0] = "LDA"
	//code[1] = "$FF"
	opcode,takte := ta.TranslateLDA(code)
	fmt.Println(opcode)
	fmt.Println(takte)
}
