package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	n := len(s)
	d := map[string]int{}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			t := Sort(s[i : j+1])
			if _, ok := d[t]; ok {
				d[t]++
			} else {
				d[t] = 1
			}
		}
	}

	c := 0
	for _, v := range d {
		c += v * (v - 1) / 2
	}
	return int32(c)
}

func Sort(s string) string {
	t := strings.Split(s, "")
	sort.Strings(t)
	return strings.Join(t, "")
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Split(bufio.ScanWords)

	q, _ := strconv.Atoi(Scan(stdin))
	for i := 0; i < q; i++ {
		fmt.Println(sherlockAndAnagrams(Scan(stdin)))
	}
}

func Scan(stdin *bufio.Scanner) string {
	stdin.Scan()
	return stdin.Text()
}
