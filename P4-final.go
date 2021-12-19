package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":8888")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	message := ""

	if scanner.Scan() {
		message = scanner.Text()
		fmt.Println("Upload file size: " + message + "\n")
	}

	f, err := os.Create("whatever.txt")
	check(err)
	defer f.Close()

	/// /\_/\ \\\

	fileWriter := bufio.NewWriter(f)
	fmt.Println("writer created")
	i := 1
	bytesize := 1
	newEye := 0
	eye := ""
	messageInt, convErr := strconv.Atoi(message)
	check(convErr)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		bytesize = len(scanner.Bytes()) + bytesize

		newString := fmt.Sprint(i) + " " + scanner.Text() + "\n"
		fileWriter.WriteString(newString)

		eye = strconv.Itoa(i)
		eyeLen := len(eye)
		newEye = newEye + eyeLen
		//fmt.Println("New Eye : ", newEye)
		i++
		// length of i
		//convert back to int add to bytesize

		//fmt.Println("Byte Size %n", bytesize)
		//fmt.Println("i incremented to: %n", i)
		fmt.Println("Size for comp: %n", messageInt)

		//fmt.Println("Final Size: %n", finSize)

		if bytesize == messageInt {
			break
		}
	}
	// close opened input file
	fileWriter.Flush()
	//fmt.Println("loop closed, file closed.")
	finSize := bytesize + newEye + i - 1
	fmt.Println("Output file size: %n", finSize)

	//outScan := bufio
	// calculate the length of i (prepended to each string and add that to the original size)

	/*length, _ := fileWriter.WriteString("This is a test!")
	fmt.Println(length)
	fileWriter.Flush()*/

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}
