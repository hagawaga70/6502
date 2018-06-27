package gfxElemente

//----NEW ...

type GfxElement interface {
	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Eine Speicherseite steht der übergebenen Position im  Grafikfenster.
	//		 Diese beinhaltet die Daten, Seiten- und Adressangaben. Die durch dem Assemblerbefehl 
	//		 geänderten Seitenzeilen sind rot markiert
	AbbildSpeicherseite1(x1, y1 uint16, groesse uint, seitenInhalt []byte, seitenInhaltAlt []byte)

	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Der Registerinhalt und die Bezeichnung des Registers steht an der übergebenen Position
	//       im Grafikfenster.Die durch dem Assemblerbefehl geänderten Registerwerte sind rot markiert
	AbbildRegister(x1, y1 uint16, name string, registerInhalt byte, registerInhaltAlt byte)

	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Der Wert des Statusbit und die Bezeichnung des Statusbits  steht an der übergebenen Position
	//       im Grafikfenster.Der Wert des durch dem Assemblerbefehl geänderten Statusbits ist rot markiert
	AbbildFlag(x1, y1 uint16, label string, flagStatus int, flagStatusAlt int)


	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Ein Label steht an der übergebenen Position und in der angegebenen Schriftgröße im Grafikfenster
	AbbildLabel(x1,y1 uint16, label string, schriftGroesse int, r uint8, g uint8, b uint8)


	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Das unter dem angegebenen Pfad befindliche Bild steht unter der angegebenen Position im Grafikfenster 
	AbbildBild(x1,y1 uint16, pfad string)
}
