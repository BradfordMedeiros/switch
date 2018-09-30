package main

import "os"
import "fmt"
import "io/ioutil"
import "errors"

import "./util/parse_args"
import "./util/input"
import "./util/query"
import "./util/statemachine"
import "./util/parse_program"
import "./util/parse_program/util/types"

func readFile(filepath string) string{
	b, err := ioutil.ReadFile(filepath) 
    if err != nil {
        fmt.Print(err)
    }
    return string(b)
}

func getHandleHookChange(hooks []types.Hook) func(string) {
	hookChange := func(hook string) {
		fmt.Println("hook change: ", hook)
	}
	return hookChange
}

func createBackendForProgram(programStructure parse_program.Program) (statemachine.StateMachine, error) {
	if !programStructure.Valid {
		return statemachine.StateMachine{}, errors.New("invalid program structure")
	}

	hookChange := getHandleHookChange(programStructure.Hooks)

	machine := statemachine.New(func(newstate string) {
		hookChange(newstate)
	})

	for _, rule := range(programStructure.Rules){
		fmt.Println("rule: ", rule)
		machine.AddState(rule.FromState, rule.ToState, rule.Transition)
	}
	for _, exit := range(programStructure.Exits){
		fmt.Println("exit: ", exit)
	}
	for _, hook := range(programStructure.Hooks){
		fmt.Println("hook: ", hook)
	}

	if programStructure.HasStart {
		fmt.Println("has start: ", programStructure.HasStart)
		fmt.Println("start: ", programStructure.Start)
		err := machine.ForceTransitionState(programStructure.Start.State)
		if err != nil {
			fmt.Println("error from starting")
			os.Exit(1)
		}else{
			fmt.Println("start is good")
		}
	}
	
	return machine, nil
}

func main(){
	options := parse_args.ParseArgs(os.Args[1:])

	programStructure := parse_program.Program { Valid: false }
	if options.ScriptPath.HasScript {
		fileContent := readFile(options.ScriptPath.ScriptPath)
		programStructure = parse_program.ParseProgram(fileContent)
	}

	machine, err := createBackendForProgram(programStructure)
	if err != nil {
		fmt.Println("program is invalid!!!!!!: ", err)
	}

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
