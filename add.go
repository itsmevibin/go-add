package main

func main() {

}

func add(val []int) int {
	var sum int
	for _, v := range val {
		sum += v
	}
	return sum
}
