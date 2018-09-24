
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