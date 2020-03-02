package services

import (
	"time"

	"github.com/aditya-nagare/d2h/services/pkg/models"
	repository "github.com/aditya-nagare/d2h/services/pkg/repositories"
)

//ServiceInterface **
type ServiceInterface interface {
	Recharge(int, *models.User) (*models.User, error)
	ViewPacks() ([]models.Channel, []models.Service, error)
	SubscribeBasePack(models.User, models.Pack, int) (*models.User, uint, uint, uint, error)
	AddChannel(models.User, models.Channel, int) (*models.User, uint, error)
	SubscribeSpecialService(models.User, models.Service, int) (*models.User, uint, error)
	ViewSubscription(models.User) ([]models.SubscriptionDetails, error)
	UpdateInfo(models.User, string, string) (*models.User, error)
	ListUsers() ([]models.User, error)
	SelectUser(id int) (*models.User, error)
	ListPacks() ([]models.Pack, error)
	ListServices() ([]models.Service, error)
	ListChannels() ([]models.Channel, error)
}

//ServiceStruct **
type ServiceStruct struct {
	repo repository.RepositoryInterface
}

//NewServiceImpl **
func NewServiceImpl(repo repository.RepositoryInterface) ServiceInterface {
	return &ServiceStruct{repo: repo}
}

//Recharge **
func (s ServiceStruct) Recharge(amount int, user *models.User) (*models.User, error) {
	user.Balance = user.Balance + uint(amount)
	user.UpdatedAt = time.Now()
	user, err := s.repo.Recharge(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//ViewPacks **
func (s ServiceStruct) ViewPacks() (channels []models.Channel, services []models.Service, err error) {
	channels, services, err = s.repo.ViewPacks()
	if err != nil {
		return nil, nil, err
	}
	return channels, services, nil
}

//SubscribeBasePack **
func (s ServiceStruct) SubscribeBasePack(user models.User, pack models.Pack, months int) (*models.User, uint, uint, uint, error) {

	subs := models.Subscription{}
	var fAmount, discount uint
	amount := pack.Price * uint(months)

	if months >= 3 {
		discount = amount / 10
		fAmount = amount - discount
	} else {
		fAmount = amount
	}

	subs.UserID = user.ID
	subs.PackID = pack.ID
	subs.Duration = uint(months)

	user.Balance = user.Balance - fAmount
	userData, err := s.repo.SubscribeBasePack(user, subs)
	if err != nil {
		return userData, amount, discount, fAmount, err
	}
	return userData, amount, discount, fAmount, nil
}

//AddChannel **
func (s ServiceStruct) AddChannel(user models.User, channel models.Channel, months int) (*models.User, uint, error) {
	subs := models.Subscription{}
	amount := channel.Price * uint(months)

	subs.UserID = user.ID
	subs.ChannelID = channel.ID
	subs.Duration = uint(months)

	user.Balance = user.Balance - amount
	userData, err := s.repo.AddChannel(user, subs)
	if err != nil {
		return userData, amount, err
	}
	return userData, amount, nil

}

//SubscribeSpecialService **
func (s ServiceStruct) SubscribeSpecialService(user models.User, service models.Service, months int) (*models.User, uint, error) {
	subs := models.Subscription{}
	amount := service.Price * uint(months)

	subs.UserID = user.ID
	subs.ChannelID = service.ID
	subs.Duration = uint(months)

	user.Balance = user.Balance - amount
	userData, err := s.repo.AddChannel(user, subs)
	if err != nil {
		return userData, amount, err
	}
	return userData, amount, nil
}

//ViewSubscription **
func (s ServiceStruct) ViewSubscription(user models.User) (subscriptions []models.SubscriptionDetails, err error) {
	subscriptions, err = s.repo.ViewSubscription(user)
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

//UpdateInfo **
func (s ServiceStruct) UpdateInfo(user models.User, email string, phone string) (*models.User, error) {
	user.Email = email
	user.Phone = phone
	user.UpdatedAt = time.Now()
	userData, err := s.repo.UpdateInfo(user)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

//ListUsers **
func (s ServiceStruct) ListUsers() ([]models.User, error) {
	users, err := s.repo.ListUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

//SelectUser **
func (s ServiceStruct) SelectUser(id int) (*models.User, error) {
	user, err := s.repo.SelectUser(id)
	return user, err
}

//ListPacks **
func (s ServiceStruct) ListPacks() (packs []models.Pack, err error) {
	packs, err = s.repo.ListPacks()
	if err != nil {
		return nil, err
	}
	return packs, nil

}

//ListServices **
func (s ServiceStruct) ListServices() (services []models.Service, err error) {
	services, err = s.repo.ListServices()
	if err != nil {
		return nil, err
	}
	return services, nil
}

//ListChannels **
func (s ServiceStruct) ListChannels() (channels []models.Channel, err error) {
	channels, err = s.repo.ListChannels()
	if err != nil {
		return nil, err
	}
	return channels, nil
}
