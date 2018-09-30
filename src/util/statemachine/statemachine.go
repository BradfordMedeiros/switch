package statemachine
import statelessmachine "./util" 
import "errors"

type StateMachine struct {
	statelessmachine statelessmachine.StatelessMachine
	currentState *string
	onStateChanged func(string, string)
} 

func New(onStateChanged func(string, string)) StateMachine{
	return StateMachine { 
		statelessmachine: statelessmachine.New(),
		currentState: nil,
		onStateChanged: onStateChanged,
	}
}

func (machine *StateMachine) AddState(stateFrom string, stateTo string, transitionName string){
	if machine.currentState == nil {
		machine.currentState = &stateFrom
	}
	machine.statelessmachine.AddState(stateFrom, stateTo, transitionName)
}

func (machine *StateMachine) GetState() (string, error){
	if machine.currentState == nil {
		return "", errors.New("no initial state")
	}
	return *machine.currentState, nil
}

func (machine *StateMachine) GetTransitions() ([]string, error){
	return machine.statelessmachine.GetTransitions(*machine.currentState)
}
func (machine *StateMachine) Transition(transitionName string) error {
	hasTransition, err := machine.statelessmachine.HasTransition(*machine.currentState, transitionName)
	if err != nil {
		return err
	}
	if hasTransition {
		stateName, _ := machine.statelessmachine.GetStateForTransition(*machine.currentState, transitionName)
		lastState := *machine.currentState
		machine.currentState = &stateName
		machine.onStateChanged(lastState, *machine.currentState)
	}else{
		return errors.New("no transition of that type " + transitionName + " for " + *machine.currentState)
	}
	return nil
}
func (machine *StateMachine) ForceTransitionState(transitionName string) error {
	if machine.statelessmachine.HasState(transitionName) {
		machine.currentState = &transitionName
		return nil
	}else{
		return errors.New("no state named " + transitionName)
	}
}
