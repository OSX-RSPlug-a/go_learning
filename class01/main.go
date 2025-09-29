package main

import (
	"fmt"
	"errors"
)


func checkNumber(x int) (string, error) {
    if x < 20 {
        return "", errors.New("number is less than 20") 
    }
    
    return "number is less than 20 or greater", nil
}

func main() { 

    x:= 19
    y:= "Strings ex" 
    z:= true

    result, err := checkNumber(x)

    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println(result, x)
    }

    fmt.Println(y, z) 
}