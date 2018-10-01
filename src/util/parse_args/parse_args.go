
/*
	parses flags, and returns options for how the program should run
	
*/

package parse_args

import "flag"

type Options struct {
	ScriptPath ScriptPath;
    InlineScript InlineScript;
    RestrictTransition bool;
}
type ScriptPath struct {
	HasScript bool;
	ScriptPath string;
}
type InlineScript struct {
    HasScript bool;
    InlineScriptContent string;
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

func (i *InlineScript) String() string {
    return i.InlineScriptContent
}
func (i *InlineScript) Set(value string) error {
    i.InlineScriptContent= value
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

    var myInlineScript InlineScript
    fs.Var(&myInlineScript, "i", "inline script")

    restrictTransition := fs.Bool("t", false, "restrict input to transition only")
    fs.Parse(arguments)

    return Options { 
        ScriptPath: myScriptPath,
        InlineScript: myInlineScript,
        RestrictTransition: *restrictTransition,
    }
}

