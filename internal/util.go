package internal

import (
    "bufio"
    "strconv"
    gp "github.com/prataprc/goparsec"
    "strings"
    "os"
)

func asInt(node gp.ParsecNode) int {
    return atoi(node.(*gp.Terminal).Value)
}

func asInts(nodes []gp.ParsecNode) gp.ParsecNode {
    var nums []int
    for _, n := range nodes {
        nums = append(nums, asInt(n))
    }
    return nums
}

var commaParser = gp.AtomExact(",", "COMMA")
var csvLineParser = gp.Many(asInts, gp.Int(), commaParser)
var whitespaceIntsParser = gp.Many(asInts, gp.Int())

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

func lineOfCsv2Slice(filename string) []int {
    // Adapted courtesy of https://stackoverflow.com/questions/5884154/read-text-file-into-string-array-and-write
    f, err := os.Open(filename)
    check(err)
    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Scan()


    var nums []int
    for _, s := range strings.Split(scanner.Text(), ",") {
        n, err := strconv.Atoi(s)
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
