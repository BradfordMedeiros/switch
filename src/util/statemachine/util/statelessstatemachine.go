
// creates the core state machines needed to run the code
/*
		addState(stateFrom, stateTo, transition)
		removeState(stateFrom, stateTo, transition)
		getTransitions(state)
		getNextStates(state, transitions)

		and an interface with added state so that you can say

		// this probably should go in a seperate file
		setInitialState(state)
		getCurrentState()
		getTransitions()
		getNextStates(transitions)
		callback for state transition

*/	


package statelessmachine

import "fmt"
import "errors"
	
type State struct {
	statename string
	reachableStates map[string] *State
}	
type StatelessMachine struct {
	states map[string] State
}

func New() StatelessMachine{
	machine := StatelessMachine { states: make(map[string] State) }
	return machine
}
func (machine *StatelessMachine) initializeState(state string){
	_, hasState := machine.states[state]
	if !hasState {
		newState := State { statename: state, reachableStates: make(map[string]*State)}
		machine.states[state] = newState
	}
}
func (machine *StatelessMachine) AddState(stateFrom string, stateTo string, transitionName string) error{
	_, hasFromState := machine.states[stateFrom]
	if !hasFromState { machine.initializeState(stateFrom) }

	_, hasToState := machine.states[stateTo]
	if !hasToState { machine.initializeState(stateTo) }
	
	stateFromState := machine.states[stateFrom]
	stateToAdd := machine.states[stateTo]

	_, hasTransition := stateFromState.reachableStates[transitionName]
	if hasTransition {
		return errors.New("already has transition")
	}

	stateFromState.reachableStates[transitionName] = &stateToAdd
	return nil
}

func (machine *StatelessMachine) GetState() string{
	return "placeholder state"
}

func (machine *StatelessMachine) GetTransitions(stateFrom string) ([]string, error){
	state, hasState := machine.states[stateFrom]

	if !hasState {
		return make([]string, 0), errors.New("does not have state")
	}

	keys := make([]string, 0)
    for key := range state.reachableStates {
        keys = append(keys, key)
    }

	return keys, nil
}
func (machine *StatelessMachine) Transition() bool {
	return false
}

func Test(){
	fmt.Println("wow")
}