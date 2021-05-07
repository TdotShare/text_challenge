package examples

import "fmt"

func GetWall(number int) {

	for i := 0; i < number; i++ {

		for k := 0; k < number; k++ {

			if i != 0 && i != number-1 {

				if k == 0 || k == number-1 {
					fmt.Print("*")
				} else {
					fmt.Print(" ")
				}

			} else {
				fmt.Print("*")
			}

		}

		fmt.Println("")

	}

}
