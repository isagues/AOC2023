package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

type MapFunc[A any, B any] func(A) B
type ReduceFunc[A any, B any] func(B, A) B

func Map[A any, B any](input []A, m MapFunc[A, B]) []B {
    output := make([]B, len(input))
    for i, element := range input {
        output[i] = m(element)
    }
    return output
}

func Reduce[A any, B any](input []A, r ReduceFunc[A, B], init B) B {
    var output B = init
    for _, element := range input {
        output = r(output, element)
    }
    return output
}

type Number struct {
    value int
    literal string
    spelled string
}

type Position struct {
    value int
    idx int
}

func optionalMin(v1 int, v2 int) int {
    if v1 == -1 {
        return v2
    }
    if v2 == -1 {
        return v1
    }
    return min(v1, v2)
}

func optionalMax(v1 int, v2 int) int {
    if v1 == -1 {
        return v2
    }
    if v2 == -1 {
        return v1
    }
    return max(v1, v2)
}

func minPosition(p1 Position, p2 Position) Position { 
    if p1.idx != -1 && p1.idx < p2.idx {
        return p1
    }
    if p2.idx == -1 {
        return p1
    }
    return p2
}

func maxPosition(p1 Position, p2 Position) Position { 
    if p1.idx != -1 && p1.idx > p2.idx {
        return p1
    }
    if p2.idx == -1 {
        return p1
    }
    return p2
}

func getFirstPosition(line string, num Number) Position {
    literalIdx := strings.Index(line, num.literal)
    spelledIdx := strings.Index(line, num.spelled)
    return Position{num.value, optionalMin(spelledIdx, literalIdx)}
}

func getLastPosition(line string, num Number) Position {
    literalIdx := strings.LastIndex(line, num.literal)
    spelledIdx := strings.LastIndex(line, num.spelled)
    return Position{num.value, optionalMax(spelledIdx, literalIdx)}
}


func main() {
    numbers := []Number{{1, "1", "one"}, {2, "2", "two"}, {3, "3", "three"}, {4, "4", "four"}, {5, "5", "five"}, {6, "6", "six"}, {7, "7", "seven"}, {8, "8", "eight"}, {9, "9", "nine"}}
    var sum int = 0

    readFile, err := os.Open("input")
    // readFile, err := os.Open("smallinput")
    check(err)
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    for fileScanner.Scan() {
        
        line := fileScanner.Text()
        fmt.Println(line)

        getF := func(num Number) Position { 
            return getFirstPosition(line, num)
        }

        firstPositions := Map(numbers, getF)
        firstPos := Reduce(firstPositions, minPosition, Position{-1, -1})
        first := firstPos.value

        getL := func(num Number) Position { 
            return getLastPosition(line, num)
        }

        lastPositions := Map(numbers, getL)
        lastPos := Reduce(lastPositions, maxPosition, Position{-1, -1})
        last := lastPos.value
    
        sum += first * 10 + last
    }
    fmt.Printf("SUM: %d\n", sum)
  
    readFile.Close()
}