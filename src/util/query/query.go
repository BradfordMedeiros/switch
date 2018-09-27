
/*
		handle the possile queries (as actual text), and return a command to run , 
		and provide reply for the user

		so parse input like:

		quit 
		:list-transitions (return next states with how to get there)
		:transition (return ok or null)
		:state

		
*/

package query 

import "fmt"
import "strings"

// : before command is a query 
// else it should be interpreted as code

func GetHandleQuery(
	getState func() (string, error), 
	getTransitions func() ([]string, error),
	transition func(string) error,
) func(string){
	return func(mainQueryString string){
		parts := strings.Split(mainQueryString, " ")
		queryString := parts[0]

		if (queryString == ":s" || queryString == ":state"){
			state, error := getState()
			if error != nil {
				fmt.Println(error)
				return
			}
			fmt.Println(state)
		}else if (queryString == ":t" || queryString == ":transitions"){
			transitions, _ := getTransitions()
			fmt.Println(transitions)
		}else if (queryString == ":m" || queryString == "move"){
			transitionState := parts[1]
			err := transition(transitionState)
			if err != nil {
				fmt.Println("error ", err)
			}else{
				fmt.Println("ok")
			}
		}else{
			fmt.Println("unknown query " + queryString)
		}
	}

}


