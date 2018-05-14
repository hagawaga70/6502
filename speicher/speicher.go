package speicher

//----NEW ...

type Speicher interface {
	// Vor.: -
	// Eff.: Das oder die Bytes sind zurückgegeben.Die Anzahl der benötigten Takte ist zurückgegeben.
	Lesen(speicherbereich []int16) (daten []byte, takte int)

	// Vor.: -
	// Erg.: Das/die Byte/s ist/sind an der/den angegebenen Adresse/n abgespeichert.Die Anzahl der benötigten Takte ist zurückgegeben.
	Schreiben(speicherbereich []int16, daten []byte) (takte int)
}
