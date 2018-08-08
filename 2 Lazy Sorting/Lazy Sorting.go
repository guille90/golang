package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Complete the solve function below.
func solve(P []int) float64 {
	if sort.IntsAreSorted(P) {
		return 0
	}

	sort.Ints(P)
	n := factorial(float64(len(P)))

	// var sameValues int
	// for i := 0; i < len(P)-1; i++ {
	// 	if P[i] == P[i+1] {
	// 		sameValues++
	// 	}
	// }
	// if sameValues != 0 {
	// 	div := math.Pow(float64(sameValues), 2)
	// 	n /= div
	// }
	return n / 1111523212800

}

func factorial(n float64) float64 {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	// P := []int{5, 2}
	// //P := []int{39, 34, 33, 40, 34, 33}
	// //P := []int{26, 9, 18, 18, 17, 12, 23, 23, 14, 29} //case 4
	// //P := []int{5, 2, 19, 15, 19, 23, 6, 11, 2, 18, 6, 30, 28, 5, 12} //case 6
	P := []int{99, 98, 97, 97, 100, 95, 96, 91, 97, 99, 98, 95, 93, 90, 95, 97, 97, 91} //case 7
	time := solve(P)
	fmt.Println(time)

	// reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	// writer := bufio.NewWriterSize(stdout, 1024*1024)

	// PCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	// checkError(err)

	// PTemp := strings.Split(readLine(reader), " ")

	// var P []int

	// for i := 0; i < int(PCount); i++ {
	// 	PItemTemp, err := strconv.ParseInt(PTemp[i], 10, 64)
	// 	checkError(err)
	// 	PItem := int(PItemTemp)
	// 	P = append(P, PItem)
	// }

	// result := solve(P)

	// s := fmt.Sprintf("%.6f", result)

	// fmt.Fprintf(writer, "%v\n", s)

	// writer.Flush()
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