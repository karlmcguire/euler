package main

func main() {
	ssq, sqs := 0, 0

	for i := 1; i <= 100; i++ {
		ssq = ssq + (i * i)
		sqs = sqs + i
	}
	sqs = sqs * sqs

	println(sqs - ssq)
}
