package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"strconv"
)

// Complete the countLuck function below.
func countLuck(matrix []string, k int32) string {
	n := len(matrix)
	m := len(matrix[0])
	count := make([][]int, n)
	prev := make([][]*Tuple, n)
	stack := NewDeque()
	goalX := 0
	goalY := 0
	for i := 0; i < n; i++ {
		count[i] = make([]int, m)
		prev[i] = make([]*Tuple, m)
		for j := 0; j < m; j++ {
			count[i][j] = INFINITY
			if matrix[i][j] == 'M' {
				count[i][j] = 0
				prev[i][j] = &Tuple{X: i, Y: j}
				stack.AppendLast(&Tuple{X: i, Y: j})
			} else if matrix[i][j] == '*' {
				goalX = i
				goalY = j
			}
		}
	}

	diffs := []*Tuple{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
	}
	for !stack.IsEmpty() {
		t := stack.PopLast().(*Tuple)
		for _, d := range diffs {
			x := t.X + d.X
			y := t.Y + d.Y
			if 0 <= x && x < n &&
				0 <= y && y < m &&
				(matrix[x][y] == '.' || matrix[x][y] == '*') &&
				count[t.X][t.Y]+1 < count[x][y] {
				count[x][y] = count[t.X][t.Y] + 1
				stack.AppendLast(&Tuple{X: x, Y: y})
				prev[x][y] = t
			}
		}
	}

	route := NewDeque()
	route.AppendLast(&Tuple{X: goalX, Y: goalY})
	for {
		a := route.PeekLast().(*Tuple)
		b := prev[a.X][a.Y]
		if a.X == b.X && a.Y == b.Y {
			break
		}
		route.AppendLast(b)
	}
	route.PopFirst()

	ans := 0
	for !route.IsEmpty() {
		t := route.PopLast().(*Tuple)
		u := prev[t.X][t.Y]
		c := 0
		for _, d := range diffs {
			x := t.X + d.X
			y := t.Y + d.Y
			if 0 <= x && x < n &&
				0 <= y && y < m &&
				!(u.X == x && u.Y == y) &&
				(matrix[x][y] == '.' || matrix[x][y] == '*') {
				c++
			}
		}

		if c >= 2 {
			ans++
		}
	}

	if ans == int(k) {
		return "Impressed"
	} else {
		return "Oops!"
	}
}

type Tuple struct {
	X int
	Y int
}

type Deque struct {
	head []interface{}
	tail []interface{}
}

func NewDeque() *Deque {
	return &Deque{
		head: []interface{}{},
		tail: []interface{}{},
	}
}

func (d *Deque) AppendLast(v interface{}) {
	d.tail = append(d.tail, v)
}

func (d *Deque) AppendFirst(v interface{}) {
	d.head = append(d.head, v)
}

func (d *Deque) PopLast() interface{} {
	if len(d.tail) == 0 {
		v := d.head[0]
		d.head = d.head[1:]
		return v
	} else {
		v := d.tail[len(d.tail)-1]
		d.tail = d.tail[:len(d.tail)-1]
		return v
	}
}

func (d *Deque) PopFirst() interface{} {
	if len(d.head) == 0 {
		v := d.tail[0]
		d.tail = d.tail[1:]
		return v
	} else {
		v := d.head[len(d.head)-1]
		d.head = d.head[:len(d.head)-1]
		return v
	}
}

func (d *Deque) IsEmpty() bool {
	return len(d.head) == 0 && len(d.tail) == 0
}

func (d *Deque) PeekFirst() interface{} {
	if len(d.head) == 0 {
		return d.tail[0]
	} else {
		return d.head[len(d.head)-1]
	}
}

func (d *Deque) PeekLast() interface{} {
	if len(d.tail) == 0 {
		return d.head[0]
	} else {
		return d.tail[len(d.tail)-1]
	}
}

const INFINITY = math.MaxInt64/2 - 1

func exec(stdin *Stdin, stdout *Stdout) {
	for t := stdin.ReadInt(); t > 0; t-- {
		n := stdin.ReadInt()
		m := stdin.ReadInt()
		Unuse(m)
		matrix := []string{}
		for i := 0; i < n; i++ {
			matrix = append(matrix, stdin.Read())
		}
		k := stdin.ReadInt()
		stdout.Println(countLuck(matrix, int32(k)))
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

func Min(a int, b ...int) int {
	for _, v := range b {
		if v < a {
			a = v
		}
	}
	return a
}

func Max(a int, b ...int) int {
	for _, v := range b {
		if a < v {
			a = v
		}
	}
	return a
}

func Abs(x int) int {
	if x > 0 {
		return x
	} else {
		return x * -1
	}
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

func IsPrime(x int) bool {
	if x < 2 {
		return false
	}

	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}

	for y > 0 {
		x, y = y, x%y
	}

	return x
}

func Lcm(x, y int) int {
	return x * y / Gcd(x, y)
}

func CeilDiv(x, y int) int {
	if x%y == 0 {
		return x / y
	} else {
		return x/y + 1
	}
}

func LCS(s, t string) string {
	n := len(s)
	m := len(t)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	x := n
	y := m
	u := make([]byte, dp[n][m])
	i := dp[n][m] - 1
	for x > 0 && y > 0 {
		if dp[x][y] == dp[x-1][y] {
			x--
		} else if dp[x][y] == dp[x][y-1] {
			y--
		} else {
			x--
			y--
			u[i] = s[x]
			i--
		}
	}

	return string(u)
}

// [0..n) から r個選んだ組合せを生成する。
// r < 0 の場合はすべての組合せを生成する。
func Combinations(n int, r int) [][]int {
	a := [][]int{}
	for i := uint(0); i < (1 << uint(n)); i++ {
		b := []int{}
		for j := uint(0); j < uint(n); j++ {
			if i&(uint(1)<<j) != 0 {
				b = append(b, int(j))
			}
		}

		if r < 0 || len(b) == r {
			a = append(a, b)
		}
	}
	return a
}

// [0..n)の順列を生成する。
func Permutations(n int) [][]int {
	type e struct {
		bit uint
		a   []int
	}

	q := []*e{}
	for i := 0; i < n; i++ {
		q = append(q, &e{bit: 1 << uint(i), a: []int{i}})
	}

	p := [][]int{}
	for x := 0; x < len(q); x++ {
		cur := q[x]
		if cur.bit == (1<<uint(n) - 1) {
			p = append(p, cur.a)
			continue
		}

		for i := 0; i < n; i++ {
			if cur.bit&(1<<uint(i)) != 0 {
				continue
			}

			m := len(cur.a) + 1
			b := make([]int, m)
			copy(b, cur.a)
			b[m-1] = i
			q = append(q, &e{bit: cur.bit | (1 << uint(i)), a: b})
		}
	}
	return p
}

func ReverseSlice(v interface{}) {
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		slice := reflect.ValueOf(v)
		for i1, n := 0, slice.Len(); i1 < n/2; i1++ {
			i2 := n - i1 - 1
			v1 := reflect.ValueOf(slice.Index(i1).Interface())
			slice.Index(i1).Set(slice.Index(i2))
			slice.Index(i2).Set(v1)
		}
	}
}

func ReverseString(s string) string {
	t := []rune(s)
	ReverseSlice(t)
	return string(t)
}

func ToStringSlice(slice interface{}) []string {
	// 本来はerrを返却すべき。 利用する側(=自分)は引数がsliceであることを知っている。
	if reflect.TypeOf(slice).Kind() != reflect.Slice {
		panic(fmt.Sprintf("%v is not slice", slice))
	}

	src := reflect.ValueOf(slice)
	n := src.Len()
	dst := make([]string, n)
	for i := 0; i < n; i++ {
		dst[i] = fmt.Sprintf("%v", src.Index(i))
	}
	return dst
}

func Unuse(a ...interface{}) {
	// PASS
}
