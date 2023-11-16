package main

import (
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/scrypt"
	"priyanshu.com/jwt/constants"
)

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	err := constants.Init()
	if err != nil {
		log.Fatalf("Failed to initialize %s\n", err)
	}
	defer constants.Db.Close()

	// createUser("Priyanshu", "Abc@123")
	verified, _ := verifyUser("Priyanshu", "Abc@123")
	fmt.Println(verified)
}

func createUser(username string, password string) error {
	salt := constants.HashingSalt
	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)

	if err != nil {
		return err
	}

	_, err = constants.Db.Exec("INSERT INTO User (username, password) VALUES (?,?);", username, hash)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func verifyUser(username string, password string) (bool, error) {
	salt := constants.HashingSalt
	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)

	if err != nil {
		return false, err
	}

	var user User
	row := constants.Db.QueryRow("SELECT * FROM USER WHERE username=? LIMIT 1;", username)

	err = row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return false, err
	}

	fmt.Println(user)
	return user.Password == string(hash), nil
}
