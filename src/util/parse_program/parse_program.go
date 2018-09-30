package parse_program
import "strings"
import "fmt"
import "./util/types"

type Token struct {
	value string
}
type Unit struct {
	UnitType string 
	Rule types.Rule
	hook types.Hook
	Start types.Start
	Exit types.Exit
}
type Program struct {
	Valid bool
	Rules []types.Rule
	Hooks []types.Hook
	Exits []types.Exit

	Start types.Start
	HasStart bool
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
		return Unit { UnitType: "rule", Rule: rule }, true
	}

	hook, isValidHook := types.TryParseHook(token.value)
	if isValidHook {
		return Unit { UnitType: "hook", hook: hook }, true
	}

	exit, isValidExit := types.TryParseExit(token.value) 
	if isValidExit {
		return Unit { UnitType: "exit", Exit: exit }, true
	}

	start, isValidStart := types.TryParseStart(token.value)
	if isValidStart {
		return Unit { UnitType: "start", Start: start }, true
	}

	return Unit { UnitType: "none" }, false
}

func ParseProgram(program string) Program{
	tokens := tokenize(program)

	var start types.Start
	hasStart := false

	units := make([]Unit, 0)
	rules := make([]types.Rule, 0)
	hooks := make([]types.Hook, 0)		// todo hooks
	exits := make([]types.Exit, 0)

	fmt.Println("num tokens: ", len(tokens))
	for _, token := range(tokens){
		unit, isValid := parseStatement(token)
		
		if unit.UnitType == "rule" {
			rules = append(rules, unit.Rule)
		}else if unit.UnitType == "exit" {
			exits = append(exits, unit.Exit)
		}else if unit.UnitType == "start" {
			start = unit.Start
			hasStart = true
		}


		if !isValid {
			fmt.Println("invalid program! exiting: (", token.value, ")")
			return  Program { Valid: false }
		}else{
			units = append(units, unit)
		}
	}

	programStructure := Program {
		Valid: true, 
		Start: start,
		HasStart: hasStart,
		Exits: exits, 
		Rules: rules, 
		Hooks: hooks, 
	}

	return programStructure
}