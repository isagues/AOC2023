package main

import (
    "fmt"
    "os"
    "bufio"
    "unicode"
    "strings"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// one, two, three, four, five, six, seven, eight, nine
var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
    // readFile, err := os.Open("input")
    readFile, err := os.Open("input")
    check(err)
    
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    var sum = 0

    for fileScanner.Scan() {
        
        var sb strings.Builder
        line := fileScanner.Text()
        fmt.Println(line)
        
        for i := 0; i < len(line); i++ {
            var fIndex [9]int
            for w_idx, w:= range words {
                fIndex[w_idx] = strings.Index(line[i:], w)
                // fmt.Printf("fIndex = %d; w= %s, w_indx= %d\n", fIndex[w_idx], w, w_idx)
            }

            // fmt.Print(values)
            var minIdx int = -1
            var min int = fIndex[0]
            for k, v := range fIndex {
                if v == -1 {
                    continue
                }
                // v += i 
                if minIdx == -1 || v <= min {
                    min = v
                    minIdx = k
                }
            }

            if min == 0 && minIdx >= 0 {
                i += len(words[minIdx])-1
                sb.WriteByte(byte(minIdx) + 1 + '0')
            } else {
                sb.WriteByte(line[i])
            }
        }
        
        transformedLine := sb.String()
        
        fmt.Println(transformedLine)

        var first int = -1
        var last int = -1

        for _, c := range (transformedLine) {
            if unicode.IsDigit(c) {
                num := int(c - '0')
                if first == -1 {
                    first = num
                }
                last = num

            }
        }

        sum += first*10 + last
        fmt.Printf("%d\n", first*10 + last)
    }
    fmt.Printf("SUM: %d\n", sum)
  
    readFile.Close()
}