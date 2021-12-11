package internal

import (
    "fmt"
    "bufio"
    "os"
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

}
