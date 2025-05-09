/*

Write a program to accept a numbers and match the number to a switchcase.
If the number is 1 it should print "You entered one"
If the number is 2 it should print "You entered two"
If the number is 3 it should print "You entered three"
If the number is 4 it should print "You entered four"
else it should print "You entered other.."

*/

package main

import "fmt"

func main() {
	fmt.Print("enter a number-->")
	var num int
	fmt.Scanf("%d", &num)

	switch num {
	case 1:
		fmt.Println("You entered one")
	case 2:
		fmt.Println("You entered two")
	case 3:
		fmt.Println("You entered three")
	case 4:
		fmt.Println("You entered four")
	default:
		fmt.Println("You entered others ..")

	}

}
