
package parse_program

type ProgramStructure struct {
	Valid bool
}

func ParseProgram(program string) ProgramStructure {
	programStructure := ProgramStructure { Valid: true }
	return programStructure
}