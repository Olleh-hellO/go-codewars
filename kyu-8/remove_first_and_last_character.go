/*It's pretty straightforward. Your goal is to create a function that removes the first and last characters of a string. You're given one parameter, the original string. You don't have to worry with strings with less than two characters.*/

package kata

func RemoveChar(word string) string {
	buffer := ""
	for i := 0; i < len(word); i++ {
		if i > 0 && i < len(word)-1 {
			buffer += string(word[i])
		}
	}
	return buffer
}
