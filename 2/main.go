package main

var t int = 2

func sum(e chan int, d chan struct{}) {
	go func() {
		for {
			select {
			case <-d:
				break
			case n := <-e:
				t = t + n
			}
		}
	}()
}

func main() {
	e := make(chan int)
	d := make(chan struct{})
	sum(e, d)

	l := 1
	p := 2
	c := 0

	for {
		c = p + l
		l = p
		p = c

		if c > 4000000 {
			break
		}

		if c%2 == 0 {
			e <- c
			println(c)
		}
	}

	d <- struct{}{}
	println(t)
}
