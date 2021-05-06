package main

import "fmt"

type tables struct {
	box []box
}

type box struct {
	value []string
}

func (b *tables) AddItem(item box) {
	b.box = append(b.box, item)
}

func main() {
	textChallenge("aaaaa", 2)
}

func textChallenge(texts string, rows int) {

	row := rows
	text := texts
	table := tables{}

	for i := 0; i < row; i++ {
		box := box{value: []string{}}
		table.AddItem(box)
	}

	for i := 0; i < len(text); i++ {
		for k := 0; k < len(table.box); k++ {
			table.box[k].value = append(table.box[k].value, "-")
		}
	}

	start := 0
	down := true

	for i := 0; i < len(text); i++ {

		table.box[start].value[i] = string(text[i])

		if down {

			start++

			if start == row {
				down = false
				start = start - 2
			}

		} else {

			start--

			if start < 0 {
				down = true
				start = start + 2
			}

		}
	}

	for _, item := range table.box {
		for _, el := range item.value {
			fmt.Print(el)
		}

		fmt.Println("")
	}

}
