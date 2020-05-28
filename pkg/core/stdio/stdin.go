package main

import (
	"fmt"
	"os"
//	"io/ioutil"
	"bufio"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter Age:")

	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	dogyears, _ := strconv.Atoi(age)

	fmt.Print("the age: ", age, " is ", dogyears * 7, " in dog years"  )
}
