package main

import "fmt"

// slice example
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7}
	s = s[2:6]
	fmt.Println("SLICE")
	fmt.Println(s)

	//maps example
	m := map[string]float64{
		"c": 1,
		"d": 2,
		"e": 3,
	}
	m["f"] = 4       // for adding
	cvalue := m["c"] // to retrive
	fmt.Println("MAPS")
	fmt.Println(cvalue)
	delete(m, "e") // to delete
	fmt.Println(m)

	for variables, values := range m {
		fmt.Println(variables, values)
	}
}
