package constants

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"os"

	"github.com/joho/godotenv"
	"priyanshu.com/jwt/types"
	_ "github.com/mattn/go-sqlite3"
)

var HMacSigningingSecret []byte
var HashingSalt []byte
var Db *sql.DB

func initializeDb() error {
	var err error
	Db, err = sql.Open("sqlite3", "database.db")

	if err != nil {
		return err
	}

	q := `CREATE TABLE IF NOT EXISTS User (
		   id INTEGER PRIMARY KEY AUTOINCREMENT,
		   username text NOT NULL UNIQUE,
		   password text NOT NULL
		);`

	_, err = Db.Exec(q)
	if err != nil {
		return err
	}

	return nil
}

func getHmacSigningSecret() ([]byte, error) {
	secret := os.Getenv("HMAC_SIGNING_SECRET")
	if len(secret) == 0 {
		return nil, types.ErrorStr("Env var not found")
	}
	key, err := hex.DecodeString(secret)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func getHashingSalt() ([]byte, error) {
	secret := os.Getenv("HASHING_SALT")
	if len(secret) == 0 {
		return nil, types.ErrorStr("Env var not found")
	}

	return []byte(secret), nil
}

func generateKey(keyLen int) ([]byte, error) {
	key := make([]byte, keyLen)
	_, err := rand.Read(key)

	if err != nil {
		return nil, err
	}

	return key, nil
}

// This loads the .env file
// Looks for the "HMAC_SIGNING_SECRET" env variable
// It decoded the variable and set `HMacSigningingSecret`
func Init() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	if key, err := getHmacSigningSecret(); err != nil {
		return err
	} else {
		HMacSigningingSecret = key
	}

	if key, err := getHashingSalt(); err != nil {
		return err
	} else {
		HashingSalt = key
	}

	if err := initializeDb(); err != nil {
		return err
	}

	return nil
}
