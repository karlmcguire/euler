package main

func prime(n int) bool {
	r := n
	for i := 2; i < r; i++ {
		if n%i == 0 {
			return false
		}
		r = n / i
	}
	return true
}
func main() {
	sum := 0

	for i := 2; i < 2000000; i++ {
		if prime(i) {
			sum = sum + i
		}
	}

	println(sum)
}
