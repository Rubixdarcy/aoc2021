package internal

import (
    "fmt"
    "bufio"
    "os"
    "sort"
)

const (
    delimShapeParen = iota
    delimShapeSquare = iota
    delimShapeBrace = iota
    delimShapeAngle = iota
)

const (
    delimDirOpen = iota
    delimDirClose = iota
)


type delimClass struct { shape, dir int }

var runeDelimClasses = map[rune]delimClass{
    '(': delimClass{ delimShapeParen,  delimDirOpen  },
    ')': delimClass{ delimShapeParen,  delimDirClose },
    '[': delimClass{ delimShapeSquare, delimDirOpen  },
    ']': delimClass{ delimShapeSquare, delimDirClose },
    '{': delimClass{ delimShapeBrace,  delimDirOpen  },
    '}': delimClass{ delimShapeBrace,  delimDirClose },
    '<': delimClass{ delimShapeAngle,  delimDirOpen  },
    '>': delimClass{ delimShapeAngle,  delimDirClose },
}

var delimClassScores = map[delimClass]int{
    runeDelimClasses[')']: 3,
    runeDelimClasses[']']: 57,
    runeDelimClasses['}']: 1197,
    runeDelimClasses['>']: 25137,
}

var shapeCompleteScores = map[int]int {
    delimShapeParen: 1,
    delimShapeSquare: 2,
    delimShapeBrace: 3,
    delimShapeAngle: 4,
}

func Day10Part1() {
    f, err := os.Open("input/day10.txt")
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)

    total := 0
    for scanner.Scan() {
        var shapes []int
        for _, r := range scanner.Text() {
            class := runeDelimClasses[r]

            if class.dir == delimDirOpen {
                shapes = append(shapes, class.shape)
            } else {
                if shapes[len(shapes) - 1] != class.shape {
                    // Invalid syntax
                    total += delimClassScores[class]
                    break
                }
                shapes = shapes[:len(shapes) - 1]
            }
        }
    }

    fmt.Println(total)
}

func Day10Part2() {
    f, err := os.Open("input/day10.txt")
    check(err)
    defer f.Close()
    scanner := bufio.NewScanner(f)

    var scores []int
    for scanner.Scan() {
        var shapes []int
        invalid := false
        for _, r := range scanner.Text() {
            class := runeDelimClasses[r]

            if class.dir == delimDirOpen {
                shapes = append(shapes, class.shape)
            } else {
                if shapes[len(shapes) - 1] != class.shape {
                    invalid = true
                    break
                }
                shapes = shapes[:len(shapes) - 1]
            }
        }
        if !invalid {
            score := 0
            for i, _ := range shapes {
                score = score * 5 + shapeCompleteScores[shapes[len(shapes) - i - 1]]
            }
            scores = append(scores, score)
        }
    }

    sort.Ints(scores)
    fmt.Println(scores[len(scores) / 2])
}
