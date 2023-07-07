package main

import "fmt"
// we can import external package like this
import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
