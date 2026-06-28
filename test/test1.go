package main

import "fmt"

func main() {
	var c [52]int
	s := "zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"

	for i := 0; i < len(s); i++ {
		c[s[i]-'A']++
	}

	for i := 0; i < len(c); i++ {
		fmt.Printf("%d\n", c[i])
	}

}
