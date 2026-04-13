package dog

import (
	"fmt"
	"strings"
)

func WhenGrownUp(s string) string {
	fmt.Println("Value of s", s)
	if strings.TrimSpace(s) == "" {
		s = "I am not Mahala the adult"
	}
	fmt.Println("Value of s", s)

	return "When the puppy grows up it says: " + strings.ToUpper(s)
}
