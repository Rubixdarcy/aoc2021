package internal

import (
    "fmt"
    "sort"
)


func Day7Part1() {
    ns := lineOfCsv2Slice("input/day07.txt")
    sort.Ints(ns)

    mid := ns[len(ns) / 2]

    fuel := 0
    for _, n := range ns {
        fuel += abs(mid - n)
    }
    fmt.Println(fuel)
}

func Day7Part2() {
    ns := lineOfCsv2Slice("input/day07.txt")

    sum := 0
    for _, n := range ns {
        sum += n
    }
    // Mean is the least squares estimator
    mean := int(float64(sum) / float64(len(ns)))

    fuel := 0
    for _, n := range ns {
        dist := abs(mean - n)
        // Triangular number
        fuel += dist * (dist + 1) / 2
    }
    fmt.Println("mean", mean)
    fmt.Println(fuel)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
