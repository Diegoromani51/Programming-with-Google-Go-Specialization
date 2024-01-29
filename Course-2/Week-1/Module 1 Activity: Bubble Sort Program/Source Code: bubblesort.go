package main

import ( "bufio"
		 "fmt"
		 "os"
		 "strconv"
		 "strings")

func BubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice)-1-i; j++ {
			if slice[j+1] < slice[j] {
				SwapNumber(slice, j)
			}
		} 
	}
}

func SwapNumber(sliceb []int, j int) {
	var swap int
	swap = sliceb[j]
	sliceb[j] = sliceb[j+1]
	sliceb[j+1] = swap
}

func main() {
	fmt.Println("Please input a sequence of up to 10 integers, separated by a space:")
	scanner := bufio.NewReader(os.Stdin)
	input, _, _ := scanner.ReadLine()
	newString := strings.Split(string(input), " ")
	var values []int
	for _, s := range newString {
		n, _ := strconv.Atoi(s)
		values = append(values, n)
	}
	BubbleSort(values)
	fmt.Println(values)

}
