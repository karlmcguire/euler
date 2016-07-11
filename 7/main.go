package main

func prime(n int) bool {
	if n < 2 {
		return false
	}

	if n == 2 {
		return true
	}

	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	c, n := 0, 0

	for {
		if prime(n) {
			c = c + 1

			if c == 10001 {
				break
			}
		}

		n++
	}

	println(n, prime(n))
}
