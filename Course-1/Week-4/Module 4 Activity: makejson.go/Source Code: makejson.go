package main

import ( "fmt"
		 "bufio"
		 "os"
		 "encoding/json")

// Function to take user input
func userInput(message string) string {
	fmt.Println(message)
	
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}

func main() {
	x := make(map[string]string)

	fmt.Println("-------------------------------")
	x["name"] = userInput("Enter a Name:")
	x["address"] = userInput("Enter their address:")

	data, err := json.Marshal(x)
	if err != nil {
		fmt.Printf("Error while converting to JSON format: %s", err)
		os.Exit(0)
	}
	fmt.Println("-------------------------------")
	fmt.Printf("JSON object: %v", string(data))
}
