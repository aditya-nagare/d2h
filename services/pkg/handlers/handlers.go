package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aditya-nagare/d2h/services/pkg/models"
	service "github.com/aditya-nagare/d2h/services/pkg/service"
)

//HandlerStruct **
type HandlerStruct struct {
	svc service.ServiceInterface
}

//NewHandlerImpl **
func NewHandlerImpl(service service.ServiceInterface) *HandlerStruct {
	return &HandlerStruct{svc: service}
}

//ViewBalance **
func (h HandlerStruct) ViewBalance(user *models.User) {
	fmt.Printf("\nHey %s,\n", user.Name)
	fmt.Printf("Your Current Balance is:%d\n", user.Balance)
}

//Recharge **
func (h HandlerStruct) Recharge(user *models.User) {
	fmt.Printf("2.Recharge Account")
	fmt.Printf("\nEnter Amount:")
	reader := bufio.NewReader(os.Stdin)
	char, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	char = strings.TrimSuffix(char, "\n")
	fmt.Println(char)
	amount, err := strconv.Atoi(char)
	if err != nil {
		fmt.Println("Enter Valid Ammount!", err)
	}
	fmt.Printf("You Entered :%d\n\n", amount)
	user, err = h.svc.Recharge(amount, user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nHey %s,\n", user.Name)
	fmt.Printf("Your Recharge Amount is:%d\n", amount)
	fmt.Printf("Your Current Balance is:%d\n", user.Balance)

}

//ViewPacks **
func (h HandlerStruct) ViewPacks() {
	channels, services, err := h.svc.ViewPacks()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nAvailable Channels for Subscription:")
	for _, v := range channels {
		fmt.Println(v.Name, ": Rs.", v.Price)
	}

	fmt.Println("\nAvailable Services for Subscription:")
	for _, v := range services {
		fmt.Println(v.Name, ": Rs.", v.Price)
	}
}

//SubscribeBasePack **
func (h HandlerStruct) SubscribeBasePack(user models.User) {
	packs, err := h.svc.ListPacks()
	for _, v := range packs {
		fmt.Printf("ID:%d \tName:%s \tMonthly Price: ₹%d\n", v.ID, v.Name, v.Price)
	}
	fmt.Printf("\nEnter ID Associated:")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	id, err := strconv.Atoi(string(char))
	fmt.Printf("You choose :%d\n\n", id)
	if id > 0 {
		id = id - 1
	} else {
		fmt.Println("\nPack ID Cannot be zero")
		return
	}
	pack := packs[id]
	fmt.Printf("\nEnter The Number of Months:")
	monthReader := bufio.NewReader(os.Stdin)
	month, _, err := monthReader.ReadRune()
	if err != nil {
		fmt.Println("Please Enter Valid Value!", err)
		return
	}
	months, _ := strconv.Atoi(string(month))
	userData, amount, discount, fAmount, err := h.svc.SubscribeBasePack(user, pack, months)
	if err != nil {
		fmt.Printf("Unable to Subscribe:%v\n", err)
		return
	}

	fmt.Printf("\nHey %s,\n", userData.Name)
	fmt.Printf("You have successfully Subscribed to: %s\n", packs[id].Name)
	fmt.Printf("Monthly Price: %d\n", packs[id].Price)
	fmt.Printf("Subscription Amount: %d\n", amount)
	fmt.Printf("Number of months: %d\n", months)
	fmt.Printf("Discount Applied: %d\n", discount)
	fmt.Printf("Final Price after Discount: %d\n", fAmount)

	fmt.Printf("Your Account Balance Is: %d\n", userData.Balance)
	fmt.Printf("Email Notification Sent Successfully!\n")
	fmt.Printf("SMS Notification Sent Successfully!\n")

}

//AddChannel **
func (h HandlerStruct) AddChannel(user models.User) {
	channels, err := h.svc.ListChannels()
	for _, v := range channels {
		fmt.Printf("ID:%d \tName:%s \tMonthly Price: ₹%d\n", v.ID, v.Name, v.Price)
	}
	fmt.Printf("\nEnter ID Associated:")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	id, err := strconv.Atoi(string(char))
	fmt.Printf("You choose :%d\n\n", id)
	if id > 0 {
		id = id - 1
	} else {
		fmt.Println("\nPack ID Cannot be zero")
		return
	}
	channel := channels[id]
	fmt.Printf("\nEnter The Number of Months:")
	monthReader := bufio.NewReader(os.Stdin)
	month, _, err := monthReader.ReadRune()
	if err != nil {
		fmt.Println("Please Enter Valid Value!", err)
		return
	}
	months, _ := strconv.Atoi(string(month))
	userData, amount, err := h.svc.AddChannel(user, channel, months)
	if err != nil {
		fmt.Printf("Unable to Subscribe:%v\n", err)
		return
	}

	fmt.Printf("\nHey %s,\n", userData.Name)
	fmt.Printf("You have successfully Subscribed to: %s\n", channels[id].Name)
	fmt.Printf("Monthly Price: %d\n", channels[id].Price)
	fmt.Printf("Subscription Amount: %d\n", amount)
	fmt.Printf("Number of months: %d\n", months)
	fmt.Printf("Your Account Balance Is: %d\n", userData.Balance)
	fmt.Printf("Email Notification Sent Successfully!\n")
	fmt.Printf("SMS Notification Sent Successfully!\n")

}

//SubscribeSpecialService **
func (h HandlerStruct) SubscribeSpecialService(user models.User) {
	services, err := h.svc.ListServices()
	for _, v := range services {
		fmt.Printf("ID:%d \tName:%s \tMonthly Price: ₹%d\n", v.ID, v.Name, v.Price)
	}
	fmt.Printf("\nEnter ID Associated:")
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	id, err := strconv.Atoi(string(char))
	fmt.Printf("You choose :%d\n\n", id)
	if id > 0 {
		id = id - 1
	} else {
		fmt.Println("\nPack ID Cannot be zero")
		return
	}
	channel := services[id]
	fmt.Printf("\nEnter The Number of Months:")
	monthReader := bufio.NewReader(os.Stdin)
	month, _, err := monthReader.ReadRune()
	if err != nil {
		fmt.Println("Please Enter Valid Value!", err)
		return
	}
	months, _ := strconv.Atoi(string(month))
	userData, amount, err := h.svc.SubscribeSpecialService(user, channel, months)
	if err != nil {
		fmt.Printf("Unable to Subscribe:%v\n", err)
		return
	}

	fmt.Printf("\nHey %s,\n", userData.Name)
	fmt.Printf("You have successfully Subscribed to: %s\n", services[id].Name)
	fmt.Printf("Monthly Price: %d\n", services[id].Price)
	fmt.Printf("Subscription Amount: %d\n", amount)
	fmt.Printf("Number of months: %d\n", months)
	fmt.Printf("Your Account Balance Is: %d\n", userData.Balance)
	fmt.Printf("Email Notification Sent Successfully!\n")
	fmt.Printf("SMS Notification Sent Successfully!\n")
}

//ViewSubscription **
func (h HandlerStruct) ViewSubscription(user models.User) {
	_, err := h.svc.ViewSubscription(user)
	if err != nil {
		fmt.Printf("No Subscription:%v\n", err)
		return
	}

}

//UpdateInfo **
func (h HandlerStruct) UpdateInfo(user models.User) {
	fmt.Printf("\nEnter Email:")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("\nEnter Phone:")
	phone, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	userData, err := h.svc.UpdateInfo(user, email, phone)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("\nHey %s,\n", userData.Name)
	fmt.Printf("Email Updated Successfully to:%s\n", userData.Email)
	fmt.Printf("Phone Updated Successfully to:%s\n", userData.Phone)

}

//ListUsers **
func (h HandlerStruct) ListUsers() ([]models.User, error) {
	users, err := h.svc.ListUsers()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range users {
		fmt.Printf("\nID:%d | Name:%s |", v.ID, v.Name)
	}
	return users, nil
}

//SelectUser **
func (h HandlerStruct) SelectUser(id int) (user *models.User) {
	user, _ = h.svc.SelectUser(id)
	return user
}
