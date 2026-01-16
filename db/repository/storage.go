package db

import "gorm.io/gorm"

type Storage struct {
	DB *gorm.DB
	ReviewRepository ReviewRepository
}

func NewStorage(DB *gorm.DB)*Storage{
	return &Storage{
		DB: DB,
		ReviewRepository: NewReviewRepository(DB),
	}
}