package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

//闭包

func add() func(int) int {
	sum := 0
	return func(v int) int {
		sum = sum + v
		return sum
	}
}

type intGen func() int

//斐波那契
func fabonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, b+a
		return a
	}
}
func (ig intGen) Read(p []byte) (n int, err error) {
	next := ig()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printReaderContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
func main() {
	adder := add()
	s := adder(5)
	fmt.Println("s:", s)

	f := fabonacci()
	printReaderContents(f)

}
