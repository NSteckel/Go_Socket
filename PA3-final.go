package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	conn, errc := net.Dial("tcp", "3.143.172.18:11111")
	check(errc)
	defer conn.Close()

	fmt.Printf("Input filename? ")
	textInput := ""
	fmt.Scanf("%s", &textInput)

	f, err := os.Open(textInput)
	check(err)

	scanner := bufio.NewScanner(f)

	writer := bufio.NewWriter(conn)

	newString := ""

	for scanner.Scan() {
		newString += scanner.Text() + "\n"
	}
	newString = newString[:len(newString) - 1]
	//fmt.Printf(newString)

	StringSize := len(newString)
	fmt.Printf("Send the file size first: " + fmt.Sprint(StringSize) + "\n")

	_, errw := writer.WriteString(fmt.Sprint(StringSize) + "\n")
	check(errw)
	_, errw = writer.WriteString(newString)
	check(errw)

	writer.Flush()

	serverScanner := bufio.NewScanner(conn)
	if serverScanner.Scan() {
		fmt.Printf("Server replies: %s\n", serverScanner.Text())
	}

	f.Close()

}
