package main

import "fmt"
import "math"

// Learning Go!

const str string = "constant"

func main() {
	// variables
	fmt.Println("////// VARIABLES //////")

	var s string = "initial"
	fmt.Println(s)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

	// constants
	fmt.Println("////// CONSTANTS //////")

	fmt.Println(str)

	const n = 500000000

	const d2 = 3e20 / n
	fmt.Println(d2)

	fmt.Println(int64(d2))

	fmt.Println(math.Sin(n))

	fmt.Println("////// LOOPS //////")

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
}