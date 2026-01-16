package models

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
    ID        int64                 `gorm:"column:id;primaryKey"`
    UserID    int64                 `gorm:"column:user_id"`
    HotelID   int64                 `gorm:"column:hotel_id"`
    BookingID int64                 `gorm:"column:booking_id"`
    Comment   string                `gorm:"column:comment"`
    Rating    float64               `gorm:"column:rating"`
    CreatedAt time.Time             `gorm:"column:created_at"`
    UpdatedAt time.Time             `gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index"`
    IsSynced  bool                  `gorm:"column:is_synced"`
}


type CreateReviewDTO struct{
	UserId  int64
	HotelId   int64
	BookingId int64
	Comment string
	Rating   float64
}

type UpdateReviewDTO struct{
	UserId  int64
	HotelId   int64
	BookingId int64
	Comment string
	Rating   float64
}