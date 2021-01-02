package number

import (
	"strconv"
	"strings"
)

/*
 * BigNumPlus
 * 1.从低位开始相加
 * 2.进位重置的时机
 * 3.数字串边界问题
 */
func BigNumPlus(n1 string, n2 string) string {
	var res string
	var carry, u1, u2, sum uint8
	var maxLen, idx1, idx2 int

	// ascii code:
	// +=>43
	// -=>45
	// 0=>48
	// 9=>57

	l1 := len(n1)
	l2 := len(n2)

	if l1 > l2 {
		maxLen = l1
	} else {
		maxLen = l2
	}

	byteN1 := []byte(n1)
	byteN2 := []byte(n2)

	for i := 1; i <= maxLen; i++ {

		if idx1 = l1 - i; idx1 >= 0 && (byteN1[idx1] >= 48 && byteN1[idx1] <= 57) {
			u1 = byteN1[idx1] - 48
		}

		if idx2 = l2 - i; idx2 >= 0 && (byteN2[idx2] >= 48 && byteN2[idx2] <= 57) {
			u2 = byteN2[idx2] - 48
		}

		sum = u1 + u2 + carry
		carry = 0
		if sum >= 10 {
			carry = 1
			sum -= 10
		}
		res = strconv.Itoa(int(sum)) + res

		u1 = 0
		u2 = 0
	}

	return res
}

/*
 * BigNumMinus
 * @param string n1: It is the bigger one.
 * @param string n2: It is the less one.
 * 1.从低位开始相加
 * 2.进位重置的时机
 * 3.借位重置的时机
 * 4.数字串边界问题
 */
func BigNumMinus(n1 string, n2 string) string {
	var res string
	var carry, u1, u2, sum, borrowing uint8
	var maxLen, idx1, idx2 int

	// ascii code:
	// +=>43
	// -=>45
	// 0=>48
	// 9=>57

	l1 := len(n1)
	l2 := len(n2)

	if l1 > l2 {
		maxLen = l1
	} else {
		maxLen = l2
	}

	byteN1 := []byte(n1)
	byteN2 := []byte(n2)

	for i := 1; i <= maxLen; i++ {

		if idx1 = l1 - i; idx1 >= 0 && (byteN1[idx1] >= 48 && byteN1[idx1] <= 57) {
			u1 = byteN1[idx1] - 48
		}

		if idx2 = l2 - i; idx2 >= 0 && (byteN2[idx2] >= 48 && byteN2[idx2] <= 57) {
			u2 = byteN2[idx2] - 48
		}

		if u1 < u2 {
			borrowing = 10
		}

		sum = u1 - carry - u2 + borrowing

		borrowing = 0
		carry = 0

		if u1 < u2 {
			carry = 1
		}

		res = strconv.Itoa(int(sum)) + res

		u1 = 0
		u2 = 0
	}

	res = strings.TrimLeft(res, "0")

	if res == "" {
		res = "0"
	}

	return res
}

/*
 * BigNumCompare: if n1 is equal or greater than n2, return bool true.
 * @param string n1
 * @param string n2
 */
func BigNumCompare(n1 string, n2 string) bool {
	// ascii code:
	// +=>43
	// -=>45
	// 0=>48
	// 9=>57

	l1 := len(n1)
	l2 := len(n2)

	if l1 > l2 {
		return true
	} else if l1 < l2 {
		return false
	} else {
		byteN1 := []byte(n1)
		byteN2 := []byte(n2)

		for i := 0; i < l1; i++ {
			if byteN1[i] > byteN2[i] {
				return true
			} else if byteN1[i] < byteN2[i] {
				return false
			}
		}
	}

	return true
}
