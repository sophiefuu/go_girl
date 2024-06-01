package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What do you want?")
	fmt.Println("Print Menu:")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)

	type menuItem struct {
		name  string
		price map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", price: map[string]float64{"small": 1.65, "medium": 1.81}},
	}

	fmt.Println(menu)

}
