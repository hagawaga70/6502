package opcode

import . "../speicher"
import . "../register"
import ."fmt"
import "regexp"
import "strings"
import ."../assembler"
import "encoding/hex"
import "encoding/binary"

func ExecuteOpcode ( 	opcode 				[]byte,
						speicher64k 		Speicher, 
						x_register 			Register,
						y_register 			Register,
						programmZaehlerHigh Register,
						programmZaehlerLow	Register,
						stapelzeiger		Register, 
						akku				Register, 
						statusbits			Register)( programmEnde bool){


/*

codeArray[0] == "ADC" 
*/
var opcodeHeadHEX = hex.EncodeToString([]byte{opcode[0]}) 
var dataAdressUINT16 uint16
var dataByte []byte
//var carryBit bool 
	switch opcodeHeadHEX{

		case "6d":	//ADC absolut
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{opcode[1],opcode[2]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = akku.SchreibenByte(dataByte[0])
					programmEnde = false 
 					break 

		case "65":	//ADC Zero-Page
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{byte(0),opcode[1]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = akku.SchreibenByte(dataByte[0]) 
					programmEnde = false 
 					break 

		case "69":	//ADC unmittelbar 
					_ = akku.SchreibenByte(opcode[1]) 
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------

		case "18":	//CLC
					_ = statusbits.SetzeBitZurueck(0)
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------
		case "d8":	//CLD
					_ = statusbits.SetzeBitZurueck(3)
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------
		case "f2":	//·END
					programmEnde = true 
 					break 
		//---------------------------------------------------------------------------------------------------
		case "ad":	//LDA absolut
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{opcode[1],opcode[2]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = akku.SchreibenByte(dataByte[0]) 
						_ = statusbits.SetzeBit(0)
					if (int(dataByte[0]) >> 7) == 1{
						_ = statusbits.SetzeBit(7)
					}else{
						_ = statusbits.SetzeBitZurueck(7)
					}

					if int(dataByte[0]) == 0{
						_ = statusbits.SetzeBit(1)
					}else{
						_ = statusbits.SetzeBitZurueck(1)
					}

					programmEnde = false 
 					break 

		case "a5":	//LDA Zero-Page
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{byte(0),opcode[1]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = akku.SchreibenByte(dataByte[0]) 
					programmEnde = false 
					if (int(dataByte[0]) >> 7) == 1{
						_ = statusbits.SetzeBit(7)
					}else{
						_ = statusbits.SetzeBitZurueck(7)
					}

					if int(dataByte[0]) == 0{
						_ = statusbits.SetzeBit(1)
					}else{
						_ = statusbits.SetzeBitZurueck(1)
					}
 					break 

		case "a9":	//LDA unmittelbar 
					_ = akku.SchreibenByte(opcode[1]) 
					buffer:=opcode[1]
					if int(buffer >> 7) == 1{
						_ = statusbits.SetzeBit(7)
					}else{
						_ = statusbits.SetzeBitZurueck(7)
					}

					buffer=opcode[1]
					if int(buffer) == 0{
						_ = statusbits.SetzeBit(1)
					}else{
						_ = statusbits.SetzeBitZurueck(1)
					}
					programmEnde = false 

 					break 
		//---------------------------------------------------------------------------------------------------
		case "ae":	//LDX absolut
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{opcode[1],opcode[2]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = x_register.SchreibenByte(dataByte[0]) 
					programmEnde = false 
 					break 

		case "a6":	//LDX Zero-Page
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{byte(0),opcode[1]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = x_register.SchreibenByte(dataByte[0]) 
					programmEnde = false 
 					break 

		case "a2":	//LDX unmittelbar
					_ = x_register.SchreibenByte(opcode[1]) 
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------
		case "ac":	//LDY absolut
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{opcode[1],opcode[2]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = y_register.SchreibenByte(dataByte[0]) 
					programmEnde = false 
 					break 

		case "a4":	//LDY Zero-Page
					dataAdressUINT16 = binary.BigEndian.Uint16([]byte{byte(0),opcode[1]})
					dataByte,_ = speicher64k.Lesen([]uint16{dataAdressUINT16,dataAdressUINT16})
					_ = y_register.SchreibenByte(dataByte[0]) 
					programmEnde = false 
 					break 

		case "a0":	//LDY unmittelbar
					_ = y_register.SchreibenByte(opcode[1]) 
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------
		case "8d":	//STA absolut
					dataAdressUINT16  = binary.BigEndian.Uint16([]byte{opcode[1],opcode[2]})
					registerByte , _ :=	akku.LesenByte() 
					_ =	speicher64k.Schreiben([]uint16{dataAdressUINT16,dataAdressUINT16},[]byte{registerByte}) 
					programmEnde = false 
 					break 

		case "85":	//STA Zero-Page
					dataAdressUINT16  = binary.BigEndian.Uint16([]byte{byte(0),opcode[1]})
					registerByte , _ :=	akku.LesenByte() 
					_ =	speicher64k.Schreiben([]uint16{dataAdressUINT16,dataAdressUINT16},[]byte{registerByte}) 
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------
		case "8e":	//STX absolut
					dataAdressUINT16  = binary.BigEndian.Uint16([]byte{opcode[1],opcode[2]})
					registerByte , _ :=	x_register.LesenByte() 
					_ =	speicher64k.Schreiben([]uint16{dataAdressUINT16,dataAdressUINT16},[]byte{registerByte}) 
					programmEnde = false 
 					break 

		case "86":	//STX Zero-Page
					Println("STX Zero",opcode)
					dataAdressUINT16  = binary.BigEndian.Uint16([]byte{byte(0),opcode[1]})
					registerByte , _ :=	x_register.LesenByte() 
					_ =	speicher64k.Schreiben([]uint16{dataAdressUINT16,dataAdressUINT16},[]byte{registerByte}) 
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------
		case "8c":	//STY absolut
					dataAdressUINT16  = binary.BigEndian.Uint16([]byte{opcode[1],opcode[2]})
					registerByte , _ :=	y_register.LesenByte() 
					_ =	speicher64k.Schreiben([]uint16{dataAdressUINT16,dataAdressUINT16},[]byte{registerByte}) 
					programmEnde = false 
 					break 

		case "84":	//STY Zero-Page
					dataAdressUINT16  = binary.BigEndian.Uint16([]byte{byte(0),opcode[1]})
					registerByte , _ :=	y_register.LesenByte() 
					_ =	speicher64k.Schreiben([]uint16{dataAdressUINT16,dataAdressUINT16},[]byte{registerByte}) 
					programmEnde = false 
 					break 
		//---------------------------------------------------------------------------------------------------

		default:
			panic("ADO Opcode:Opcode"+opcodeHeadHEX+" nicht vorhanden")
	}
	return
} 


func GetOpcodeList(hagCode string)( map[int][]string,map[int][]string,map[string][]string){

	var assemble Assembler = NewAssembler() 
	var codeLine 		string
	var startAdresse 	string
	var opcode 			[]string
	var err				bool	
	var naechsteAdresse string
	var debug bool = false
	//var buffer 				[]byte
	var counterOpcodeList 	int						// Zeilennummerierung für die opcodeList
	//var counter 			int

	opcodeList 				:= map[int][]string{} //Deklarierung der HashListe
	assemblerCodeList 		:= map[int][]string{} //Deklarierung der HashListe

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
	"TXA":{},"TYA":{},"END":{}}


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
	pseudoBefehle			:= regexp.MustCompile(`^\s*[A-Za-z0-9]+\s*=\s*(#\$|\$)[A-Fa-f0-9]+\s*$`)
	pseudoBefehleExakt		:= regexp.MustCompile(`^\s*([A-Z0-9]+)\s*=\s*((#\$|\$)[0-9A-Fa-f]{1,4})\s*$`)
	startAdresseRegex		:= regexp.MustCompile(`^\s*\*\s*=.*$`)
	startAdresseRegexExakt	:= regexp.MustCompile(`^\s*\*\s*=\s*\$([0-9A-Fa-f]{1,4})\s*$`)
	pseudoBefehleHASH		:= map[string][]string{}	// Der Name des Pseudobefehls ist der Key des Hashes



	for _,codeLine = range(hagCodeArray){
		//Println(codeLine)

		// Überspringe leere Zeilen
		if ueberspringeLeereZeilen.MatchString(codeLine){
			continue
		}

		// Überspringe reine Semikolonzeilen 
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

		// Lesen der Startadresse
		if  startAdresseRegex.MatchString(codeLine){

			if	startAdresseRegexExakt.MatchString(codeLine){
				startAdresse  = startAdresseRegexExakt.ReplaceAllString(codeLine, `$1`)
				pseudoBefehleHASH["$t@rt@dre$$e"] = []string{startAdresse}
				continue
			}else{
				panic("002: Die Angabe der Startadresse ist nicht korrekt -> *=$xxxx - x = [A-Fa-f0-9] ") 
			}
		}
		// Lesen des Pseudobefehls
		if  pseudoBefehle.MatchString(codeLine){
			if	pseudoBefehleExakt.MatchString(codeLine){

				pseudoBefehlName	:= pseudoBefehleExakt.ReplaceAllString(codeLine, `$1`)
				pseudoBefehlInhalt	:= pseudoBefehleExakt.ReplaceAllString(codeLine, `$2`)
				pseudoBefehleHASH[pseudoBefehlName] = []string{pseudoBefehlInhalt}
				continue
			}else{
				panic("001:Die Zuweisung des Pseudocodes entspricht nicht den Anforderungen: z.B ADR1=$1111")
			}
		}
		if debug{
			Println("---------------------------------------------------------------------")
			Println("testDateien pseudoBefehleHASH codeLine")
			Println(pseudoBefehleHASH)
			Println(codeLine)
			Println("°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")
		}
		codeArray:=strings.Fields(codeLine)
		assemblerCodeList[counterOpcodeList] = append(assemblerCodeList[counterOpcodeList],codeArray...)

		// Eine Sprungmarke wird abgespeichert
		if _, ok := befehleListe[codeArray[0]]; !ok {

			pseudoBefehleHASH[codeArray[0]] = []string{"$"+startAdresse}
			codeArray = codeArray[1:]
		}
		if			codeArray[0] == "ADC" ||
					codeArray[0] == "LDA" ||
					codeArray[0] == "LDX" ||
					codeArray[0] == "LDY" ||
					codeArray[0] == "STA" ||
					codeArray[0] == "STX" ||
					codeArray[0] == "STY" {

			 opcode,err,naechsteAdresse = assemble.TranslateXXX(codeArray,pseudoBefehleHASH,startAdresse)

		}else if 	codeArray[0] == "CLC" ||
					codeArray[0] == "CLD"	{

			 opcode,err,naechsteAdresse = assemble.TranslateModifyFlags(codeArray,startAdresse)

		}else if  	codeArray[0] == "END" {
			 opcode,err,naechsteAdresse = assemble.TranslateEnd(codeArray,startAdresse)
		}else{
			panic("Der Befehl ist nicht im Befehlssatz vorhanden!!!")
		}

		opcodeList[counterOpcodeList] = []string{startAdresse}
		opcodeList[counterOpcodeList] = append(opcodeList[counterOpcodeList],opcode...)

		startAdresse = naechsteAdresse
		counterOpcodeList++

		if debug{
			Println("---------------------------------------------------------------------")
			Println("testDateien -opcode -err")
			Println(opcode)
			Println(err)
			Println("°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°")
		}
    }
	return opcodeList,assemblerCodeList,pseudoBefehleHASH
}




