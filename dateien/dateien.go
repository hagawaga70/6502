package dateien

/* Eine Datei kann als Strom unbekannter Länge von Bytes betrachtet werden. 
 * Zu jeder Datei gehört eine Zeichenkette, der sogenannte Dateiname, über 
 * den man den Strom (später) wieder finden kann.
 * Ein Strom kann durch Anhängen von Bytes an sein Ende modifiziert 
 * werden, das nennt man Schreiben.
 * Zu jedem Zeitpunkt ist nur ein einziges Byte (aktuelles Byte) sichtbar,
 * es sei denn, man ist am Ende der Datei.
 * Der Zugriff auf das aktuelle Byte heißt Lesen, dadurch wird das folgende 
 * Element sichtbar. Jeder Strom besitzt einen Modus: Er ist entweder
 * zum Lesen oder Schreiben geöffnet. */
 
type Datei interface {
	
	// Vor.: dateiname ist ein gültiger Dateiname. modus ist 'l', 's' oder 'a'. 
	// Erg.: Eine Variable vom Typ Datei ist initialisiert und geliefert.
	//       War modus 'l', so ist die zugehörige Datei zum Lesen geöffnet.
	//       Das allererste Byte des Stroms ist aktuell.
	//       War modus 's', so ist die zugehörige Datei leer und zum
	//       Schreiben geöffnet.
	//       War modus 'a', so ist die zugehörige Datei zum Schreiben geöffnet,
	//       jedoch ist der alte Inhalt erhalten geblieben und es wird
	//       beim Schreiben hinten angefügt.
	//       Konnte die Datei nicht geöffnet werden, ist das Programm
	//       abgebrochen.
	// Oeffnen (dateiname string, modus byte) Datei 
	// ^^^^^^^
	// entspricht dem sonst zu verwendenden New
	
	//Vor.: Die Datei ist zum Schreiben geöffnet.
	//Erg.: b ist an das Ende des Datenstroms angefügt. Wenn dies nicht
	//      möglich war, ist das Programm abgebrochen.
	Schreiben (b byte)
	
	// Vor.: Die Datei ist zum Lesen geöffnet.
	// Erg.: True ist geliefert, gdw. es kein aktuelles Byte gibt, d.h. 
	//       das Ende des Stroms ist erreicht. Misslingt der Test, so ist das
	//       Programm abgebrochen.
	Ende () bool
	
	// Vor.: Es gibt ein aktuelles Byte.
	// Erg.: Das aktuelle Byte ist geliefert.
	// Eff.: Das darauf folgende Byte ist nun aktuell, wenn es eins gibt.
	//       Andernfalls ist das Ende des Stroms erreicht. Bei einem Fehler
	//       ist das Programm abgebrochen.
	Lesen () byte
	
	// Vor.: Die Datei wurde noch nicht geschlossen.
	// Eff.: Die Datei wurde geschlossen und steht zum Lesen und Schreiben
	//       nicht mehr zur Verfügung.Gab es beim Schliessen der Datei ein
	//       Problem, so ist das Programm abgebrochen!
	Schliessen ()
}
