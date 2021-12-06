package internal

import "fmt"

type lanternState [9]int

func (s lanternState) nextState() lanternState {
    var next lanternState

    copy(next[:], s[1:])
    next[6] += s[0]
    next[8] = s[0]

    return next
}

func Day6Part1() {
    ns := lineOfCsv2Slice("input/d06p1.txt")
    var state lanternState

    for _, n := range ns {
        state[n]++
    }

    for i := 0; i < 80; i++ {
        state = state.nextState()
    }

    count := 0
    for _, n := range state {
        count += n
    }

    fmt.Println(count)
}

func Day6Part2() {
}
