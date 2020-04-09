package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func exec(stdin *Stdin, stdout *Stdout) {
	n := stdin.ReadInt()
	r := stdin.ReadInt()
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, stdin.ReadInt())
	}

	a := make([]int, n)
	c := map[int]int{}
	for i, v := range arr {
		if v%r == 0 {
			a[i] = Get(c, v/r)
		}
		c[v] = Get(c, v) + 1
	}

	ans := 0
	d := map[int]int{}
	for i, v := range arr {
		if v%r == 0 {
			ans += Get(d, v/r)
		}
		d[v] = Get(d, v) + a[i]
	}

	stdout.Println(ans)
}

func Get(counter map[int]int, key int) int {
	if v, ok := counter[key]; ok {
		return v
	} else {
		return 0
	}
}

func main() {
	stdout := NewStdout()
	defer stdout.Flush()
	exec(NewStdin(bufio.ScanWords), stdout)
}

type Stdin struct {
	stdin *bufio.Scanner
}

func NewStdin(split bufio.SplitFunc) *Stdin {
	s := Stdin{bufio.NewScanner(os.Stdin)}
	s.stdin.Split(split)
	s.stdin.Buffer(make([]byte, bufio.MaxScanTokenSize), int(math.MaxInt32))
	return &s
}

func (s *Stdin) Read() string {
	s.stdin.Scan()
	return s.stdin.Text()
}

func (s *Stdin) ReadInt() int {
	n, _ := strconv.Atoi(s.Read())
	return n
}

func (s *Stdin) ReadFloat64() float64 {
	n, _ := strconv.ParseFloat(s.Read(), 64)
	return n
}

type Stdout struct {
	stdout *bufio.Writer
}

func NewStdout() *Stdout {
	return &Stdout{bufio.NewWriter(os.Stdout)}
}

func (s *Stdout) Flush() {
	s.stdout.Flush()
}

func (s *Stdout) Println(a ...interface{}) {
	fmt.Fprintln(s.stdout, a...)
}
