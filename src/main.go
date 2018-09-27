package main

import "./util/input"
import "./util/query"
import "./util/statemachine"



func main(){
	machine := statemachine.New()

	machine.AddState("wet", "frozen", "freeze")		
	machine.AddState("frozen", "wet", "unfreeze")
	machine.AddState("wet", "dry", "airdry")
	machine.AddState("dry", "wet", "rain")

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
