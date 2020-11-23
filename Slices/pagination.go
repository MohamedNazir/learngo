package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func pagination() {

	items := []string{"apple", "banana", "cherry", "Durian",
		"eat", "fig", "Guava", "Horlicks :)",
		"ink", "jamun", "kite", "lemon",
		"orange", "plums"}
	//fmt.Println(names)

	//	s.Show("0:4", items[0:4])
	//	s.Show("4:8", items[4:8])
	//	s.Show("8:12", items[8:12])
	// 	s.Show("12:16", items[12:16]) // this will lead to an runtime error sinve the slice length is only 14

	const pagesize = 4
	l := len(items)

	for from := 0; from < l; from += pagesize {
		to := from + pagesize

		if to > l {
			to = l
		}
		currentPage := items[from:to]
		head := fmt.Sprintf("Page #%d", (from/pagesize)+1)
		s.Show(head, currentPage)
	}

}
