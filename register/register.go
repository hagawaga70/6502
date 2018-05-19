package register

//----NEW ...

type Register interface {
	// Vor.: -
	// Eff.: Der Inhalt des Registers ist zurückgegeben.Die Anzahl der benötigten Takte ist zurückgegeben.
	LesenByte() (registerinhalt byte, takte int)

	// Vor.: -
	// Erg.: Das Register ist neu beschrieben.Die Anzahl der benötigten Takte ist zurückgegeben.
	SchreibenByte(daten byte) (takte int)

	// Vor.: -
	// Erg.: Das Register ist neu beschrieben.Die Anzahl der benötigten Takte ist zurückgegeben.
	SetzeBit(pos uint) 	(takte int)
	SetzeBitZurueck		(pos uint) 	(takte int)
	LeseBit(pos uint) (bitstatus bool,takte int)
}
