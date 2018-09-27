/*

	handles the io from the console , maybe from stdin, etc and just delivers the raw stuff forward
*/


package input

import "bufio"
import "os"
import "fmt"

func StartRepl (commandChannel chan string) {
	fmt.Println("Switch 1.0.0")
	fmt.Println("---------------------")
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {		// happens if user just hits enter, so filter it out
			continue
		}
		commandChannel <- text 
	}
}