package main
import (
	"fmt"
	"os"
	"strings"
)

func PrintArgs() {
	for index, str := range os.Args[0:] {
		fmt.Println(index, "," ,str)
	}
}

func main() {
//	var s, sep string	
//	for i := 1; i < len(os.Args); i++ {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("### PrintArgs ###")
 	PrintArgs()
	fmt.Println("### strings.Join(os.Args[0:], \" \") ###")
	fmt.Println(strings.Join(os.Args[0:], " "))
}

