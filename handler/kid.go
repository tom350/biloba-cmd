package handler

import (
	"biloba-cmd/driver"
	"biloba-cmd/model"
	"biloba-cmd/service"
	"fmt"
)

type KidHandler struct {
	kidService *service.KidService
}

type App struct {
	DB *driver.DB
}

func KidPrint(app *App, UserID string) model.Kids{
	kidHandler := NewKidHandler(app)
	kids := *kidHandler.getKid(UserID)
	return kids
}

func NewKidHandler(app *App) *KidHandler {
	return &KidHandler{service.NewKidService(app.DB)}
}

func (k *KidHandler) getKid(UserID string) (kids *model.Kids){
	kids, err := k.kidService.RetrieveKids(UserID)
	if err != nil {
		fmt.Println(err)
	}

	return kids
}
