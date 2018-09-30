package parse_program
import "strings"
import "fmt"
import "./util/filter"
import "./util/types"
import "./util/keywords"

type TokenUnion struct {
	rule types.Rule
	hook types.Hook
	start types.Start
	exit types.Exit
}

type Token struct {
	value string
}
type Unit struct {
	unitType string 
}
type ProgramStructure struct {
	Units []Unit
	Valid bool
}
func tokenize(program string) []Token {
	tokens := make([]Token, 0)
	lines := strings.Split(program, "\n")

	for _, line := range(lines){
		lineToAdd := filter.Filter(line, func(char rune) bool {
			return char != ' '
		})
		if len(lineToAdd) > 0 {
			fmt.Println(lineToAdd)
			tokens = append(tokens, Token { value: lineToAdd })
		}
	}
	return tokens
}


func getTryParseStatement(tryParseExit func(string) (types.Exit, bool)) func(token Token)(Unit, bool){
	tryParseStatement := func(token Token) (Unit, bool) {
		return Unit { unitType: "statement" }, true
	}
	return tryParseStatement
}



func ParseProgram(program string) ProgramStructure {
	fmt.Println("parse program--------")
	parseExit := types.GetTryParseExit(keywords.IsValidSymbol)
	tryParseStatement := getTryParseStatement(parseExit)

	tokens := tokenize(program)

	units := make([]Unit, len(tokens))
	for _, token := range(tokens){
		unit, isValid := tryParseStatement(token)
		if !isValid {
			return  ProgramStructure { Valid: false, Units: units }
		}
		units = append(units, unit)
	}

	programStructure := ProgramStructure { Valid: true, Units: units }
	return programStructure
}