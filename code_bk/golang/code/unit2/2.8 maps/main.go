package main

import "fmt"

func main() {
	// TH => Thailand
	// EN => English
	var langs map[string]string
	if langs == nil {
		fmt.Println("nil map")
	}

	langs2 := map[string]string{
		"TH": "Thailand",
		"EN": "English",
	}
	fmt.Println(langs2)
	fmt.Println(langs2["TH"]) // Thailand

	for k, v := range langs2 {
		fmt.Println(k, v)
	}
}
