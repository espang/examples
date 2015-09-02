package bit_manipulation

import "testing"

func bitsToInt(s string) int32 {
	n := uint(len(s) - 1)
	var result int32
	for _, r := range s {
		if r == '1' {
			result += 1 << n
		}
		n--
	}
	return result
}

var testmap = map[string]int32{
	"0":      0,
	"1":      1,
	"10":     2,
	"11":     3,
	"100":    4,
	"101":    5,
	"110":    6,
	"111":    7,
	"101011": 43, //1+2+8+32
}

func TestBitsToInt(t *testing.T) {
	for s, i := range testmap {
		if bitsToInt(s) != i {
			t.Errorf("bitsToInt(%s) -> %b; expected %b", s, bitsToInt(s), i)
		}
	}
}

func TestInsertion(t *testing.T) {
	n := bitsToInt("10000000000")
	m := bitsToInt("10011")

	result := Insertion(n, m, 2, 6)

	if exp := bitsToInt("10001001100"); result != exp {
		t.Errorf("Expected %d got %d", exp, result)
	}
}
