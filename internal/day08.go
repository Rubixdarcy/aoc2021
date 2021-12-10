package internal

import (
    "fmt"
    gp "github.com/prataprc/goparsec"
    "io/ioutil"
    "math/bits"
)

var pow10 = [...]int{ 1, 10, 100, 1000 }

func Day8Part1() {
    input, err := ioutil.ReadFile("input/day08.txt")
    check(err)

    scanner := gp.NewScanner(input)
    parser := gp.Many(nil,
        gp.And(nil,
            gp.Many(nil, gp.Ident()),
            gp.AtomExact(" | ", "BAR"),
            gp.ManyUntil(nil, gp.Ident(), gp.AtomExact("\n", "NEWLINE")),
    ))
    node, _ := parser(scanner)

    count := 0
    for _, line := range node.([]gp.ParsecNode) {
        for _, word := range line.([]gp.ParsecNode)[2].([]gp.ParsecNode) {
            switch len(word.(*gp.Terminal).Value) {
            case 2, 3, 4, 7:
                count += 1
            }
        }
    }
    fmt.Println(count)
}

type easyMasks struct { one, seven, four, eight uint }

func digitToBitmask(digit string) uint {
    var result uint
    for _, c := range digit {
        result += 1 << (int(c) - 97)
    }
    return result
}

func maskContains(a, b uint) bool {
    notA := 1 << 7 - 1 - a
    return notA & b == 0
}

func determineDigit(easy *easyMasks, mask uint) int {
    switch mask {
    case easy.one:
        return 1
    case easy.seven:
        return 7
    case easy.four:
        return 4
    case easy.eight:
        return 8
    }

    if bits.OnesCount(mask) == 6 {
        if maskContains(mask, easy.four) {
            return 9
        }
        if maskContains(mask, easy.one) {
            return 0
        }
        return 6
    }

    if maskContains(mask, easy.one) {
        return 3
    }
    if bits.OnesCount(mask & easy.four) == 2 {
        return 2
    }
    return 5
}

func Day8Part2() {
    input, err := ioutil.ReadFile("input/day08.txt")
    check(err)

    scanner := gp.NewScanner(input)
    parser := gp.Many(nil,
        gp.And(nil,
            gp.Many(nil, gp.Ident()),
            gp.AtomExact(" | ", "BAR"),
            gp.ManyUntil(nil, gp.Ident(), gp.AtomExact("\n", "NEWLINE")),
    ))
    node, _ := parser(scanner)

    total := 0
    for _, line := range node.([]gp.ParsecNode) {
        // Easy pass
        var easy easyMasks
        for _, word := range line.([]gp.ParsecNode)[0].([]gp.ParsecNode) {
            mask := digitToBitmask(word.(*gp.Terminal).Value)
            switch bits.OnesCount(mask) {
            case 2:
                easy.one = mask
            case 3:
                easy.seven = mask
            case 4:
                easy.four = mask
            case 7:
                easy.eight = mask
            }
        }

        // Calculation
        for i, word := range line.([]gp.ParsecNode)[2].([]gp.ParsecNode) {
            mask := digitToBitmask(word.(*gp.Terminal).Value)
            digit := determineDigit(&easy, mask)
            total += digit * pow10[3 - i]
        }
    }
    fmt.Println(total)
}

