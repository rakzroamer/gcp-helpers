package main

import ( "fmt"
 	"os")

func main() {

	n := 42
	s := fmt.Sprintf("%d", n)
	s1 := fmt.Printf("%d", n)
	fmt.Printf("s = %s (type %T)\n",s ,s)
	fmt.Printf("s1 = %s (type %T)\n",s1 ,s1)
}
