package main

import "fmt"

func main() {
	/* Integer: int --> Based on OS architecture (x86 or x64)
	* int8: 8 bits = -128 to 127
	* int16 = -2^15 to (2^16 - 1)
	* int32 = -2^31 to (2^32 - 1)
	* int64 = -2^63 to (2^64 - 1)
	 */

	/* Positive integer: uint --> Same as regular (int) but only positive
	* uint8: 8 bits = 0 to (2^8 - 1)
	* uint16 = 0 to (2^16 - 1)
	* uint32 = 0 to (2^32 - 1)
	* uint64 = 0 to (2^64 - 1)
	 */

	/* Decimal number:
	* float32: 32 bits = (+/-)1.18e^-38 to (+/-)3.4e^38
	* float64: 64 bits = (+/-)2.23e^-308 to (+/-)1.8e^308
	 */

	/* Strings and booleans
	* string = "" --> Always double quotes
	* bool = true or false
	 */

	/*
		Complex numbers:
		* complex64: Real and imaginary float32
		* complex128: Real and imaginary float64
		* Example: c := 10 + 8i
	*/

	fmt.Println("Detail in code comments")
}
