package main

import( "fmt"
        "log"
)

var inputNumber float64

func main(){
		fmt.Println("Input a Float Number:")
		_, err :=
		fmt.Scan(&inputNumber)
		if err != nil{
			log.Printf("[Error] Invalid Input")
		}

		fmt.Printf("This is the truncated version of your input '%v'.\n", int64(inputNumber))
}
