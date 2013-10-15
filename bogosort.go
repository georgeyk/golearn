package main

import "fmt"
import "math/rand"
import "os"
import "strconv"

func sorted(elements []int) bool {
	size := len(elements)
	for i := 0; i < size-1; i++ {
		x := elements[i]
		y := elements[i+1]
		if x > y {
			return false
		}
	}
	return true
}

// fisher-yates shuffle
func shuffle(elements []int) []int {
	size := len(elements)
	for i := 0; i < size-1; i++ {
		j := rand.Intn(size)
		tmp := elements[i]
		elements[i] = elements[j]
		elements[j] = tmp
	}
	return elements
}

func bogosort(elements []int) {
	for i := 0; ; i++ {
		if sorted(elements) {
			fmt.Println("Sorted:", elements, "Iters:", i)
			break
		}
		shuffle(elements)
	}
}

func main() {
	size := len(os.Args[1:])
	if size <= 0 {
		fmt.Println("You should pass a sequence of integers to sort.")
		os.Exit(1)
	}

	elements := make([]int, size)

	for i, element := range os.Args[1:] {
		element, err := strconv.Atoi(element)

		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(1)
		}

		elements[i] = element
	}

	bogosort(elements[:])
}
