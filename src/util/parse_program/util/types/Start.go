package types

type Start struct {

}

func tryParseStart(value string) (Start, bool){
	return Start{ }, true
}