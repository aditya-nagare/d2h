package cmd

import (
	db "github.com/aditya-nagare/d2h/db"
	handler "github.com/aditya-nagare/d2h/services/pkg/handlers"
	"github.com/aditya-nagare/d2h/services/pkg/menu"
	repository "github.com/aditya-nagare/d2h/services/pkg/repositories"
	service "github.com/aditya-nagare/d2h/services/pkg/service"
)

//Run **
func Run() {
	conn := db.NewDbConnection()
	repository := repository.NewRepositoryImpl(conn)
	service := service.NewServiceImpl(repository)
	handler := handler.NewHandlerImpl(service)
	menu.Menu(handler)
}
