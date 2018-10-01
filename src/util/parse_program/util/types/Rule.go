
package types
import "strings"
import "errors"
import "./common"

// (wet -> frozen) : freeze
// or
// (wet -> frozen) : somelabelhere - freeze

type Rule struct {
	FromState string
	ToState string
	Transition string
	HasLabel bool
	Label string  
}


type leftSide struct {
	fromState string
	toState string
}
type rightSide struct {
	transition string
	label string
	hasLabel bool
}

func parseLeftSide(leftSideString string) (leftSide, error) {
	if leftSideString[0] != '(' || leftSideString[len(leftSideString)-1] != ')' {
		return leftSide {}, errors.New("invalid")
	}
	leftSideSplit := strings.Split(leftSideString, "->")
	if len(leftSideSplit) != 2 {
		return leftSide{}, errors.New("invalid")
	}
	fromState := leftSideSplit[0][1:]
	toState := leftSideSplit[1][0:len(leftSideSplit[1])-1]
	return leftSide{ fromState: fromState, toState: toState }, nil
}

func parseRightSide(rightSideString string) (rightSide, error) {
	rightSideSplit := strings.Split(rightSideString, "-")
	if len(rightSideSplit) != 1 && len(rightSideSplit) != 2 {
		return rightSide{}, errors.New("invalid num params")
	}
	if len(rightSideSplit) == 1 {
		return rightSide { transition: rightSideSplit[0], label: "", hasLabel: false }, nil
	}
	return rightSide{ transition: rightSideSplit[0], label: rightSideSplit[1], hasLabel: true }, nil
}

func TryParseRule(value string) (Rule, bool){
	values := strings.Split(common.Filter(value, func(val rune) bool {
		return val != ' '
	}), ":")

	if len(values) != 2 {
		return Rule{}, false
	}

	ls, errLeftSide := parseLeftSide(values[0])
	rs, errRightSide := parseRightSide(values[1])

	if errLeftSide != nil || errRightSide != nil {
		return Rule{}, false
	}

	rule := Rule { 
		FromState: ls.fromState, 
		ToState: ls.toState, 
		Transition: rs.transition, 
		Label: rs.label,
		HasLabel: rs.hasLabel,
	}
	return rule, true
}

func (rule *Rule) AsString() string{
	return "rule placeholder"
}