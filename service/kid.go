package service

import (
	"biloba-cmd/driver"
	"biloba-cmd/model"
	"biloba-cmd/repositories"
)

type KidService struct {
	kidRepository *repositories.KidRepository
}

func NewKidService(db *driver.DB) *KidService {
	return &KidService{repositories.NewKidRepository(db)}
}

func (k *KidService) RetrieveKids(UserID string) (kids *model.Kids, err error) {
	return k.kidRepository.RetrieveKids(UserID)
}
