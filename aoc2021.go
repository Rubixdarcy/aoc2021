package main

import (
    "fmt"
    "bufio"
    "io/ioutil"
    "os"
    "strconv"
    gp "github.com/prataprc/goparsec"
)

func main() {
    //day1part1()
    //day1part2()
    day5part1()
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

type point struct { x, y int }
type line struct { start, end point }

// Check if lines are hor/ver and canonicise them if they are
func (l line) isHorizontal() (line, bool) {
    y := l.start.y
    canonical := line{
        start: point{x: min(l.start.x, l.end.x), y: y},
        end: point{x: max(l.start.x, l.end.x), y: y},
    }
    return canonical, y == l.end.y
}
func (l line) isVertical() (line, bool) {
    x := l.start.x
    canonical := line{
        start: point{x: x, y: min(l.start.y, l.end.y)},
        end: point{x: x, y: max(l.start.y, l.end.y)},
    }
    return canonical, x == l.end.x
}

func intersectHorizontal(line1, line2 line, xs map[point]struct{}) {
    if y := line1.start.y; y == line2.start.y {
        for x := max(line1.start.x, line2.start.x); x <= min(line1.end.x, line2.end.x); x++ {
            xs[point{ x: x, y: y}] = struct{}{}
        }
    }
}

func intersectVertical(line1, line2 line, xs map[point]struct{}) {
    if x := line1.start.x; x == line2.start.x {
        for y := max(line1.start.y, line2.start.y); y <= min(line1.end.y, line2.end.y); y++ {
            xs[point{x: x, y: y}] = struct{}{}
        }
    }
}

func intersectHorVer(hor, ver line, xs map[point]struct{}) {
    x := ver.start.x
    y := hor.start.y

    intersect := hor.start.x <= x && x <= hor.end.x &&
        ver.start.y <= y && y <= ver.end.y

    if intersect {
        xs[point{x: x, y: y}] = struct{}{}
    }
}

func day5part1() {
    input, err := ioutil.ReadFile("input/d05p1.txt")
    check(err)

    lines := parseLines(input)
    var horizontal, vertical []line

    for _, line := range lines {
        if canonical, is := line.isHorizontal(); is {
            horizontal = append(horizontal, canonical)
            continue
        }
        if canonical, is := line.isVertical(); is {
            vertical = append(vertical, canonical)
        }
    }

    xs := make(map[point]struct{})
    for i, line1 := range horizontal {
        for _, line2 := range horizontal[i + 1:] {
            intersectHorizontal(line1, line2, xs)
        }
    }
    for i, line1 := range vertical {
        for _, line2 := range vertical[i + 1:] {
            intersectVertical(line1, line2, xs)
        }
    }
    for _, hor := range horizontal {
        for _, ver := range vertical {
            intersectHorVer(hor, ver, xs)
        }
    }

    fmt.Println(len(xs))
}

func parseLines(input []byte) []line {
    comma := gp.AtomExact(",", "COMMA")
    arrow := gp.AtomExact(" -> ", "ARROW")
    nl := gp.AtomExact("\n", "NEWLINE")

    lineParser := gp.And(
        func(ns []gp.ParsecNode)(gp.ParsecNode) {
            return line{
                start: point{ x: asInt(ns[0]), y: asInt(ns[2]) },
                end: point{ x: asInt(ns[4]), y: asInt(ns[6]) },
            }
        },
        gp.Int(), comma, gp.Int(), arrow, gp.Int(), comma, gp.Int(),
    )
    linesParser := gp.ManyUntil(
        func(ns []gp.ParsecNode)(gp.ParsecNode) {
            var lines []line
            for _, n := range ns {
                lines = append(lines, n.(line))
            }
            return lines
        },
        lineParser, nl, gp.End(),
    )

    node, _ := linesParser(gp.NewScanner(input))
    return node.([]line)
}

func asInt(node gp.ParsecNode) int {
    return Atoi(node.(*gp.Terminal).Value)
}

func Atoi(s string) int {
    i, err := strconv.Atoi(s)
    check(err)
    return i
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
