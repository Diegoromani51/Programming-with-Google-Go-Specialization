package main

import ( "fmt"
		 "log"
		 "os"
		 "sort"
		 "strconv"
)

const ( ExitCode string = "X")

func main() {
		var input string
		var initslice = make([]int, 3)
		fmt.Printf("Initialized slice with Len: %v \n", len(initslice))
		for {
				fmt.Println("\t-----------------------------------")
				fmt.Printf("Current slice: %v\n", initslice)
				fmt.Printf("Enter a number to add to slice ('X' to exit program):")
				if c, err := fmt.Scanf("%s", &input); err != nil {
					if c ==0 {
						log.Println("[ERROR] Empty Input")
						continue
					}
					log.Println(err)
					break
				}
				if input == ExitCode {
					fmt.Println("Exiting program")
					os.Exit(0)

				}
				intValue, err :=strconv.Atoi(input)
				if err != nil {
					fmt.Printf( "[ERROR] Invalid Input: ->'%s'\n", input)
					continue

				}
				i, found := Find(initslice, 0)
				if !found {
					initslice = append(initslice, intValue)
				} else {
					initslice[i] = intValue

				}
				sort.Ints(initslice)
				fmt.Println(initslice)

		}
}

func Find(slice []int, val int) (int,bool) {
	for i, item := range slice {
		if item == val {
			return i, true

		}
	}
	return -1, false
}
