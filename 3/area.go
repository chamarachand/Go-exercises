package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Length: ")
	lengthInput, err := reader.ReadString('\n')

	lengthInput = strings.TrimSpace(lengthInput)
	length, err := strconv.ParseInt(lengthInput, 10, 64)
	if err != nil {
		fmt.Println("Invalid length")
		return
	}

	fmt.Print("Width: ")
	widthInput, err := reader.ReadString('\n')

	widthInput = strings.TrimSpace(widthInput)
	width, err := strconv.ParseInt(widthInput, 10, 64)
	if err != nil {
		fmt.Println("Invalid length")
		return
	}

	area := length * width
	fmt.Print("The area of the rectangle is: ", area)
}
