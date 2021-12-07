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
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
