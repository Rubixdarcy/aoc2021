package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
    fmt.Println("hello world")
    day1part1()
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
