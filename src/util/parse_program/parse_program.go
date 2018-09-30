package parse_program
import "strings"
import "fmt"
import "./util/types"

type Token struct {
	value string
}
type Unit struct {
	unitType string 
	rule types.Rule
	hook types.Hook
	start types.Start
	exit types.Exit
}
type ProgramStructure struct {
	Units []Unit
	Valid bool
}
func tokenize(program string) []Token {
	tokens := make([]Token, 0)
	lines := strings.Split(program, "\n")

	for _, line := range(lines){
		lineToAdd := strings.Trim(line, " ")
		if len(lineToAdd) > 0 {
			tokens = append(tokens, Token { value: lineToAdd })
		}
	}
	return tokens
}


func parseStatement(token Token)(Unit, bool){
	rule, isValidRule := types.TryParseRule(token.value)
	if isValidRule {
		return Unit { unitType: "rule", rule: rule }, true
	}

	hook, isValidHook := types.TryParseHook(token.value)
	if isValidHook {
		return Unit { unitType: "hook", hook: hook }, true
	}

	exit, isValidExit := types.TryParseExit(token.value) 
	if isValidExit {
		return Unit { unitType: "exit", exit: exit }, true
	}

	start, isValidStart := types.TryParseStart(token.value)
	if isValidStart {
		return Unit { unitType: "start", start: start }, true
	}

	return Unit { unitType: "none" }, false
}



func ParseProgram(program string) ProgramStructure {
	fmt.Println("parse program--------")

	tokens := tokenize(program)

	units := make([]Unit, len(tokens))
	for _, token := range(tokens){
		unit, isValid := parseStatement(token)
		if !isValid {
			fmt.Println("invalid program! exiting: (", token.value, ")")
			return  ProgramStructure { Valid: false, Units: units }
		}
		fmt.Println("! found valid element: ", unit.unitType)
		units = append(units, unit)
	}

	programStructure := ProgramStructure { Valid: true, Units: units }
	return programStructure
}