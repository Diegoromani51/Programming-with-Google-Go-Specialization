package main

import ( "fmt"
		 "bufio"
		 "io/ioutil"
		 "strings")

type Name struct {
	fname string
	lname string 
}

var People []Name

func retrievedata(filename string) {
	var count int64=0

	barr, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error unable to read file %s:", err)
	}

	fmt.Printf("Reading file provided...")
	scanner := bufio.NewScanner(strings.NewReader(string(barr)))
	for scanner.Scan() {
		var name Name
		line := strings.Split(scanner.Text(), " ")
		name.fname = line[0]
		name.lname = line[1]

		People = append(People, name)
		count += 1

	}
	fmt.Println("Total lines read: ", count)
	fmt.Println("-----------------------------------------------")

}

func main() {
	var filename string

	fmt.Println("Enter the name of the file you would like read:")
	_, err := fmt.Scanf("%s", &filename)
	if err != nil {
		fmt.Println("Error has occured: %s", err)
	}

	retrievedata(filename)
	fmt.Println("------------------------------------------")

	for _, name := range People {
		fmt.Println("First_Name:", name.fname, "\tLast_Name:", name.lname)
	}
}
