package gfxElemente

//----NEW ...

type GfxElement interface {
	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Eine Speicherseite ist ins gfx-Fenster geschrieben.
	AbbildSpeicherseite(x1, y1 uint16, groesse uint, seiteninhalt []byte)
	//--------------------------------------------------------------------
	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Eine Speicherseite ist ins gfx-Fenster geschrieben.
	AbbildSpeicherseite1(x1, y1 uint16, groesse uint, seiteninhalt []byte)
	//--------------------------------------------------------------------
	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Der Registerinhalt und die Bezeichnung des Registers ist ins gfx-Fenster geschrieben.
	AbbildRegister(x1, y1 uint16, name string, registerInhalt byte, registerInhaltAlt byte)
	//--------------------------------------------------------------------
}