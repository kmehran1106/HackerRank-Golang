package staircase

import (
    "strings"
)


func Staircase(n int) string {
    var s string = ""
    for i := 0; i < n; i++ {
        es := strings.Repeat(" ", (n-1-i))
        hv := strings.Repeat("#", (i+1))
        s += (es + hv)
        if i < (n-1) {
            s += "\n"
        }
    }
    // s = strings.TrimSuffix(s, "\n")
	return s
}

// package main
// 
// import (
//     "strings"
//     "fmt"
//     "bufio"
//     "os"
//     "strconv"
// )
// func main () {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter First Number : ")
// 	s, _ := reader.ReadString('\n')
// 	s = strings.TrimSpace(s)
// 	n, _ := strconv.Atoi(s)
// 	fmt.Println(Staircase(n))
// }