package main

import (
	"database/sql"
	"log"

	//	"golang.org/x/crypto/scrypt"
	_ "github.com/mattn/go-sqlite3"
	"priyanshu.com/jwt/initializer"
)

type User struct {
	Id       int
	Username string
	Password string
}

func main() {
	err := initializer.Init()
	if err != nil {
		log.Fatalf("Failed to initialize %s\n", err)
	}

	db, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	q := `CREATE TABLE IF NOT EXISTS User (
		   id INTEGER PRIMARY KEY AUTOINCREMENT,
		   username text NOT NULL,
		   password text NOT NULL
		);`

	_, err = db.Exec(q)
	if err != nil {
		log.Fatalln(err)
	}
}

// func createUser(username string, password string) {
// 	salt := initializer.HashingSalt
// 	dk, err := scrypt.Key([]byte("some password"), salt, 16384, 8, 1, 32)
//
// 	fmt.Println(string(dk))
// 	fmt.Println(err)
// 	fmt.Println(salt)
// }
