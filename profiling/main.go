package profiling

// Ejecutar las pruebas de consumo de memoria y rendimiento de cpu.
// go test -bench=. -benchmem -cpuprofile profile.out

// Entrar en la cli interactiva pprof
// go tool pprof profile.out

func FibonacciFastest(number int) int {
	if number <= 1 {
		return number
	}

	fibonacci := 0
	memo := make(map[int]int)

	for i := 0; i < number - 1; i++ {
		if i <= 1 {
			memo[i] = i
		} else {
			memo[i] = memo[i - 1] + memo[i - 2]
		}

		fibonacci = fibonacci + memo[i]
	}

	return fibonacci + 1
}

func FibonacciSlowest(number int) int {
	if number <= 1 {
		return number
	}

	return FibonacciSlowest(number - 1) + FibonacciSlowest(number - 2)
}
