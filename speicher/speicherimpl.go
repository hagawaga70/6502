package speicher

//import "strconv"

type impl struct {
	speicher [65536]byte // Der 65k B-Speicher wird durch ein Array simuliert
}

func NewSpeicher() *impl {
	var s *impl
	s = new(impl)
	return s
}
func (s *impl) Lesen(speicherbereich []uint16) (daten []byte, takte int) {
	if len(speicherbereich) == 1 { // - Übergabe einer einzelnen Adresse
		daten = append(daten, s.speicher[speicherbereich[0]]) // - Der Inhalt der Adresse wird i, Array Daten gespeichert
		return daten, 0                                       // - Der zweite Rückgabewert ist die Anzahl der benötigten Takte. Derzeit ohne Funktion
	} else if len(speicherbereich) == 2 { // - Übergabe von zwei Adressen, zur Rückgabe eines Speicherbereichs
		for i := speicherbereich[0]; i <= speicherbereich[1]; i++ { // -
			//s.speicher[i], daten = daten[0], daten[1:]
			daten = append(daten, s.speicher[i])
		}
		return daten, 0 // - Rückgabe der Daten der Adressen des übergebenen Speicherbereichs
	} else {
		panic("package speicher - Error 001:Es wurden mehr als 2 Speicheradressen übergeben ")
	}
}

func (s *impl) Schreiben(speicherbereich []uint16, daten []byte) (takte int) {
	if len(speicherbereich) == 1 { // - Der Speicher einer einzelnen Adresse soll neu beschreiben werden
		s.speicher[speicherbereich[0]] = daten[0] // - Beschreiben der Speicheradresse mit neuen Daten
		takte = 0
		return takte // - Rückgabe der benötigten Taktzyklen (noch nicht umgesetzt)
	} else if len(speicherbereich) == 2 { // - Übergabe der Adressen eines Speicherbereichs
		for i := speicherbereich[0]; i <= speicherbereich[1]; i++ { // - Beschreiben der Zellen eines Speicherbereichs
			s.speicher[i], daten = daten[0], daten[1:]
		}
		takte = 0
		return takte // - Rückgabe der benötigten Taktzyklen (noch nicht umgesetzt)
	} else {
		panic("package speicher - Error 002:Es wurden mehr als 2 Speicheradressen übergeben ")
	}
}
