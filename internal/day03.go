package internal

import (
    "fmt"
    "bufio"
    "os"
)

const nbits = 12

func Day3Part1() {
    f, err := os.Open("input/day03.txt")
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)

    var bits [nbits]int
    count := 0

    for scanner.Scan() {
        count++
        s := scanner.Text()

        for i, c := range s {
            if c == '1' {
                bits[i]++
            }
        }
    }

    var gamma, epsilon int
    threshold := count / 2
    for i, n := range bits {
        if n >= threshold {
            gamma += 1 << (nbits - i - 1)
        }
    }
    epsilon = 1 << nbits - 1 - gamma

    fmt.Println(epsilon * gamma)
}

func Day3Part2() {
}
