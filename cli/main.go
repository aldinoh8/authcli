package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/aldinoh8/authcli/config"
	"github.com/aldinoh8/authcli/handler"
)

type Cli struct {
	UserHandler handler.User
}

func (c Cli) AuthMenu() {
	fmt.Println("\n\n -----")
	fmt.Println("Welcome to Example App")
	fmt.Println("[COMMAND]: Description")
	fmt.Println("[register]: Register/create new user")
	fmt.Println("[login]: login into existing user")
	fmt.Println("press anything to exit")

	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err.Error())
	}

	switch input {
	case "register":
		c.Register()
	case "login":
		c.Login()
	default:
		os.Exit(1)
	}
}

func (c Cli) MainMenu() {
	fmt.Println("[COMMAND]: Description")
	fmt.Println("[list-products]: Retrieve all products")
	fmt.Println("[list-orders]: Retrieeve all ordrs")
	fmt.Println("[create-products]: Create new products")
	fmt.Println("press anything to exit")
}

func (c Cli) Register() {
	// handlerUser.Register
	var username, password string
	var age int

	fmt.Println("input username:")
	_, errUsername := fmt.Scanln(&username)

	fmt.Println("input password:")
	_, errPassword := fmt.Scanln(&password)

	fmt.Println("input age:")
	_, errAge := fmt.Scanln(&age)

	if errUsername != nil ||
		errPassword != nil ||
		errAge != nil {
		log.Fatal("failed to scan input")
	}

	err := c.UserHandler.Register(username, password, age)
	if err != nil {
		fmt.Println("failed to create new user ...")
	} else {
		fmt.Println("success register user")
	}
	c.AuthMenu()
}

func (c Cli) Login() {
	var username, password string

	fmt.Println("input username:")
	_, errUsername := fmt.Scanln(&username)

	fmt.Println("input password:")
	_, errPassword := fmt.Scanln(&password)

	if errUsername != nil ||
		errPassword != nil {
		log.Fatal("failed to scan input")
	}
	user, err := c.UserHandler.Login(username, password)
	if err != nil {
		fmt.Println("error: ", err.Error())
		c.Login()
	} else {
		fmt.Println("user berhasil login")
		config.LoggedInUser = user
		c.MainMenu()
	}

}
