package main

func main() {
	n := 1
	g := false

	for {
		for i := 1; i <= 20; i++ {
			if n%i != 0 {
				g = false
				break
			} else {
				g = true
			}
		}
		if g {
			break
		} else {
			n++
		}
	}

	println(n)
}
