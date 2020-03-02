package repositories

import (
	"fmt"

	"github.com/aditya-nagare/d2h/services/pkg/models"

	"github.com/jinzhu/gorm"
)

//RepositoryInterface **
type RepositoryInterface interface {
	Recharge(*models.User) (*models.User, error)
	ViewPacks() ([]models.Channel, []models.Service, error)
	SubscribeBasePack(models.User, models.Subscription) (*models.User, error)
	AddChannel(user models.User, subs models.Subscription) (*models.User, error)
	SubscribeSpecialService(models.User, models.Subscription) (*models.User, error)
	ViewSubscription(models.User) ([]models.SubscriptionDetails, error)
	UpdateInfo(models.User) (*models.User, error)
	ListUsers() ([]models.User, error)
	SelectUser(id int) (*models.User, error)
	ListPacks() ([]models.Pack, error)
	ListServices() ([]models.Service, error)
	ListChannels() ([]models.Channel, error)
}

//RepositoryStruct **
type RepositoryStruct struct {
	dbConn *gorm.DB
}

//NewRepositoryImpl **
func NewRepositoryImpl(dbConn *gorm.DB) RepositoryInterface {
	return &RepositoryStruct{dbConn: dbConn}
}

//Recharge **
func (r RepositoryStruct) Recharge(user *models.User) (*models.User, error) {
	db := r.dbConn
	err := db.Table("users").Where("id=?", user.ID).Update(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

//ViewPacks **
func (r RepositoryStruct) ViewPacks() (channels []models.Channel, services []models.Service, err error) {
	db := r.dbConn
	packs := []models.Pack{}
	packComp := []models.PackComposition{}
	err = db.Find(&packs).Error
	if err != nil {
		return nil, nil, err
	}
	fmt.Println("Available Packs for Subscription:")
	for _, v := range packs {
		fmt.Println("\n", v.Name, ": Rs.", v.Price)

		fmt.Println("Channels:")
		err = db.Raw("SELECT channels.name FROM pack_compositions JOIN channels ON channels.id = pack_compositions.channel_id AND pack_compositions.pack_id = ?", v.ID).Scan(&packComp).Error
		if err != nil {
			fmt.Printf("No Channles Found In The Database!\n")
			return nil, nil, err
		}
		for _, v := range packComp {
			fmt.Println(v.Name)
		}
	}

	err = db.Find(&channels).Error
	if err != nil {
		fmt.Printf("No Channles Found In The Database!\n")
		return nil, nil, err
	}

	err = db.Find(&services).Error
	if err != nil {
		fmt.Printf("No Services Found In The Database!\n")
		return nil, nil, err
	}

	return channels, services, nil
}

//SubscribeBasePack **
func (r RepositoryStruct) SubscribeBasePack(user models.User, subs models.Subscription) (*models.User, error) {
	db := r.dbConn
	err := db.Create(&subs).Error
	if err != nil {
		return nil, err
	}
	err = db.Table("users").Where("id=?", user.ID).Update(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//AddChannel **
func (r RepositoryStruct) AddChannel(user models.User, subs models.Subscription) (*models.User, error) {
	db := r.dbConn
	err := db.Create(&subs).Error
	if err != nil {
		return nil, err
	}
	err = db.Table("users").Where("id=?", user.ID).Update(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//SubscribeSpecialService **
func (r RepositoryStruct) SubscribeSpecialService(user models.User, subs models.Subscription) (*models.User, error) {
	db := r.dbConn
	err := db.Create(&subs).Error
	if err != nil {
		return nil, err
	}
	err = db.Table("users").Where("id=?", user.ID).Update(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//ViewSubscription **
func (r RepositoryStruct) ViewSubscription(user models.User) (sub []models.SubscriptionDetails, err error) {
	db := r.dbConn
	err = db.Table("subscriptions").Select("packs.name AS pack_name").Joins("JOIN packs ON subscriptions.pack_id=packs.id").Where("user_id=?", user.ID).Scan(&sub).Error
	if err != nil {
		fmt.Printf("No Services Found In The Database!\n")
	}
	fmt.Printf("\nPacks:")
	for _, v := range sub {
		fmt.Printf("\n%s\n", v.PackName)
	}

	err = db.Table("subscriptions").Select("channels.name AS channel_name").Joins("JOIN channels ON subscriptions.channel_id= channels.id").Where("user_id=?", user.ID).Scan(&sub).Error
	if err != nil {
		fmt.Printf("No Services Found In The Database!\n")
	}
	fmt.Printf("\nAdd-On Channels:")
	for _, v := range sub {
		fmt.Printf("\n%s\n", v.ChannelName)
	}

	err = db.Table("subscriptions").Select("services.name AS service_name").Joins("JOIN services ON subscriptions.service_id=services.id").Where("user_id=?", user.ID).Scan(&sub).Error
	if err != nil {
		fmt.Printf("No Services Found In The Database!\n")
	}
	for _, v := range sub {
		fmt.Printf("\nServices:\n%s\n", v.ServiceName)
	}
	return sub, nil

}

//UpdateInfo **
func (r RepositoryStruct) UpdateInfo(user models.User) (*models.User, error) {
	db := r.dbConn
	err := db.Table("users").Where("id=?", user.ID).Update(&user).Error
	if err != nil {
		fmt.Printf("Unable to Recharge:%v\n", err)
		return nil, err
	}
	return &user, nil
}

//ListUsers **
func (r RepositoryStruct) ListUsers() ([]models.User, error) {
	users := []models.User{}
	db := r.dbConn
	err := db.Find(&users).Error
	if err != nil {
		fmt.Printf("No Users Found In The Database!\n")
		return nil, err
	}
	return users, nil
}

//SelectUser **
func (r RepositoryStruct) SelectUser(id int) (*models.User, error) {
	user := models.User{}
	db := r.dbConn

	err := db.Where("id=?", id).First(&user).Error
	if err != nil {
		fmt.Printf("No User Found with Username\n")
		return nil, err
	}
	return &user, nil
}

//ListPacks **
func (r RepositoryStruct) ListPacks() (packs []models.Pack, err error) {
	db := r.dbConn
	err = db.Find(&packs).Error
	if err != nil {
		return nil, err
	}
	return packs, nil
}

//ListServices **
func (r RepositoryStruct) ListServices() (services []models.Service, err error) {
	db := r.dbConn
	err = db.Find(&services).Error
	if err != nil {
		return nil, err
	}
	return services, nil
}

//ListChannels **
func (r RepositoryStruct) ListChannels() (channels []models.Channel, err error) {
	db := r.dbConn
	err = db.Find(&channels).Error
	if err != nil {
		return nil, err
	}
	return channels, nil
}
