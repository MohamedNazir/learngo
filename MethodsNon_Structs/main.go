package main

func main() {

	var (
		soccer  = game{"Football", 20.24}
		cricket = game{"IPL", 20.24}
	)

	var items []*game

	items = append(items, &soccer, &cricket)

	my := list(items)

	my.print()
}

type game struct {
	title string
	price float64
}
