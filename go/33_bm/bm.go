package main

import (
	"fmt"
	"math"
)

const SIZE = 126

func GenerateBC(b string, bc []int) {

	m := len(b)

	for i := 0; i < SIZE; i++ {
		bc[i] = -1
	}

	for i := 0; i < m; i++ {
		ascii := int(b[i])
		bc[ascii] = i
	}
}

func BM(a, b string) int {

	bc := make([]int, SIZE, SIZE)
	GenerateBC(b, bc)

	n := len(a)
	m := len(b)

	suffix := make([]int, m)
	prefix := make([]bool, m)
	GenerateGS(b, suffix, prefix)

	i := 0 //模式串滑动的位数
	for i <= n-m {

		j := 0 //坏字符的位置
		for j = m - 1; j >= 0; j-- {
			if a[i+j] != b[j] {
				break
			}
		}

		if j == -1 {
			return i
		}

		x := j - bc[a[i+j]]

		y := 0
		if j < m-1 {
			y = MoveByGS(j, m, suffix, prefix)
		}

		i = i + int(math.Max(float64(x), float64(y)))
	}

	return -1
}

func GenerateGS(b string, suffix []int, prefix []bool) {

	m := len(b)

	for i := 0; i < m; i++ {
		suffix[i] = -1
		prefix[i] = false
	}

	for i := 0; i < m-1; i++ {
		j := i
		k := 0
		for j >= 0 && b[j] == b[m-1-k] {
			j--
			k++
			suffix[k] = j + 1
		}

		if j == -1 {
			prefix[k] = true
		}
	}
}

func MoveByGS(j, m int, suffix []int, prefix []bool) int {

	k := m - j - 1 //好后缀长度

	if suffix[k] != -1 {
		return j - suffix[k] + 1
	}

	for r := j + 2; r <= m-1; r++ {
		if prefix[m-r] {
			return r
		}
	}

	return m
}

func main() {
	s := "abc abcdab abcdabcdabde"
	pattern := "bcdabd"
	fmt.Println(BM(s, pattern)) //16

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "abab"
	fmt.Println(BM(s, pattern)) //11

	s = "aabbbbaaabbababbabbbabaaabb"
	pattern = "ababacd"
	fmt.Println(BM(s, pattern)) //-1

	s = "hello"
	pattern = "ll"
	fmt.Println(BM(s, pattern)) //2

	s = "abcacabcbcbacabc"
	pattern = "cbacabc"
	fmt.Println(BM(s, pattern)) //9

	s = "abcacabcbcbacabc"
	pattern = "cabcab"
	fmt.Println(BM(s, pattern)) //-1
}
