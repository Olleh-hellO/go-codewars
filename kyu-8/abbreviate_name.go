/* Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

It should look like this:

Sam Harris => S.H

patrick feeney => P.F */

package kata

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	names := strings.Split(name, "")
	firstName := names[0]
	lastName := names[1]

	firstNameInit := strings.ToUpper(string(firstName[0:1]))
	lastNameInit := string.ToUpper(string(lastName[0:1]))

	return fmt.Sprintf("%s.%s", firstNameInit, lastNameInit)

}
