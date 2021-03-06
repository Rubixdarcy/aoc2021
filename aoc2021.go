package main

import (
    "flag"
    "fmt"
    "strconv"
    . "github.com/Rubixdarcy/aoc2021/internal"
)

type challenge struct { day, part int }

func chal(day, part int) challenge {
    return challenge { day, part }
}

var challenges = map[challenge]func(){
    chal(1, 1): Day1Part1,
    chal(1, 2): Day1Part2,
    chal(2, 1): Day2Part1,
    chal(2, 2): Day2Part2,
    chal(3, 1): Day3Part1,
    chal(3, 2): Day3Part2,
    chal(4, 1): Day4Part1,
    chal(4, 2): Day4Part2,
    chal(5, 1): Day5Part1,
    chal(5, 2): Day5Part2,
    chal(6, 1): Day6Part1,
    chal(6, 2): Day6Part2,
    chal(7, 1): Day7Part1,
    chal(7, 2): Day7Part2,
    chal(8, 1): Day8Part1,
    chal(8, 2): Day8Part2,
    chal(9, 1): Day9Part1,
    chal(9, 2): Day9Part2,
    chal(10, 1): Day10Part1,
    chal(10, 2): Day10Part2,
}

func main() {
    flag.Parse()
    args := flag.Args()
    if len(args) < 2 {
        panic("Day and Part args required")
    }

    day, part := atoi(args[0]), atoi(args[1])
    sol, found := challenges[chal(day, part)]
    if !found {
        panic(fmt.Errorf("No solution for day %v part %v", day, part))
    }

    sol()
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func atoi(s string) int {
    i, err := strconv.Atoi(s)
    check(err)
    return i
}
