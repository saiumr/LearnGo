package main

import (
	"fmt"
	"sort"
	"bufio"
	"os"
)

func main() {
	// make allocate memory
	ages := make(map[string]int)
	ages = map[string]int{
		"gmq": 23,
		"cs":  22,
		"ls":  22,
		"ccx": 23,
	}

	ages["error"] = ages["error"] + 1
	delete(ages, "error")
	// value == 0 maybe means that key does not exist
	// if age, ok := ages["bob"]; !ok { /* ... */ }

	// map range is not same each time
	/*
	for i := 0; i < 10; i++ {
		for name, age := range ages {
			fmt.Printf("%v, %v\n", name, age)
		}
		fmt.Println()
	}
	*/

	// print map elements by the same order
	var names []string
	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for i := 0; i < 10; i++ {
		for _, name := range names {
			fmt.Println(name, ages[name])
		}
		fmt.Println()
	}

	// empty map is nil
	var mapa map[string]int
	fmt.Println(mapa == nil)
	fmt.Println(len(mapa) == 0)

	// Invalid, no make, no assignment
	// mapa["ff"] = 9
	
	if equal(map[string]int{"A": 0}, map[string]int{"A": 0}) {
		fmt.Println("Yes, equal.")
	} else {
		fmt.Println("No, not equal.")
	}

	// the key in map are not same
	// we can use map key make a set
	seen := make(map[string]bool)  // a set of strings
	input := bufio.NewScanner(os.Stdin) 
	//fmt.Println(input.Scan())  // bool
	//fmt.Println(input.Text())  // text
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
			return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
				return false
		}
	}
	return true
}

