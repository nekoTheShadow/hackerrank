package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Complete the highestValuePalindrome function below.
func highestValuePalindrome(s string, n int32, k int32) string {
	m := int(n)
	l := int(k)

	if m == 1 {
		return "9"
	}

	t := []byte(s)
	for i := 0; i < m/2; i++ {
		if t[i] == t[m-i-1] {
			continue
		}

		l--
		if t[i] < t[m-i-1] {
			t[i] = t[m-i-1]
		} else {
			t[m-i-1] = t[i]
		}
	}

	if l < 0 {
		return "-1"
	}

	for i := 0; i < m/2; i++ {
		if l == 0 {
			break
		}

		if t[i] == '9' {
			continue
		}

		if s[i] == s[m-i-1] {
			if l >= 2 {
				t[i] = '9'
				t[m-i-1] = '9'
				l -= 2
			}
		} else {
			t[i] = '9'
			t[m-i-1] = '9'
			l--
		}
	}

	if l > 0 && m%2 != 0 {
		t[m/2] = '9'
	}

	return string(t)
}

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Split(bufio.ScanWords)

	n, _ := strconv.Atoi(Scan(stdin))
	k, _ := strconv.Atoi(Scan(stdin))
	s := Scan(stdin)
	fmt.Println(highestValuePalindrome(s, int32(n), int32(k)))
}

func Scan(stdin *bufio.Scanner) string {
	stdin.Scan()
	return stdin.Text()
}
