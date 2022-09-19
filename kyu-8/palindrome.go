/*Write a function that checks if a given string (case insensitive) is a palindrome.*/

package kata

import "strings"

func IsPalindrome(str string) bool {

	str = strings.ToUpper(str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
