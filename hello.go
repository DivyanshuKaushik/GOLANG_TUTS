package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const PI float32 = 3.14

// add to numbers
func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello, world")
	reader  := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')
	fmt.Printf("value = %s",input)
	fmt.Println(add(7, 11))
	value := time.Now()
	fmt.Println(value.Format("01-02-2006"))
	fmt.Printf("value of pi is %.2f \n", PI)
}
