package main


import ."./assembler"
import  "fmt"



func  main(){
/*	
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

*/

	var ta Assembler = NewAssembler() 
	var code []string = []string{"LDA","#$FF"}
	//code[0] = "LDA"
	//code[1] = "$FF"
	opcode,takte := ta.TranslateLDA(code)
	for _,value := range(opcode){
		fmt.Println(value)
	}
	fmt.Println(takte)
}
