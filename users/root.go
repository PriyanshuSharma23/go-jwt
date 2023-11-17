package users

import (
	"golang.org/x/crypto/scrypt"
	"priyanshu.com/jwt/constants"
)

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func CreateUser(username string, password string) (User, error) {
	salt := constants.HashingSalt
	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)

	if err != nil {
		return User{}, err
	}

	result, err := constants.Db.Exec("INSERT INTO User (username, password) VALUES (?,?);", username, hash)
	if err != nil {
		return User{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return User{}, err
	}

	return User{Id: int(id), Username: username}, nil
}

func VerifyUser(username string, password string) (bool, *User, error) {
	salt := constants.HashingSalt
	hash, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, 32)

	if err != nil {
		return false, nil, err
	}

	var user User
	row := constants.Db.QueryRow("SELECT * FROM USER WHERE username=? LIMIT 1;", username)

	err = row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return false, nil, err
	}

	return user.Password == string(hash), &user, nil
}
