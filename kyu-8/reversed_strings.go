/*
Complete the solution so that it reverses the string passed into it.
```
'world'  =>  'dlrow'
'word' => 'drow'
```
*/

package kata

import "fmt"

func Solution(word string) string {

	reversedWord := ""

	for _, char := range word {
		reversedWord = string(char) + reversedWord
	}

	return reversedWord
}

func main() {
	fmt.Println(Solution(""))      // ""
	fmt.Println(Solution("a"))     // a
	fmt.Println(Solution("abba"))  // abba
	fmt.Println(Solution("world")) // dlrow
	fmt.Println(Solution("word"))  // drow

}
