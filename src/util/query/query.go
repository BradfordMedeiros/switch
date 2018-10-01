package query 

import "fmt"
import "strings"
import "errors"

func GetHandleQuery(
	getState func() (string, error), 
	getTransitions func() ([]string, error),
	transition func(string) error,
	printResponse bool,
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
			if len(parts) < 2 {
				fmt.Println("error: ", errors.New("no transition specified"))
				return
			}
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


