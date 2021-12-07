package internal

import (
    "fmt"
    "bufio"
    "os"
    "sort"
    "strconv"
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
    f, err := os.Open("input/day03.txt")
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    var nums []string
    for scanner.Scan() {
        nums = append(nums, scanner.Text())
    }
    sort.Strings(nums)

    // Oxygen
    var oxygen int64
    oxygenSlice := nums
    for i := 0; i < nbits; i++ {
        first1 := indexOfFirst1InCol(oxygenSlice, i)
        if first1 * 2 <= len(oxygenSlice) {
            oxygenSlice = oxygenSlice[first1:]
        } else {
            oxygenSlice = oxygenSlice[:first1]
        }
        if len(oxygenSlice) == 1{
            oxygen, err = strconv.ParseInt(oxygenSlice[0], 2, 64)
            check(err)
            break
        }
    }

    // CO2
    var co2 int64
    co2Slice := nums
    for i := 0; i < nbits; i++ {
        first1 := indexOfFirst1InCol(co2Slice, i)
        if first1 * 2 <= len(co2Slice) {
            co2Slice = co2Slice[:first1]
        } else {
            co2Slice = co2Slice[first1:]
        }
        if len(co2Slice) == 1{
            co2, err = strconv.ParseInt(co2Slice[0], 2, 64)
            check(err)
            break
        }
    }

    fmt.Println(oxygen * co2)
}

func indexOfFirst1InCol(nums []string, col int) int {
    for i, s := range nums {
        if s[col] == '1' {
            return i
        }
    }
    return len(nums)
}
