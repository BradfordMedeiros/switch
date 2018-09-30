package keywords

/*func isKeyword(wordstring) bool {
    switch category {
    case 
    	"when",
        "start",
        "exit",
        return true
    }
    return false
}*/

// must only contain [a-z, A-Z, 0-9]
// A = 65, Z= 90, a = 97, z = '122'
func IsValidSymbol(wordstring string) bool {
	/*if isKeyword(wordstring){
		return false
	}*/

	for _, char := range(wordstring){
		if char < 'a' {
			return false
		}
		if char > 'Z' &&  char < 'a' {
			return false 
		}
		if char > 'z' {
			return false
		}
	}
	return true
}