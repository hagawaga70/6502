package gfxElemente

//----NEW ...

type GfxElement interface {
	// Vor.: Das gfx-Fenster ist geöffnet.
	// Eff.: Eine Speicherseite ist ins gfx-Fenster geschrieben.
	AbbildSpeicherseite(x1, y1 uint16, groesse uint, seiteninhalt []byte)
}
