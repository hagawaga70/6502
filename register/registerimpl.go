package register

type impl struct {
	register byte
}

func NewRegister() *impl {
	var r *impl
	r = new(impl)
	return r
}
func (r *impl) LesenByte() (registerinhalt byte, takte int) {
	registerinhalt = r.register
	takte = 0
	return registerinhalt, takte
}

func (r *impl) SchreibenByte(daten byte) (takte int) {
	r.register = daten
	return takte
}

func (r *impl) SetzeBit(pos uint) (takte int) {
	r.register |= (1 << pos)
	takte = 0
	return 0
}
func (r *impl) SetzeBitZurueck(pos uint) (takte int) {
	//*(zahlByte) |= (0 << pos)
	mask := ^(1 << pos)
	r.register &= byte(mask)
	takte = 0
	return 0
}

func (r *impl) LeseBit(pos uint) (bitstatus bool, takte int) {
	val := r.register & (1 << pos)
	bitstatus = val > 0
	takte = 0
	return bitstatus, takte
}
