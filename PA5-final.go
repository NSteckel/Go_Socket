package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":11111")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()

		defer conn.Close()

		reader := bufio.NewReader(conn)
		Size, sizeErr := reader.ReadString('\n')
		check(sizeErr)
		//fmt.Printf("%s", messageSize)

		messageSize := "" + Size
		line := strings.TrimSuffix(messageSize, "\n")
		fmt.Println("Upload file size: ", line)

		/*messageFull, messErr := reader.ReadString('\n')
		check(messErr)
		//fmt.Printf("%s", messageFull)*/

		//\\//\\//\\

		// create whatever
		f, err := os.Create("whatever.txt")
		check(err)
		defer f.Close()

		messageInt, convErr := strconv.Atoi(line)
		check(convErr)

		//fmt.Println("hello... %n", messageInt)
		//fmt.Println("Buuut")

		//fmt.Println(5 + messageInt)

		fileWriter := bufio.NewWriter(f)

		numBytes := 0
		numLines := 0

		for numBytes < messageInt {
			numLines = numLines + 1
			stringmessage, erf := reader.ReadString('\n')
			check(erf)

			numBytes = numBytes + len(stringmessage) + 2
			_, errwr := fileWriter.WriteString(fmt.Sprintf("%d %s", numLines, stringmessage))
			check(errwr)
			fileWriter.Flush()
			//fmt.Println("num bytes: ", numBytes)
		}

		numBytes = numBytes - 1
		//fileWriter.WriteString("hello")
		fmt.Println("Output file size: ", numBytes)

		//\\//\\//\\

		writer := bufio.NewWriter(conn)
		newline := fmt.Sprintf("%d byte file generated \n", numBytes)
		_, errw := writer.WriteString(newline)
		check(errw)
		writer.Flush()
	}
}
