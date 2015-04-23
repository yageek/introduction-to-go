package main

import "fmt"

type AccessLevel int
type User struct {
	Login    string
	Password string
}
type Administrator struct {
	User  User
	Level AccessLevel
}

const (
	Nothing AccessLevel = iota
	God
)

func main() {
	admin := &Administrator{Level: God, User: User{Login: "su", Password: "3.1419527"}}
	fmt.Println("Admin:", admin)
}
