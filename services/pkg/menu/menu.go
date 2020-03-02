package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	handler "github.com/aditya-nagare/d2h/services/pkg/handlers"
)

//Menu **
func Menu(handler *handler.HandlerStruct) {
user:
	fmt.Println("\nPlease Select Your Account:")
	users, _ := handler.ListUsers()
	fmt.Printf("\n\nEnter Your ID:")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	id, _ := strconv.Atoi(string(char))
	if id > 0 {
		id = id - 1
	} else {
		fmt.Println("\nUser ID Cannot be zero")
		goto user
	}
	user := handler.SelectUser(users[id].ID)
	fmt.Println("\nWelcome to D2H:\n", "\n1.View current balance in the account", "\n2.Recharge Account", "\n3.View available packs, channels and services", "\n4.Subscribe to base packs", "\n5.Add channels to an existing subscription", "\n6.Subscribe to special services", "\n7.View current subscription details", "\n8.Update email and phone number for notifications", "\n0.Exit")
	fmt.Printf("\nEnter the option:")
	reader = bufio.NewReader(os.Stdin)
	char, _, err = reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	menu, _ := strconv.Atoi(string(char))
	fmt.Printf("You choose option number: %d\n\n", menu)

	switch menu {
	case 1:
		handler.ViewBalance(user)
		MenuPrompt(handler)
	case 2:
		handler.Recharge(user)
		MenuPrompt(handler)
	case 3:
		handler.ViewPacks()
		MenuPrompt(handler)
	case 4:
		handler.SubscribeBasePack(*user)
		MenuPrompt(handler)
	case 5:
		handler.AddChannel(*user)
		MenuPrompt(handler)
	case 6:
		handler.SubscribeSpecialService(*user)
		MenuPrompt(handler)
	case 7:
		handler.ViewSubscription(*user)
		MenuPrompt(handler)
	case 8:
		handler.UpdateInfo(*user)
		MenuPrompt(handler)
	case 0:
		fmt.Println("Bye!")
		os.Exit(0)

	}

}

//MenuPrompt **
func MenuPrompt(handler *handler.HandlerStruct) {
menu:
	fmt.Printf("\nDo Your Wish to continue(Y|y to continue or N|n to exit):")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	switch char {
	case 'Y':
		Menu(handler)
	case 'y':
		Menu(handler)
	case 'N':
		fmt.Println("Bye!")
		os.Exit(0)
	case 'n':
		fmt.Println("Bye!")
		os.Exit(0)
	default:
		fmt.Println("Enter A Valid Input!")
		goto menu
	}
}
