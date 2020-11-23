package main

import (
	"fmt"
	"os"
)

const (
	user            = "jack"
	passwd          = "1888"
	errUser         = "Access denied for %q. \n"
	accessOK        = "Access granted to %q. \n"
	invalidPassword = "Invalid password for %q. \n"
	invalidArgs     = "Usage : [username] [password]"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(invalidArgs)
		return
	}
	if os.Args[1] == user && os.Args[2] == passwd {
		fmt.Printf(accessOK, user)
	} else if os.Args[2] != passwd {
		fmt.Printf(invalidPassword, user)
	} else if os.Args[1] != user {
		fmt.Println(errUser, os.Args[1])
	}

}
