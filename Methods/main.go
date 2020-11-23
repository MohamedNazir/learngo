package main

func main() {

	var (
		mobidick = book{"Mobidick", 10.03}
		soccer   = game{"Football", 20.24}
		cricket  = game{"IPL", 20.24}
	)

	printBook(mobidick)
	soccer.print() // by defualt they type will be passed
	cricket.print()

	game.print(cricket) // this is called method expressions. this will also work.
}
