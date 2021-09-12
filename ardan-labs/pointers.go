package ardanlabs

import "testing"

func TestC(t *testing.T) {
	count := 10

	println("count: \tValue Of[", count, "]\t\tAddr Of[", &count, "]")
	increment(count)

	println("count: \tValue Of[", count, "]\t\tAddr Of[", &count, "]")
	increment1(&count)
	println("count: \tValue Of[", count, "]\t\tAddr Of[", &count, "]")
}

func increment(inc int) {
	inc++
	println("count: \tValue Of[", inc, "]\t\tAddr Of[", &inc, "]")
}

func increment1(inc *int) {
	*inc++
	println("count: \tValue Of[", inc, "]\t\tAddr Of[", &inc, "]")
}