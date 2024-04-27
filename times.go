package main

import (
    "fmt"
    "math/rand"
	"time"
)

// number of questions to ask
const NUM_QUESTIONS = 5

// min and max numbers to quiz on, e.g. times tables will range from 2*2 to 12*12
const MIN = 2
const MAX = 12


func main() {
	total_correct := 0
	start_time := time.Now().UnixMilli()
	for i := 0; i < NUM_QUESTIONS; i++ {
		total_correct += question()
	}
	elapsed := float32(time.Now().UnixMilli() - start_time)/1000
	fmt.Println("------------\nfinal stats: ", total_correct, " out of ", NUM_QUESTIONS, " questions answered correctly.") 
	fmt.Printf("you took %.3f seconds to answer, or %.3f seconds per question\n", elapsed, elapsed/float32(NUM_QUESTIONS))
}

// returns 1 if answered correctly, 0 if wrong
func question () int {
	v1 := rangeIn(MIN,MAX)
	v2 := rangeIn(MIN,MAX)
    fmt.Println(v1, " * ", v2)
	var answer int
	fmt.Scan(&answer)
	if answer == (v1 * v2) {
		fmt.Println("correct!")
		return 1
	} else {
		fmt.Println("wrong! answer is ", v1 * v2)
		return 0
	}
}

func rangeIn(low, hi int) int {
	rand.Seed(time.Now().UnixNano())
    return low + rand.Intn(hi-low)
}
