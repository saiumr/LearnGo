package main
import (
	"fmt"
)

type Currency int
const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	// r := [...]int{99: -1}
	// defined 100 variables, r[99] = -1 and others l 	
	// arr := [...]int{2,3,4,5}
	// var arr [3]int = [...]int{7, 8, 9} also takes effect, but not elegant enough
	var arr [3]int = [3]int{7, 8, 9}  
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	
	for k, v := range arr {
		fmt.Printf("k:%d v:%d\n", k, v)
	}
	fmt.Printf("%v\n", symbol)
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
}
