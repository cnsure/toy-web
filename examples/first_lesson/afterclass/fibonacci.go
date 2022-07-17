package main

func main() {
	println(fibonacci(2))
}

func fibonacci(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
