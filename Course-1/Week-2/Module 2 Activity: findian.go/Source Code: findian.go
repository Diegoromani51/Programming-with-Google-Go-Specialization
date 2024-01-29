package main

import ( "fmt"
		 "strings"
		 "os"
		 "bufio"
)

func main() {

	var inputString string
	fmt.Println("Please Input a String")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputString = scanner.Text()
		break
	}


	inputString = strings.ToLower(inputString)

	if strings.HasPrefix(inputString, "i") &&
		strings.HasSuffix(inputString, "n") &&
		strings.Contains(inputString, "a") {

			fmt.Printf("Found!\n") 
			fmt.Printf("Your string starts with (i), ends with (n), and has an (a) in it")
	} else {
		fmt.Printf("Not Found!\n")
		fmt.Printf("Your string does no start with (i), end with (n) and have an (a) in it")
	}



}


