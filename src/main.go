package main

import "os"
import "fmt"
import "io/ioutil"

import "./util/parse_args"
import "./util/input"
import "./util/query"
import "./util/statemachine"
import "./util/parse_program"

func readFile(filepath string) string{
	b, err := ioutil.ReadFile(filepath) 
    if err != nil {
        fmt.Print(err)
    }
    return string(b)
}

func createBackendForProgram(programStructure parse_program.ProgramStructure) statemachine.StateMachine{
	machine := statemachine.New()

	machine.AddState("wet", "frozen", "freeze")		
	machine.AddState("frozen", "wet", "unfreeze")
	machine.AddState("wet", "dry", "airdry")
	machine.AddState("dry", "wet", "rain")

	return machine
}

func main(){
	options := parse_args.ParseArgs(os.Args[1:])
	fmt.Println(options)

	programStructure := parse_program.ProgramStructure { Valid: false }
	if options.ScriptPath.HasScript {
		fileContent := readFile(options.ScriptPath.ScriptPath)
		programStructure = parse_program.ParseProgram(fileContent)
	}

	machine := createBackendForProgram(programStructure)
	

	handleQuery := query.GetHandleQuery(
		machine.GetState, 
		machine.GetTransitions,
		machine.Transition,
	)

	commandChannel := make(chan string) 
		go input.StartRepl(commandChannel)
		for true {
			select {
				case commandString := <-commandChannel: {
					handleQuery(commandString)
				}

			}
		}
}
