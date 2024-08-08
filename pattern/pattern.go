package pattern

import "fmt"

// func PatternTriangle1() {
// 	for i := 1; i <= 5; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}

// }
func ReversePatternTriangle1() {
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func TrianglePattern3(n int) {
	n = 5
	for i := 1; i <= n; i++ {
		// Print leading spaces
		for k := n - i; k > 0; k-- {
			fmt.Print(" ")
		}
		// Print asterisks
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func TrianglePattern4() {
	n := 5
	for i := n; i >= 1; i-- {
		for k := n - i; k > 0; k-- {
			fmt.Print(" ")
		}
		for j := i; j > 0; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
