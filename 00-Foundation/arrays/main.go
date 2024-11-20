package main

import "fmt"

type ID int

const a =  "Hello World"

var (
	b bool    = true
	c int     = 10
	d string  = "Matheus"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 10 
	meuArray[1] = 20 
	meuArray[2] = 33 

	//fmt.Printf("O tipo de E é %T", e)
	//fmt.Println(len(meuArray) -1)
	for i, v := range meuArray{
		fmt.Printf("O valor do indice %d é %d\n", i, v)
	}

}