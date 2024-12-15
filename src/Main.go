package main

import (
	"fmt"
	"java/repl"
	"os"
	"os/user"
)

func main() {
	logo := `
     ____.                    
    |    |____ ___  _______   
    |    \__  \\  \/ /\__  \  
/\__|    |/ __ \\   /  / __ \_
\________(____  /\_/  (____  /
              \/           \/ 
`
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf(logo, '\n')
	fmt.Printf("Hello %s! This is the Java programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
