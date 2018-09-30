package event_actions

import "os"

func ExitProgram(exitcode int){
	os.Exit(exitcode)
}