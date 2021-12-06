package internal

import (
    "bufio"
    "strconv"
    "os"
    gp "github.com/prataprc/goparsec"
)

func asInt(node gp.ParsecNode) int {
    return atoi(node.(*gp.Terminal).Value)
}

func atoi(s string) int {
    i, err := strconv.Atoi(s)
    check(err)
    return i
}

func linesOfNumbers2Slice(filename string) []int {
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

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func cmp(a, b int) int {
    if a < b {
        return -1
    }
    if a > b {
        return 1
    }
    return 0
}
