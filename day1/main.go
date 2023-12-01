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

// one, two, three, four, five, six, seven, eight, nine
var words = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var nums = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
    readFile, err := os.Open("input")
    // readFile, err := os.Open("smallinput")
    check(err)
    
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
  
    var sum = 0

    for fileScanner.Scan() {
        
        line := fileScanner.Text()
        fmt.Println(line)
        
        var minIdx, minValue int = -1, -1
        for idx, w:= range words {
            wIdx := strings.Index(line, w)
            
            if wIdx == -1 {
                continue
            }
            // v += i 
            if minIdx == -1 || wIdx <= minIdx {
                minIdx = wIdx
                minValue = idx+1
            }
        }
        for idx, w:= range nums {
            wIdx := strings.Index(line, w)
            if wIdx == -1 {
                continue
            }
            // v += i 
            if minIdx == -1 || wIdx <= minIdx {
                minIdx = wIdx
                minValue = idx+1
            }
        }

        var maxIdx, maxValue int = -1, -1
        for idx, w:= range words {
            wIdx := strings.LastIndex(line, w)
            
            if wIdx == -1 {
                continue
            }
            // v += i 
            if maxIdx == -1 || wIdx >= maxIdx {
                maxIdx = wIdx
                maxValue = idx+1
            }
        }
        for idx, w:= range nums {
            wIdx := strings.LastIndex(line, w)
            
            if wIdx == -1 {
                continue
            }
            // v += i 
            if maxIdx == -1 || wIdx >= maxIdx {
                maxIdx = wIdx
                maxValue = idx+1
            }
        }

        sum += minValue*10 + maxValue
        fmt.Printf("%d\n", minValue*10 + maxValue)
    }
    fmt.Printf("SUM: %d\n", sum)
  
    readFile.Close()
}