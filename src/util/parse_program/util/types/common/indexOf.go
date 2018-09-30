package common
import "fmt"

func IndexOf(stringToLookIn string, stringToSearchFor string) int {
	for i:= 0; i < len(stringToLookIn); i++ {
		fmt.Println(stringToLookIn[i])
	}
	return 0
}