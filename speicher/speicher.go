package speicher

//----NEW ...

type Speicher interface {
	// Vor.: -
	// Eff.: Das oder die Bytes ist/sind zurückgegeben. (Die Anzahl der benötigten Takte ist zurückgegeben - nicht realisert)
	Lesen(speicherbereich []uint16) (daten []byte, takte int)

	// Vor.: -
	// Erg.: Das/die Byte/s ist/sind an der/den angegebenen Adresse/n abgespeichert.(Die Anzahl der benötigten Takte ist zurückgegeben - nicht realisiert)
	Schreiben(speicherbereich []uint16, daten []byte) (takte int)
}
