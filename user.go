package main

import "fmt"

// User ...
type User struct {
	ID        int    `json:"ID,-"`
	Login     string `json:"login"`
	Password  string `json:"Password,-"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Role      string `json:"Role,-"`
}

// Create создание нового пользователя в базе
func (u User) Create() error {
	row := connection.QueryRow(`INSERT INTO "User"
    ("Login", "Password", "FirstName", "LastName", "Role")
    VALUES ($1, $2, $3, $4, 'manager') RETURNING "ID"`,
		u.Login, u.Password, u.FirstName, u.LastName)
	e := row.Scan(&u.ID)
	if e != nil {
		return e
	}

	fmt.Println("Create new user with ID", u.ID)

	return nil
}

func (u User) Select() error {
	row := connection.QueryRow(
		`SELECT "Role", "FirstName", "LastName"
			FROM "User" WHERE "Login"=$1 AND "Password"=$2`,
		u.Login, u.Password)
	e := row.Scan(&u.Role, &u.FirstName, &u.LastName)
	if e != nil {
		return e
	}

	fmt.Println("Authorization user with Role", u.Role)

	return nil
}
