package statemachine
import statelessmachine "./util" 

type StateMachine struct {
	name string 	
}

func New() StateMachine{
	return StateMachine {  name: "test" }
}

func (machine *StateMachine) AddState(stateFrom string, stateTo string, transitionName string){
	statelessmachine.Test()
}

func (machine *StateMachine) GetState() string{
	return "placeholder state"
}

