package internal

import (
    "fmt"
    gp "github.com/prataprc/goparsec"
    "io/ioutil"
)

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

func (l line) yieldCoords(counts map[point]int) {
    dx := cmp(l.end.x, l.start.x)
    dy := cmp(l.end.y, l.start.y)

    p := l.start
    for {
        counts[p]++
        if p == l.end {
            break
        }
        p.x, p.y = p.x + dx, p.y + dy
    }
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

func Day5Part1() {
    input, err := ioutil.ReadFile("input/day05.txt")
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

// I solved part 1 efficiently. Doing part 2 efficiently would be convoluted
// so instead I'm using a huge hashmap that maps all coords to their count.
func Day5Part2() {
    input, err := ioutil.ReadFile("input/day05.txt")
    check(err)

    lines := parseLines(input)

    counts := make(map[point]int)
    for _, line := range lines {
        line.yieldCoords(counts)
    }

    n := 0
    for _, count := range counts {
        if count > 1 {
            n++
        }
    }

    fmt.Println(n)
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
