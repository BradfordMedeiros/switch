package event_actions

import "fmt"
func Dump(getStartInfo func() string, getExitInfo func() string, getRuleInfo func() string, getHookInfo func() string) {
	fmt.Println(getStartInfo())
	fmt.Println(getExitInfo())
	fmt.Println(getRuleInfo())
	fmt.Println(getHookInfo())	
} 