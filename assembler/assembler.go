package assembler

//----NEW ...

type Assembler interface {
	// Vor.: -
	// Erg.: Die hexadezimale Übersetzung des (LDA,LDX,LDY,STA,STX,STY)-Befehls ist zurückgegeben. Ein
	//		 oder zwei hexadezimale Adressbytes sind zurückgegeben. Die hexadezimale
	//		 Anfangsadresse des nächsten Speicherplatz für den nächsten Befehl ist zurück
	//		 gegeben. Eine Fehlermeldung "err" ist zurückgegeben. Im Fehlerfall enthält das OPcodeSlice eine
	//		 Aussagekräftige Fehlermeldung
	TranslateXXX(assemblerCode []string, pseudoBefehle map[string][]string, aktuelleAdresse string)(optcode []string, err bool,naechsteAdresse string)
	// Vor.: -
	// Erg.:  Die hexadezimale Übersetzung des (CLC,CLD)-Befehls ist zurückgegeben 
	TranslateModifyFlags(assemblerCode []string, aktuelleAdresse string)(optcode []string, err bool,naechsteAdresse string)
}
