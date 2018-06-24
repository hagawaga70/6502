package main

/* Autor: Björn Woltering
* Datum:
* Zweck:
 */
import "fmt"

func main() {
	var zahlByte byte = 00000000
	var zeigerZahlByte *byte = &zahlByte

	fmt.Println(zahlByte)
	setzeBit(zeigerZahlByte, 0)
	setzeBit(zeigerZahlByte, 1)
	setzeBit(zeigerZahlByte, 2)
	setzeBit(zeigerZahlByte, 3)
	setzeBit(zeigerZahlByte, 4)
	fmt.Println(leseBit(zeigerZahlByte, 0))
	fmt.Println(leseBit(zeigerZahlByte, 1))
	fmt.Println(leseBit(zeigerZahlByte, 2))
	fmt.Println(leseBit(zeigerZahlByte, 3))
	fmt.Println(leseBit(zeigerZahlByte, 4))
	fmt.Println(leseBit(zeigerZahlByte, 5))
	fmt.Println(leseBit(zeigerZahlByte, 6))
	fmt.Println(leseBit(zeigerZahlByte, 7))
	fmt.Println(leseBit(zeigerZahlByte, 8))

	//setzeBitZurueck(zeigerZahlByte, 7)

	fmt.Println()
	// Sets the bit at pos in the integer n.
	/*

		// Clears the bit at pos in n.
	*/
}

func setzeBit(zahlByte *byte, pos uint) {
	*(zahlByte) |= (1 << pos)
}

func setzeBitZurueck(zahlByte *byte, pos uint) {
	//*(zahlByte) |= (0 << pos)
	mask := ^(1 << pos)
	*(zahlByte) &= byte(mask)
}
func leseBit(zahlByte *byte, pos uint) bool {
	val := *(zahlByte) & (1 << pos)
	return (val > 0)
}
