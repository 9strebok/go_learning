package recursive_fact

import (
    "fmt"
    "testing"
)


func TestRecursiveFact(t *testing.T) {
    fmt.Println("Start testing")
    if recursive_fact(2) != 1*2 {
        t.Fail()
    }
    if recursive_fact(0) != 1 {
        t.Fail()
    }
    if recursive_fact(10) != 1*2*3*4*5*6*7*8*9*10 {
        t.Fail()
    }
}
