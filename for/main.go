package main

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []string{"um", "dois", "tres"}
	for k, v := range numeros {
		println(k, v)
	}

	for _, val := range numeros {
		println(val)
	}

	t := 0
	for t < 10 {
		println(t)
		t++
	}
}
