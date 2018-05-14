package convert

/* Autor: Björn Woltering
 * Datum: 30.04.2018
 * Zweck: Konvertiert Zahlen in die unterschiedlichen Zahlensysteme
 */

import "strconv"

func hex2int16_array(hex []string) (result []int16) {
	var dez int16
	for i, hex := range hex {
		dez, _ := strconv.ParseInt(hex, 16, 16)
		result = append(result, int16(dez))
	}
	return result
}
