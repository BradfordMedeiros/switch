package event_actions

import "fmt"
func Dump(getStartInfo func() string, getExitInfo func() string, getRuleInfo func() string, getHookInfo func() string) {
	fmt.Print(getStartInfo())
	fmt.Print(getExitInfo())
	fmt.Print(getRuleInfo())
	fmt.Print(getHookInfo())	
} 