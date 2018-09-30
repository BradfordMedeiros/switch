package types
import "strings"
import "./common"

type Start struct {
	state string
}

// start as somestate
func TryParseStart(value string) (Start, bool){
	values := strings.Split(value, " ")
	values = common.FilterArray(values, func(val string) bool {
		return val != " "
	})
	
	if len(values) != 3 {
		return Start{}, false
	}
	if values[0] != "start" {
		return Start{}, false
	}
	if values[1] != "as" {
		return Start{}, false
	}

	return Start{ state: values[2] }, true
}