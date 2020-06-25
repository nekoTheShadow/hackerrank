package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

// Complete the maxCircle function below.
func maxCircle(queries [][]int32) []int32 {
	n := len(queries)
	uf := NewUnionFind(1e9)
	ans := make([]int32, n)
	max := -INFINITY
	for i := 0; i < n; i++ {
		a := int(queries[i][0] - 1)
		b := int(queries[i][1] - 1)
		uf.Union(a, b)
		max = Max(max, uf.Size(a))
		ans[i] = int32(max)
	}
	return ans
}

const INFINITY = math.MaxInt64/2 - 1

type UnionFind struct {
	parents map[int]int
	sizes   map[int]int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		parents: map[int]int{},
		sizes:   map[int]int{},
	}
	return &uf
}

func (uf *UnionFind) Find(x int) int {
	if _, ok := uf.parents[x]; !ok {
		return x
	}
	uf.parents[x] = uf.Find(uf.parents[x])
	return uf.parents[x]
}

func (uf *UnionFind) Union(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if _, ok := uf.sizes[x]; !ok {
		uf.sizes[x] = 1
	}
	if _, ok := uf.sizes[y]; !ok {
		uf.sizes[y] = 1
	}

	if uf.sizes[x] < uf.sizes[y] {
		uf.parents[x] = y
		uf.sizes[y] += uf.sizes[x]
	} else {
		uf.parents[y] = x
		uf.sizes[x] += uf.sizes[y]
	}
}

func (uf *UnionFind) Size(x int) int {
	x = uf.Find(x)
	if _, ok := uf.sizes[x]; !ok {
		uf.sizes[x] = 1
	}
	return uf.sizes[x]
}

func Pow(x, y int) int {
	z := 1
	for y > 0 {
		if y%2 == 0 {
			x *= x
			y /= 2
		} else {
			z *= x
			y -= 1
		}
	}
	return z
}

func Max(a int, b ...int) int {
	for _, v := range b {
		if a < v {
			a = v
		}
	}
	return a
}

func exec(stdin *Stdin, stdout *Stdout) {
	n := stdin.ReadInt()
	queries := make([][]int32, n)
	for i := 0; i < n; i++ {
		a := int32(stdin.ReadInt())
		b := int32(stdin.ReadInt())
		queries[i] = []int32{a, b}
	}
	for _, ans := range maxCircle(queries) {
		stdout.Println(ans)
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
	s.stdin.Buffer(make([]byte, bufio.MaxScanTokenSize), INFINITY)
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

func (s *Stdin) ReadIntSlice(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = s.ReadInt()
	}
	return a
}

func (s *Stdin) ReadStringSlice(n int) []string {
	a := make([]string, n)
	for i := 0; i < n; i++ {
		a[i] = s.Read()
	}
	return a
}

func (s *Stdin) ReadFloat64Slice(n int) []float64 {
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = s.ReadFloat64()
	}
	return a
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
