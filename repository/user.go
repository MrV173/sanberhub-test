package repository

import (
	"errors"
	"sanberhub-test/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	Transaction(transaction models.Transaction) (models.Transaction, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Transactions").Find(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Transactions").First(&user, ID).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	var existingUser models.User
	if err := r.db.Where("nik = ? ", user.Nik).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return existingUser, err
		}
	} else {
		err = errors.New("NIK sudah digunakan")
		return existingUser, err
	}
	if err := r.db.Where("no_hp = ? ", user.NoHp).First(&existingUser).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return existingUser, err
		}
	} else {
		err = errors.New("No Handphone sudah digunakan")
		return existingUser, err
	}
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) Transaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}
