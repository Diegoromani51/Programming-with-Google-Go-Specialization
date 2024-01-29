package main

import ("fmt"
		"sync"
)

const UintSize = 32 << (^uint(0) >> 32 & 1)     /* Number of bits in an 'int' variable*/
const MaxInt = 1<<(UintSize-1) - 1              /* Max possible value of an 'int' varaible*/

func BubbleSort(arr []int, low, high int, wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done()

	/* Bubble Sort algorithm to sort the sub-arrays from low to high */
	for i := low; i < high; i++ {
		for j := i + 1; j < high; j++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i] 
			}
		}
	}

	/* Printing sorted subarrays, using mutex (mu) to make sure other goroutines do not print at the same time */
	mu.Lock()
	fmt.Printf("Sorted by gorountine #%d:", (low/(len(arr)/4))+1)
	for i := low; i < high; i++ {
		fmt.Printf("%d", arr[i])} 
	fmt.Print("\n")
	mu.Unlock()

}

/* This function is used to identify which number is minimum among the four numbers */
func Minof4(a, b, c, d int) int {
	if a < b && a < c && a < d {
		return 0
	} else if b < c && b < d {
		return 1
	} else if c < d {
		return 2
	}
	return 3
}


/* Take array as input from the user */ 
func Inputfunc() []int {
	var size, num int
	arr := make([]int, 0)

	fmt.Println("Enter size of input array:")
	fmt.Scan(&size)
	fmt.Println("Enter", size, " integers:")
	for i := 0; i < size; i++ {
		fmt.Scan(&num)
		arr = append(arr, num)
	}
	return arr
}

func main() {
	arr := Inputfunc() /* Take array user input */
    
	/* Calculate the lower and upper bounds for the 4 partitions of the array */
	var lo, hi [4]int
	for i := 0; i < 4; i++ {
		lo[i] = (len(arr)/4) * i
		hi[i] = (len(arr)/4) * (i + 1)
	}
	hi[3] = len(arr)

	var wg sync.WaitGroup /* Make the main() wait for the the other gorountines to complete*/
	var mu sync.Mutex     /* Makes one goroutine print at a time*/
	
	wg.Add(4)

	go BubbleSort(arr, lo[0], hi[0], &wg, &mu)
	go BubbleSort(arr, lo[1], hi[1], &wg, &mu)
	go BubbleSort(arr, lo[2], hi[2], &wg, &mu)
	go BubbleSort(arr, lo[3], hi[3], &wg, &mu)

	wg.Wait()

	aux := make([]int, 0, len(arr))  /* Empty array used for merging the sorted array*/

	/* Merge the arrays and store the result inot the empty array aux*/
	for i := 0; i < len(arr); i++ {
		pos := Minof4(arr[lo[0]], arr[lo[1]], arr[lo[2]], arr[lo[3]])
		aux = append(aux, arr[lo[pos]])
		lo[pos] = lo[pos] + 1

		if lo[pos] == hi[pos] {
			lo[pos]--
			arr[lo[pos]] = MaxInt
		}
	}

	fmt.Print("Merged Array:")
	fmt.Print(aux)

}
