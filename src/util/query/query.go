
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

// : before command is a query 
// else it should be interpreted as code

func GetHandleQuery(getState func() string) func(string){
	return func(queryString string){
		if (queryString == ":state"){
			fmt.Println(getState())
		}else if (queryString == ":transitions"){
			fmt.Println("get transitions")
		}else{
			fmt.Println("unknown query " + queryString)
		}
	}

}


