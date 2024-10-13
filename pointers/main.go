package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {

	var1 := 10
	var2 := 20

	println(soma(&var1, &var2))
	println(var1)

	// // Mmória -> Endereço -> Valor
	// a := 10
	// //println(a)
	// //println(&a)

	// var ponteiro *int = &a
	// //println(ponteiro)
	// *ponteiro = 20
	// println(a)

	// b := &a
	// *b = 30

	// println(b)
	// println(*b)
	// println(a)
}
