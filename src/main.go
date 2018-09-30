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

func generateHookChangeMap(rules []types.Rule, hooks []types.Hook) map[string]map[string][]string{
	hookchange := make(map[string]map[string][]string)

	// [fromState][toState] => label
	// label => eventname 
	// gets
	// [fromState][toState] => [] eventname

	// creating mapping based on hook transitions
	// but the associated rule label is the only ones we should care about
	// so the string at the end probably should be array assocated rule labels

	for _, rule := range(rules){
		fmt.Println("creating?")
		if rule.HasLabel {
			fmt.Println("adding label: ", rule.Label)

			if hookchange[rule.FromState] == nil {
				hookchange[rule.FromState] = make(map[string][]string)
			}
			if hookchange[rule.FromState][rule.ToState] == nil {
				hookchange[rule.FromState][rule.ToState] = make([]string, 0)
			}

			hookchange[rule.FromState][rule.ToState] = append(
				hookchange[rule.FromState][rule.ToState],
				rule.Label, // this should really be the events not the labels
			)
		}
	}
	return hookchange
}

func getHandleHookChange(
	hooks []types.Hook, 
	rules []types.Rule, 
	exits []types.Exit,
	callLabel func(string),
	exitFunc func(code int),
) func(string, string) {
	hookChangeMap := generateHookChangeMap(rules, hooks)

	hookChange := func(laststate string, newstate string) {
		if hookChangeMap[laststate] !=nil {
			labels, hasMapping := hookChangeMap[laststate][newstate]
			if hasMapping {
				for _, label := range(labels){
					callLabel(label)
				}
			}
		}

		for _, exit := range(exits){
			if exit.State == newstate {
				exitFunc(exit.Exitcode)
			}
		}
	}


	return hookChange
}

func createBackendForProgram(programStructure parse_program.Program) (statemachine.StateMachine, error) {
	if !programStructure.Valid {
		return statemachine.StateMachine{}, errors.New("invalid program structure")
	}

	machine := statemachine.New(getHandleHookChange(
		programStructure.Hooks,
		programStructure.Rules, 
		programStructure.Exits,
		func(label string) {
			fmt.Println("call label (should be event): ", label)
		},
		func(code int){
			fmt.Println("exit with code: ", code)
			os.Exit(code)
		},
	))

	for _, rule := range(programStructure.Rules){
		fmt.Println("rule: ", rule)
		machine.AddState(rule.FromState, rule.ToState, rule.Transition)
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
		fileContentBytes, _ := ioutil.ReadFile(options.ScriptPath.ScriptPath)
		programStructure = parse_program.ParseProgram(string(fileContentBytes))
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
