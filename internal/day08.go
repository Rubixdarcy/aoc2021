package internal

import (
    "fmt"
    gp "github.com/prataprc/goparsec"
    "io/ioutil"
)

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
            //fmt.Println(word.(*gp.Terminal).Value)
            switch len(word.(*gp.Terminal).Value) {
            case 2, 3, 4, 7:
                count += 1
            }
        }
    }
    fmt.Println(count)
}

func Day8Part2() {
}

//func parseLines(input []byte) []line {
//    comma := gp.AtomExact(",", "COMMA")
//    arrow := gp.AtomExact(" -> ", "ARROW")
//    nl := gp.AtomExact("\n", "NEWLINE")
//
//    lineParser := gp.And(
//        func(ns []gp.ParsecNode)(gp.ParsecNode) {
//            return line{
//                start: point{ x: asInt(ns[0]), y: asInt(ns[2]) },
//                end: point{ x: asInt(ns[4]), y: asInt(ns[6]) },
//            }
//        },
//        gp.Int(), comma, gp.Int(), arrow, gp.Int(), comma, gp.Int(),
//    )
//    linesParser := gp.ManyUntil(
//        func(ns []gp.ParsecNode)(gp.ParsecNode) {
//            var lines []line
//            for _, n := range ns {
//                lines = append(lines, n.(line))
//            }
//            return lines
//        },
//        lineParser, nl, gp.End(),
//    )
//
//    node, _ := linesParser(gp.NewScanner(input))
//    return node.([]line)
//}
