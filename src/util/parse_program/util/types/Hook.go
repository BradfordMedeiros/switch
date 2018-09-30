package types
import "strings"
import "./common"

type Hook struct {
	Label string
	Event string
}

func TryParseHook(value string) (Hook, bool){
	values := strings.Split(common.Filter(value, func(val rune) bool {
		return val != ' '
	}), "-")

	if len(values) != 3 {
		return Hook{}, false
	}
	if values[0] != "when" {
		return Hook{}, false
	}

	return Hook { Label: values[1], Event: values[2] }, true
}