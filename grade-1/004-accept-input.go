/*

accept the radius as the input and calculate the perimeter of a circle

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("calculate the perimeter of a circle")
	fmt.Println("eneter the radius::")
	var radius float64
	fmt.Scanf("%f", &radius)

	fmt.Println("Perimeter is %f", 2*3.14*radius)
}
