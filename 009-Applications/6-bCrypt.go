/*
Stores passwords using bCrypt
*/

package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `passworrd12345`
	fmt.Println("The password is : ", password);
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost);
	if err != nil {
		fmt.Println("Error while generating the hashed function of the password")
	}
	fmt.Println("The hashed password is :", bs)
	fmt.Printf("The hashed password is :%s\n", bs)

	
	loginPassword  := `password12345`;

	errorMessage := bcrypt.CompareHashAndPassword(bs, []byte(loginPassword));
	if errorMessage != nil {
		fmt.Println("Incorrect credentials for login");
		return;
	}
    fmt.Println("Login Success!!")

}