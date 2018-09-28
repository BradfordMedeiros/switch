
package parse_program

/*
	(thing -> go) : wow   # statement (no label)
	(wow->yo):yes
//
	when - dry the wet : someexternalevent  #statement
	start as wet 	# start

	exit 0 when frozen #exit statement
	exit 1 when airdry #exit statement
*/

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

// @todo implement
func tokenize(program string) []Token {
	tokens := make([]Token, 0)
	return tokens
}

// @todo implement
func parseToken(token Token) Unit {
	return Unit { unitType: "test" }
}



func ParseProgram(program string) ProgramStructure {
	tokens := tokenize(program)

	units := make([]Unit, len(tokens))
	for _, token := range(tokens){
		units = append(units, parseToken(token))
	}

	programStructure := ProgramStructure { Valid: true, Units: units }
	return programStructure
}