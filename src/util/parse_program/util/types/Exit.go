
package types
//import "strings"

type Exit struct {

}


func GetTryParseExit(isValidKeyword func(string) bool) func(value string)(Exit,bool) {
	tryParseExit := func (value string) (Exit, bool) {
	
		
		return Exit{} , false
	}
	return tryParseExit
}
