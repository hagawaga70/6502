package register

//----NEW ...

type Register interface {
	// Vor.: ---
	// Erg.: Der Inhalt des Registers ist zurückgegeben. (Die Anzahl der benötigten Takte ist zurückgegeben - nicht umgesetzt).
	LesenByte() (registerinhalt byte, takte int)

	// Vor.: ---
	// Erg.: Das Register ist neu beschrieben. (Die Anzahl der benötigten Takte ist zurückgegeben - nicht umgesetzt)
	SchreibenByte(daten byte) (takte int)

	// Vor.: ---
	// Erg.: Das Statusbit an der übergebenen Position ist gesetzt. (Die Anzahl der benötigten Takte ist zurückgegeben - nicht umgesetzt).
	SetzeBit(pos uint) 	(takte int)

	// Vor.: ---
	// Erg.: Das Statusbit an der übergebenen Position ist zurückgesetzt. (Die Anzahl der benötigten Takte ist zurückgegeben - nicht umgesetzt).
	SetzeBitZurueck		(pos uint) 	(takte int)

	// Vor.: ---
	// Erg.: Das Statusbit an der übergebenen Position ist gelesen. (Die Anzahl der benötigten Takte ist zurückgegeben - nicht umgesetzt).
	LeseBit(pos uint) (bitstatus bool,takte int)
}
