package repositories

import (
	"biloba-cmd/driver"
	"biloba-cmd/model"
	"context"
	"errors"
	"fmt"
)
type KidRepository struct {
	DB *driver.DB
}

func NewKidRepository(database *driver.DB) *KidRepository {
	return &KidRepository{database}
}


func (k *KidRepository) RetrieveKids(UserID string) (userKids *model.Kids, err error) {
	path := fmt.Sprintf("user-kids/%s", UserID)
	ref := k.DB.Client.NewRef(path)

	var data *map[string]model.Kid

	if err := ref.Get(context.Background(), &data); err != nil {
		return userKids, errors.New(fmt.Sprintf("Error reading from database:", err))
	}

	if data == nil {
		return userKids, errors.New(fmt.Sprintf("Cannot find kids"))
	}

	kids := make(model.Kids, 0)

	for id, kid := range *data {
		kid.ID = id
		kids = append(kids, kid)
	}

	return &kids, nil
}
