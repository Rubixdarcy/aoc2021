package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    //day1part1()
    day1part2()
}

func day1part1() {
    nums := LinesOfNumbers2Slice("input/d01p1.txt")

    count := 0
    for i, n := range nums[1:] {
        if nums[i] < n {
            count++
        }
    }
    fmt.Println(count)
}

func day1part2() {
    nums := LinesOfNumbers2Slice("input/d01p2.txt")

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

func LinesOfNumbers2Slice(filename string) []int {
    // Adapted courtesy of https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
    f, err := os.Open(filename)
    check(err)
    defer f.Close()

    var nums []int
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        n, err := strconv.Atoi(scanner.Text())
        check(err)
        nums = append(nums, n)
    }

    return nums
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
