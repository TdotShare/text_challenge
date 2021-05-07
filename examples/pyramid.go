package examples

import "fmt"

func GetPyramid_no_one(number int) {

	for i := 0; i < number; i++ {

		for k := 0; k <= i; k++ {
			fmt.Print("*")
		}

		fmt.Println("")

	}
}
