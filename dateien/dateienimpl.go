package dateien

import ( "os" ; "io" )

type impl struct {
	f *os.File
}
 	
// Vor.: dateiname ist ein gültiger Dateiname. modus ist 'l', 's' oder 'a'. 
// Erg.: Eine Variable vom Typ Datei ist initialisiert und geliefert.
//       War modus 'l', so ist die zugehörige Datei zum Lesen geöffnet.
//       Das allererste Byte des Stroms ist aktuell.
//       War modus 's', so ist die zugehörige Datei leer und zum
//       Schreiben geöffnet.
//       War modus 'a', so ist die zugehörige Datei zum Schreiben geöffnet,
//       jedoch ist der alte Inhalt erhalten geblieben und es wird
//       beim Schreiben hinten angefügt.
func Oeffnen (dateiname string, modus byte) *impl {
	var d = new(impl)
	var err error
	
	switch modus {
		case 'l':
		((*d).f), err= os.Open(dateiname) 
		case 's':
		((*d).f), err= os.Create(dateiname)
		case 'a':
		((*d).f), err= os.OpenFile(dateiname,os.O_WRONLY | os.O_APPEND,0)
		default:
		panic("Falscher Modus-Parameter! Programmabbruch!")
	} 
	if err != nil {
		panic ("Fehler beim Öffnen der Datei! Programmabbruch!")
	}
	return d
}
	
	
//Vor.: Die Datei ist zum Schreiben geöffnet.
//Erg.: b ist an das Ende des Datenstroms angefügt. Wenn dies nicht
//      möglich war, ist das Programm abgebrochen.
func (d *impl)	Schreiben (b byte) {
	var b1 = make([]byte,1)
	var n int
	b1[0] = b
	n, _ = (*d).f.Write (b1)
	if n != 1 {
		panic("Konnte nicht in die Datei schreiben! Programmabbruch!")
	}
}
	
// Vor.: Die Datei ist zum Lesen geöffnet.
// Erg.: True ist geliefert, gdw. es kein aktuelles Byte gibt, d.h. 
//       das Ende des Stroms ist erreicht. Misslingt der Test, so ist das
//       Programm abgebrochen.
func (d *impl)	Ende () bool {
	var b1 = make([]byte,1)
	var n int
	var err error
	n, err =(*d).f.Read (b1)
	if n == 1 {
		(*d).f.Seek(-1,1)
		return false
	}
	if n == 0 && err == io.EOF {
		return true
	}
	panic("Fehler beim Test auf Ende! Programmabbruch!")
}
	
// Vor.: Es gibt ein aktuelles Byte.
// Erg.: Das aktuelle Byte ist geliefert.
// Eff.: Das darauf folgende Byte ist nun aktuell, wenn es eins gibt.
//       Andernfalls ist das Ende des Stroms erreicht. Bei einem Fehler
//       ist das Programm abgebrochen.
func (d *impl) Lesen () byte {
	var b1 = make([]byte,1)
	var n int
	n, _ = (*d).f.Read (b1)
	if n == 1 {
		return b1[0]
	} else {
		panic("Fehler beim Lesen! Programmabbruch!")
	}
}
	
// Vor.: Die Datei wurde noch nicht geschlossen.
// Eff.: Die Datei wurde geschlossen und steht zum Lesen und Schreiben
//       nicht mehr zur Verfügung. Gab es beim Schliessen der Datei ein
//       Problem, so ist das Programm abgebrochen!
func (d *impl) Schliessen () {
	var err error
	err = (*d).f.Close ()
	if err != nil {
		panic ("Fehler beim Schließen der Datei! Programmabbruch!")
	}
}

