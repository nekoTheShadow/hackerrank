package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the arrayManipulation function below.
func arrayManipulation(n int32, queries [][]int32) int64 {
	imos := make([]int64, n+1)
	for _, query := range queries {
		a := query[0] - 1
		b := query[1] - 1
		k := query[2]
		imos[a] += int64(k)
		imos[b+1] -= int64(k)
	}
	for i := 0; i < int(n); i++ {
		imos[i+1] += imos[i]
	}

	max := int64(-1)
	for i := 0; i < int(n); i++ {
		if max < imos[i] {
			max = imos[i]
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)
	stdout := os.Stdout

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	var queries [][]int32
	for i := 0; i < int(m); i++ {
		queriesRowTemp := strings.Split(readLine(reader), " ")

		var queriesRow []int32
		for _, queriesRowItem := range queriesRowTemp {
			queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
			checkError(err)
			queriesItem := int32(queriesItemTemp)
			queriesRow = append(queriesRow, queriesItem)
		}

		if len(queriesRow) != int(3) {
			panic("Bad input")
		}

		queries = append(queries, queriesRow)
	}

	result := arrayManipulation(n, queries)

	fmt.Fprintf(writer, "%d\n", result)

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
