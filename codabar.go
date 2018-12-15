package main

import (
	"fmt"
)

const mod = 16

var encodingTable = map[rune]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'-': 10,
	'$': 11,
	':': 12,
	'/': 13,
	'.': 14,
	'+': 15,
	'A': 16,
	'B': 17,
	'C': 18,
	'D': 19,
}

func main() {
	originalCode := "012345678901"
	fmt.Println("originalCode", originalCode)
	codeWithCheckSum := addCheckSum(originalCode)
	fmt.Println("codeWithCheckSum", codeWithCheckSum)
}

func addCheckSum(originalCode string) (codeWithCheckSum string) {
	// Codabar required check digit according to Modulo 16
	// See also: https://www.activebarcode.com/codes/checkdigit/modulo16.html
	var ref int
	for _, v := range originalCode {
		ref = ref + encodingTable[v]
	}
	modulo := ref % mod
	diff := mod - modulo

	var checkSum rune
	for k, v := range encodingTable {
		if v == diff {
			checkSum = k
		}
	}

	codeWithCheckSum = originalCode + string(checkSum)
	return
}
