package main

import (
    "fmt"
    "strconv"
)

func main() {
    zahl1, err := strconv.ParseInt("FF", 16, 0)
    zahl2, err := strconv.ParseInt("2", 16, 0)
    if err != nil {
        panic(err)
    }
	zahl3 := zahl1 + zahl2
	fmt.Println(strconv.FormatInt(zahl3, 16))
}
