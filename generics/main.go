package main

type Number interface {
	~int | ~float64
}

func soma[T int | float64](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

func soma2[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}

	return soma
}

// comparable
func compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "Matheus": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000.50, "Matheus": 2000.60, "Maria": 3000.10}

	println(soma(m))
	println(soma(m2))
	println(compara(10, 10.0))
}
