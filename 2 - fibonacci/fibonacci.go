package main

import (
	"fmt"
	"os"
	"strconv"
)

func processArgs() int {
    args := os.Args
    if (len(args) > 1) {
        num, err := strconv.Atoi(args[1]);
        if err != nil {
            fmt.Println("Error - Can't convert input to an int")
            os.Exit(1)
        } else {
            return num;
        }
    }
    fmt.Println("Error - Expected an int input")
    os.Exit(1)
    return 0;
}

func fibonacci(n int) int {
    if (n == 0) {
        return 0;
    }

    var sign int = 1
    if n < 0 {
        n = n * -1
        if n % 2 == 0 {
            sign = -1
        }
    }

    var a, b int = 1, 1
    for i := 1; i < n; i++ {
        a, b = b, a+b
    }
    return a * sign;
}

func main() {
    fmt.Println(fibonacci(processArgs()));
}