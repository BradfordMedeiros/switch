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
import "./util/event_actions"

func getHasEvent(hooks []types.Hook) (getEventForLabel func(types.Rule) (bool, string)){
	doGetEventForLabel := func(rule types.Rule) (bool, string){
		for _, hook := range(hooks) {
			if hook.Label == rule.Label {
				return true, hook.Event
			}
		}
		return false, ""
	}
	return  doGetEventForLabel
}

func generateHookChangeMap(rules []types.Rule, hooks []types.Hook) map[string]map[string][]string{
	hookchange := make(map[string]map[string][]string)
	hasEvent := getHasEvent(hooks)

	for _, rule := range(rules){
		hasEventHook, eventHookName := hasEvent(rule)
		if !hasEventHook {
			continue
		}
		if rule.HasLabel {
			if hookchange[rule.FromState] == nil {
				hookchange[rule.FromState] = make(map[string][]string)
			}
			if hookchange[rule.FromState][rule.ToState] == nil {
				hookchange[rule.FromState][rule.ToState] = make([]string, 0)
			}

			hookchange[rule.FromState][rule.ToState] = append(
				hookchange[rule.FromState][rule.ToState],
				eventHookName, 
			)
		}
	}
	return hookchange
}

func getHandleHookChange(
	hooks []types.Hook, 
	rules []types.Rule, 
	exits []types.Exit,
	callEvent func(string),
	exitFunc func(code int),
) func(string, string) {
	hookChangeMap := generateHookChangeMap(rules, hooks)
	hookChange := func(laststate string, newstate string) {
		if hookChangeMap[laststate] !=nil {
			labels, hasMapping := hookChangeMap[laststate][newstate]
			if hasMapping {
				for _, label := range(labels){
					callEvent(label)
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
		event_actions.PublishEvent,
		event_actions.ExitProgram,
	))

	for _, rule := range(programStructure.Rules){
		machine.AddState(rule.FromState, rule.ToState, rule.Transition)
	}
	
	if programStructure.HasStart {
		err := machine.ForceTransitionState(programStructure.Start.State)
		if err != nil {
			fmt.Println("error from starting")
			// maybe think of exit code to distinguish internal vs user program
			os.Exit(1)
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
		true,
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
