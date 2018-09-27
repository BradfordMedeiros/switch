package statemachine
import statelessmachine "./util" 
import "errors"

type StateMachine struct {
	statelessmachine statelessmachine.StatelessMachine
	currentState *string
} 

func New() StateMachine{
	return StateMachine { 
		statelessmachine: statelessmachine.New(),
		currentState: nil,
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
func (machine *StateMachine) Transition() bool {
	return false
}
