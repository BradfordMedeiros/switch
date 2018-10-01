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


// this probably should be moved to parse_program 
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


	if options.ScriptPath.HasScript && options.InlineScript.HasScript {
		fmt.Println("invalid options")
		os.Exit(1)
	}else if options.ScriptPath.HasScript {
		fileContentBytes, _ := ioutil.ReadFile(options.ScriptPath.ScriptPath)
		programStructure = parse_program.ParseProgram(string(fileContentBytes))
	}else if options.InlineScript.HasScript {
		programStructure = parse_program.ParseProgram(options.InlineScript.InlineScriptContent)
	}else{
		fmt.Println("no options specified")
		os.Exit(1)
	}



	if options.Dump {
		event_actions.Dump(func() string {
			return programStructure.Start.AsString() + "\n"
		}, func() string {
			eventActions := ""
			for _, exit := range(programStructure.Exits){
				eventActions = eventActions + exit.AsString() + "\n"
			}
			return eventActions

		}, func() string {
			eventActions := ""
			for _, rule := range(programStructure.Rules){
				eventActions = eventActions + rule.AsString() + "\n"
			}
			return eventActions
		}, func() string {
			eventActions := ""
			for _, hook := range(programStructure.Hooks){
				eventActions = eventActions + hook.AsString() + "\n"
			}
			return eventActions

		})
		os.Exit(0)
	}
	
	machine, err := createBackendForProgram(programStructure)
	if err != nil {
		fmt.Println("program is invalid!!!!!!: ", err)
	}

	handleQuery := query.GetHandleQuery(
		machine.GetState, 
		machine.GetTransitions,
		machine.Transition,
		options.RestrictTransition,
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
