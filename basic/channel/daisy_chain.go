package channel

import "fmt"

// daisy chain
 func daisy(left, right chan int) {
     left <- 1 + <-right
 }

 func chain() {
     const n = 10000
     leftmost := make(chan int)
     right := leftmost
     left := leftmost
     for i := 0; i < n; i++ {
         right = make(chan int)
         go daisy(left, right)
         left = right
     }
     go func(c chan int) { c <- 1 }(right)
     fmt.Println(<-leftmost)
 }
