package internal

import (
    "fmt"
    "io/ioutil"
    gp "github.com/prataprc/goparsec"
)


type direction int
const (
    forward direction = iota
    up
    down
)

func Direction(s string) direction {
    switch s {
    case "forward":
        return forward
    case "up":
        return up
    case "down":
        return down
    default:
        panic(fmt.Errorf("Unknown direction '%s'", s))
    }
}

type inst struct { dir direction; dist int }

func Day2Part1() {
    input, err := ioutil.ReadFile("input/d02p1.txt")
    check(err)

    insts := parseLinesDay5(input)

    var horizontal, depth int
    for _, inst := range insts {
        switch inst.dir {
        case forward:
            horizontal += inst.dist
        case up:
            depth -= inst.dist
        case down:
            depth += inst.dist
        }
    }

    fmt.Println(horizontal * depth)
}

func Day2Part2() {
    input, err := ioutil.ReadFile("input/d02p2.txt")
    check(err)

    insts := parseLinesDay5(input)

    var horizontal, depth, aim int
    for _, inst := range insts {
        switch inst.dir {
        case forward:
            horizontal += inst.dist
            depth += inst.dist * aim
        case up:
            aim -= inst.dist
        case down:
            aim += inst.dist
        }
    }

    fmt.Println(horizontal * depth)
}

func parseLinesDay5(input []byte) []inst {
    space := gp.AtomExact(" ", "SPACE")
    nl := gp.AtomExact("\n", "NEWLINE")

    lineParser := gp.And(
        func(ns []gp.ParsecNode)(gp.ParsecNode) {
            return inst{
                dir: Direction(ns[0].(*gp.Terminal).Value),
                dist: asInt(ns[2]),
            }
        },
        gp.Ident(), space, gp.Int(),
    )
    linesParser := gp.ManyUntil(
        func(ns []gp.ParsecNode)(gp.ParsecNode) {
            var insts []inst
            for _, n := range ns {
                insts = append(insts, n.(inst))
            }
            return insts
        },
        lineParser, nl, gp.End(),
    )

    node, _ := linesParser(gp.NewScanner(input))
    return node.([]inst)
}
