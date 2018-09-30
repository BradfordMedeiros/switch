package types

type Start struct {

}

func TryParseStart(value string) (Start, bool){
	return Start{ }, false
}