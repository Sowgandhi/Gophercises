package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)


func caesarCipher(s string, n int32, k int32) string {
	runes := []rune(s)
	k = k % 26
	var i int32
	for i = 0; i < n; i++ {
		if runes[i] >= 'a' && runes[i] <= 'z' {
			runes[i] = rune('a' + ((int32(runes[i])+k)%'a')%26)
		} else if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = rune('A' + ((int32(runes[i])+k)%'A')%26)
		} else {
			continue
		}

	}

	return string(runes)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(s, n, k)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
