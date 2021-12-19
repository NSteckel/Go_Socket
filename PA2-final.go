// fill with PA2.go:

// package main

package main // ALWAYS THE FIRST LINE!!!

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("Enter an Input File: \n")
	textInput := "" // asks for value, and assigns value to 'text' at the same time
	fmt.Scanf("%s", &textInput)

	// Read Section
	f, err := os.Open(textInput)
	check(err)

	// scanner := bufio.NewScanner(f)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }
	// f.Close()
	////
	fmt.Printf("Enter an Output File: \n")
	textOutput := ""
	fmt.Scanf("%s", &textOutput)

	// fmt.Fprintf(os.Stdout, "Input: %s\n", textInput)
	// fmt.Fprintf(os.Stdout, "Output: %s\n", textOutput)
	//  ^All three print the same thing

	// Write section
	outputfile, err := os.Create(textOutput)
	check(err)
	defer outputfile.Close()

	scanner := bufio.NewScanner(f)
	writer := bufio.NewWriter(outputfile)
	i := 1
	for scanner.Scan() {
		newString := fmt.Sprint(i) + " " + scanner.Text() + "\n"
		writer.WriteString(newString)
		i++
	}

	f.Close()
	writer.Flush()
}
