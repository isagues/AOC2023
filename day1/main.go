package main

import (
    "fmt"
    "os"
    "bufio"
    "unicode"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

// one, two, three, four, five, six, seven, eight, nine

func main() {
    readFile, err := os.Open("input")
    check(err)
    
    fileScanner := bufio.NewScanner(readFile)
 
    fileScanner.Split(bufio.ScanLines)
  
    var sum = 0

    for fileScanner.Scan() {
        var f int = -1
        var l int = -1
        var n int = -1
        

        for _, c := range (fileScanner.Text()) {
            // fmt.Printf("%s :", fileScanner.Text())
            if unicode.IsDigit(c) {
                n = int(c - '0')
                if f == -1 {
                    f = n
                }
                l = n

            }
        }
        fmt.Printf("%d%d \n", f, l)
        sum += f*10 + l
    }
    fmt.Printf("SUM: %d\n", sum)
  
    readFile.Close()
}