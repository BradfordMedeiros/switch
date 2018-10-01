package types
import "strings"
import "strconv"
import "./common"

type Exit struct {
	State string
	Exitcode int
}

// exit 0 when somestate
func TryParseExit(value string) (Exit, bool){
	values := strings.Split(value, " ")
	values = common.FilterArray(values, func(val string) bool {
		return val != " "
	})
	
	if len(values) != 4 {
		return Exit{}, false
	}
	if values[0] != "exit" {
		return Exit{}, false
	}

	if values[2] != "when" {
		return Exit{}, false
	}

	exitcode, err := strconv.Atoi(values[1])
	if err != nil {
		return Exit{}, false
	}

	return Exit { State: values[3], Exitcode: exitcode }, true
	return Exit { }, false
}

func (exit *Exit) AsString() string{
	return "exit placeholder"
}