
/*
	parses flags, and returns options for how the program should run
	
*/

package parse_args

import "flag"

type Options struct {
	ScriptPath ScriptPath;
}

type ScriptPath struct {
	HasScript bool;
	ScriptPath string;
}
func (i *ScriptPath) String() string {
    return i.ScriptPath
}
func (i *ScriptPath) Set(value string) error {
	i.ScriptPath = value
    if value == "" {
    	i.HasScript = false
    }else{
    	i.HasScript = true
    }
    return nil
}


func ParseArgs(arguments []string) Options{
	fs := flag.NewFlagSet("main", flag.ExitOnError)

    var myScriptPath ScriptPath
    fs.Var(&myScriptPath, "f", "program to run")

    fs.Parse(arguments)
    return Options { ScriptPath: myScriptPath }
}

