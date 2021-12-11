package internal

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

func Day9Part1() {
    f, err := os.Open("input/day09.txt")
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)

    var grid [][]int
    for scanner.Scan() {
        var line []int
        for _, c := range scanner.Text() {
            line = append(line, int(c) - int('0'))
        }
        grid = append(grid, line)
    }

    total := 0
    for y, line := range grid {
        for x, val := range line {
            if isLowPoint(grid, x, y) {
                total += val + 1
            }
        }
    }
    fmt.Println(total)
}

var neighbours = [...][2]int{
    [2]int{ -1,  0 },
    [2]int{  1,  0 },
    [2]int{  0, -1 },
    [2]int{  0,  1 },
}

func isLowPoint(grid [][]int, x, y int) bool {
    center, found := gridLookup(grid, x, y)
    if !found {
        return false
    }
    for _, neighbour := range neighbours {
        side, found := gridLookup(grid, x + neighbour[0], y + neighbour[1])
        if found && center >= side {
            return false
        }
    }
    return true
}

func gridLookup(grid [][]int, x, y int) (int, bool) {
    if y < 0 || y > len(grid) - 1 {
        return -1, false
    }
    line := grid[y]
    if x < 0 || x > len(line) - 1 {
        return -1, false
    }
    return line[x], true
}

func Day9Part2() {
    f, err := os.Open("input/day09.txt")
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)

    var grid [][]int
    for scanner.Scan() {
        var line []int
        for _, c := range scanner.Text() {
            line = append(line, int(c) - int('0'))
        }
        grid = append(grid, line)
    }

    var basinSizes []int
    for y, line := range grid {
        for x, _ := range line {
            if isLowPoint(grid, x, y) {
                basinSizes = append(basinSizes, basinSize(grid, x, y))
            }
        }
    }
    sort.Ints(basinSizes)
    l := len(basinSizes)
    fmt.Println(basinSizes[l - 1] * basinSizes[l - 2] * basinSizes[l - 3])
}

// Flood fill algorithm
func basinSize(grid [][]int, x, y int) int {
    open := [][2]int{ [2]int{ x, y } }
    closed := make(map[[2]int]struct{})

    for len(open) > 0 {
        point := open[len(open) - 1]
        x, y := point[0], point[1]
        open = open[:len(open) - 1]

        _, alreadyClosed := closed[point]
        height, inBounds := gridLookup(grid, x, y)
        if alreadyClosed || !inBounds || height == 9 {
            continue
        }

        closed[point] = struct{}{}

        for _, neighbour := range neighbours {
            open = append(open, [2]int{ x + neighbour[0], y + neighbour[1] })
        }
    }
    return len(closed)
}

