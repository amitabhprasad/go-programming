package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

/**
- Encrypt sensitive data
**/
func bcryptExample() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println("bs ", bs)
	fmt.Println("password ", s)
	loginp := `password124`
	err = bcrypt.CompareHashAndPassword(bs, []byte(loginp))
	if err != nil {
		fmt.Println("Error during match ", err)
		return
	}
	fmt.Println("Login success!")
}
