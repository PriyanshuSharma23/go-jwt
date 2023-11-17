package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"priyanshu.com/jwt/constants"
	"priyanshu.com/jwt/users"
)

func main() {
	err := constants.Init()
	if err != nil {
		log.Fatalf("Failed to initialize %s\n", err)
	}
	defer constants.Db.Close()

	fmt.Println("Select Action:")
	fmt.Println("1. Sign In")
	fmt.Println("2. Sign Up")
	fmt.Println("3. Check Auth")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	switch option {
	case 1:
	signIn()
	case 2:
	signUp()
	case 3:
	checkAuth()
	default:
	fmt.Println("Select a valid operation")
	}

}

func signUp() {
	fmt.Println("Sign Up")
	scanner := bufio.NewScanner(os.Stdin)

	// Ask for username
	fmt.Print("Username: ")
	var username string
	scanner.Scan()
	username = strings.TrimSpace(scanner.Text())

	// Ask for password
	fmt.Print("Password: ")
	var password string
	scanner.Scan()
	password = strings.TrimSpace(scanner.Text())

	// Create user
	user, err := users.CreateUser(username, password)
	if err != nil {
		log.Printf("Failed to create user: %s\n", err)
		return
	}

	// Create accesstoken
	token, err := users.GenerateJwt(user)
	if err != nil {
		log.Printf("Failed to create token: %s\n", err)
		return
	}

	fmt.Printf("Successfully signed up, here is the token:\n%s\n", token)
}

func signIn() {
	fmt.Println("Sign In")
	scanner := bufio.NewScanner(os.Stdin)

	// Ask for username
	fmt.Print("Username: ")
	var username string
	scanner.Scan()
	username = strings.TrimSpace(scanner.Text())

	// Ask for password
	fmt.Print("Password: ")
	var password string
	scanner.Scan()
	password = strings.TrimSpace(scanner.Text())

	// Create user
	ok, user, err := users.VerifyUser(username, password)
	if err != nil {
		log.Printf("Failed to verify user: %s\n", err)
		return
	}
	if !ok {
		fmt.Println("User does not exist, try signing up...")
		return
	}

	// Create accesstoken
	token, err := users.GenerateJwt(*user)
	if err != nil {
		log.Printf("Failed to create token: %s\n", err)
		return
	}

	fmt.Printf("Successfully signed in, here is the token:\n%s\n", token)
}

func checkAuth() {
	fmt.Println("Paste token")
	scanner := bufio.NewScanner(os.Stdin)

	// Ask for username
	var token string
	scanner.Scan()
	token = strings.TrimSpace(scanner.Text())

	user, err := users.VerifyToken(token)
	if err != nil {
		fmt.Printf("Invalid token: %s\n", err)
		return
	}

	fmt.Println("\nValid!")
	fmt.Printf("%#v", user)
}
