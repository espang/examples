package bit_manipulation

import "fmt"

func Insertion(n, m int32, posI, posJ uint8) int32 {

	// insert in bits: xxxxxxxxxx yyyy from i to j:
	//    could be     xxxxyyyyxx for i = 2 and j = 5
	//    or           xxx00yyyyx for i = 1 and j = 6

	// elements left from insertion
	var left int32 = n >> posJ
	// elements right from inversion
	var right int32 = n & (2 << posI)
	// elements in the middle with leading zeros
	var middle int32 = m & (2<<(posJ-posI) - 1)
	fmt.Printf("%b %d\n", ^0, ^0)
	// combine left middle right, shifting left and middle to their position
	return left<<posJ | middle<<posI | right
}
