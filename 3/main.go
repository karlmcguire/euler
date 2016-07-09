package main

func main() {
	n := 600851475143

	var i int

	for i = 2; i < n; i++ {
		if n%i == 0 {
			n = n / i
			i--
		}
	}

	println(i)
}
