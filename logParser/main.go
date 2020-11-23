package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	dict := make(map[string]int)

	for in.Scan() {
		fields := strings.Fields(in.Text())

		domain := fields[0]
		visit, _ := strconv.Atoi(fields[1])
		dict[domain] += visit
	}

	fmt.Printf("%-30s - %10s \n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for domain, visit := range dict {
		fmt.Printf("%-30s - %10d\n", domain, visit)

	}
}
