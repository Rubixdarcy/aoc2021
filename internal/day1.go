package internal

import "fmt"

func Day1Part1() {
    nums := linesOfNumbers2Slice("input/d01p1.txt")

    count := 0
    for i, n := range nums[1:] {
        if nums[i] < n {
            count++
        }
    }
    fmt.Println(count)
}

func Day1Part2() {
    nums := linesOfNumbers2Slice("input/d01p2.txt")

    var count, prev int
    prev = nums[0] + nums[1] + nums[2]
    for i, _ := range nums[3:] {
        next := nums[i + 1] + nums[i + 2] + nums[i + 3]
        if prev < next {
            count++
        }
        prev = next
    }
    fmt.Println(count)
}
