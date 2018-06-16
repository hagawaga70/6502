package assembler

//----NEW ...

type Assembler interface {
	// Vor.: -
	// Erg.: Der Inhalt des Registers ist zurückgegeben.Die Anzahl der benötigten Takte ist zurückgegeben.
	TranslateLDA(assemblerCode []string) (optcode []string, takte int)
}
