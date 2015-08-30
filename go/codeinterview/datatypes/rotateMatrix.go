package datatypes

import (
	"fmt"
	"reflect"
)

type Pixel [4]byte

func NewPixel(b byte) Pixel {
	p := Pixel{}
	p[0] = b
	p[1] = b
	p[2] = b
	p[3] = b
	return p
}
func (p Pixel) String() string {
	return string(p[0])
}

type Image struct {
	data [][]Pixel
	n    int
}

func (i *Image) Equals(compare *Image) bool {
	return reflect.DeepEqual(i, compare)
}

func (i *Image) Print() {
	fmt.Println("Image:")
	for row := 0; row < i.n; row++ {

		for col := 0; col < i.n; col++ {
			fmt.Printf("%s\t", i.data[row][col])
		}
		fmt.Println("")
	}

}

// 1 2 3
// 4 5 6
// 7 8 9
// --->
// 3 6 9
// 2 5 8
// 1 4 7
//
// [0, 0] --> [n-1, 0]
// [0, 1] --> [n-2, 0]
// [0, 2] --> [n-3, 0]
// [1, 0] --> [n-1, 1]
// [1, 1] --> [n-2, 1]
// [1, 2] --> [n-3, 1]
// [2, 0] --> [n-1, 2]
// [2, 1] --> [n-2, 2]
// [2, 2] --> [n-3, 2]
// [row, col] --> [n - 1 - col, row]
// n x n swaps
func RotateMatrix(image *Image) {

	var save Pixel
	n := image.n
	d := image.data
	for row := 0; row < n/2; row++ {
		for col := row; col < n-1-row; col++ {
			d[n-1-col][row], save = d[row][col], d[n-1-col][row]
			// d[n-1-col, row] --> d[ n-1-row, n-1-col]
			d[n-1-row][n-1-col], save = save, d[n-1-row][n-1-col]
			// d[n-1-row, n-1-col] --> d[ n-1-(n-1-col), n-1-row ] = d[ col, n-1-row ]
			d[col][n-1-row], save = save, d[col][n-1-row]
			// d[ col, n-1-row ] --> d[ n-1-(n-1-row), col ] = d[ row, col ]
			d[row][col] = save
		}
	}
}
