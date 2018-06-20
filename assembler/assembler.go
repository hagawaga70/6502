package assembler

//----NEW ...

type Assembler interface {
	// Vor.: -
	// Erg.: Die hexadezimale Übersetzung des LDA-Befehls ist zurückgegeben. Eine  zwei
	//		 oder vierstellige hexadezimale Adressen ist zurückgegeben. Die hexadezimale
	//		 Anfangsadresse des nächsten Speicherplatz für den nächsten Befehl ist zurück
	//		 gegeben. 
	TranslateLDA(assemblerCode []string, pseudoBefehle map[string][]string, aktuelleAdresse string) (optcode []string, takte string ,naechsteAdresse string)

	// Vor.: -
	// Erg.: Die hexadezimale Übersetzung des LDX-Befehls ist zurückgegeben. Eine  zwei
	//		 oder vierstellige hexadezimale Adressen ist zurückgegeben. Die hexadezimale
	//		 Anfangsadresse des nächsten Speicherplatz für den nächsten Befehl ist zurück
	//		 gegeben. 
	TranslateLDX(assemblerCode []string, pseudoBefehle map[string][]string, aktuelleAdresse string) (optcode []string, takte string ,naechsteAdresse string)

}
