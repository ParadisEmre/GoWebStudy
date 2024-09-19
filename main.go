package main

import "fmt"

func main() {
	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30}

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:4] // doesn't include pos 4 element
	rangeTwo := names[2:]  //includes the last element
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	fmt.Println(rangeOne)

	for index, value := range names {
		fmt.Printf("the index is %v and the value is %v \n", index, value)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	menu := map[string]float64{"pie": 5.95, "cake": 3.50, "ice cream": 2.50}

	fmt.Printf("the value of pie is %v \n", menu["pie"])
}
