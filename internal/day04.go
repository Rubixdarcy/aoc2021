package internal

import (
    "fmt"
    "io/ioutil"
    gp "github.com/prataprc/goparsec"
)

const boardSize = 5
const boardArea = boardSize * boardSize

var checkStartOffset [10][2]int = [10][2]int{
    [2]int{  0, 1 },
    [2]int{  5, 1 },
    [2]int{ 10, 1 },
    [2]int{ 15, 1 },
    [2]int{ 20, 1 },
    [2]int{  0, 5 },
    [2]int{  1, 5 },
    [2]int{  2, 5 },
    [2]int{  3, 5 },
    [2]int{  4, 5 },
}

type board struct { marked [boardArea]bool; nums []int }

// Mark off number and return true if Bingo
func (b *board) markNumber(num int) bool {
    idx, found := indexOf(b.nums, num)
    if found {
        b.marked[idx] = true
    }

    // Check for bingo
    for _, startOffset := range checkStartOffset {
        bingo := true
        for i := 0; i < boardSize; i++ {
            if !b.marked[startOffset[0] + i * startOffset[1]] {
                bingo = false
                break
            }
        }
        if bingo {
            return true
        }
    }
    return false
}

func (b *board) value() int {
    total := 0
    for i, marked := range b.marked {
        if !marked {
            total += b.nums[i]
        }
    }
    return total
}

func playBingo(numbers []int, boards []*board) (int, bool) {
    for _, number := range numbers {
        for _, board := range boards {
            if board.markNumber(number) {
                return board.value() * number, true
            }
        }
    }
    return 0, false
}

func playBingoUntilComplete(numbers []int, boards []*board) (int, bool) {
    score := 0
    winner := false
    completeBoards := make([]bool, len(boards))
    for _, number := range numbers {
        for i, board := range boards {
            if !completeBoards[i] && board.markNumber(number) {
                score = board.value() * number
                winner = true
                completeBoards[i] = true
            }
        }
    }
    return score, winner
}

func Day4Part1() {
    input, err := ioutil.ReadFile("input/day04.txt")
    check(err)

    numbers, boardNums := parseBingoInput(input)

    var boards []*board
    boardNumsCount := len(boardNums)
    for i := 0; i < boardNumsCount; i += boardArea {
        boards = append(boards, &board{ nums: boardNums[i:i + boardArea]})
    }

    score, winner := playBingo(numbers, boards)

    if !winner {
        fmt.Println("No winner")
    }
    fmt.Println(score)
}

func Day4Part2() {
    input, err := ioutil.ReadFile("input/day04.txt")
    check(err)

    numbers, boardNums := parseBingoInput(input)

    var boards []*board
    boardNumsCount := len(boardNums)
    for i := 0; i < boardNumsCount; i += boardArea {
        boards = append(boards, &board{ nums: boardNums[i:i + boardArea]})
    }

    score, winner := playBingoUntilComplete(numbers, boards)

    if !winner {
        fmt.Println("No winner")
    }
    fmt.Println(score)
}

func parseBingoInput(input []byte) ([]int, []int) {
    parser := gp.And(nil, csvLineParser, whitespaceIntsParser)

    node, _ := parser(gp.NewScanner(input))
    result := node.([]gp.ParsecNode)

    return result[0].([]int), result[1].([]int)
}

func indexOf(ns []int, n int) (int, bool) {
    for i, elem := range ns {
        if elem == n {
            return i, true
        }
    }
    return -1, false
}
