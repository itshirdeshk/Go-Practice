package main

import (
	"github.com/fatih/color"
	"github.com/itshirdeshk/go_learning/auth"
	"github.com/itshirdeshk/go_learning/user"
)

func main() {
	auth.LoginWithCredentials("hirdesh", "12345")

	user := user.User{
		Name:  "Hirdesh",
		Email: "hirdesh@example.com",
	}

	color.Red(user.Email)
}
