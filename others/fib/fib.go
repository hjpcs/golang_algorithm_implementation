package fib

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	sum := 0
	prev := 0
	cur := 1
	for i := 2; i <= n; i++ {
		sum = prev + cur
		prev = cur
		cur = sum
	}
	return sum
}
