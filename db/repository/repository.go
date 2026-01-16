package db

import (
	"reviewservice/models"

	"gorm.io/gorm"
)

type ReviewRepository interface {
	GetById(id int64) (*models.Review,error)
	CreateReview(review *models.Review)(int64,error)
	GetAllReviews()([]models.Review,error)
	DeleteReview(int64)error
	UpdateReview(*models.Review,string)(*models.Review,error)
}

type ReviewRepositoryImpl struct {
	DB *gorm.DB
}

func NewReviewRepository(_db *gorm.DB) ReviewRepository{
	return &ReviewRepositoryImpl{
		DB:_db,
	}
}

func (rw *ReviewRepositoryImpl) GetById(id int64) (*models.Review,error){
	var review models.Review
	result := rw.DB.Where("id = ?", id).First(&review)
	if result.Error != nil {
		return nil,result.Error
	}
	return &review,nil
}


func (rw *ReviewRepositoryImpl) CreateReview(reviewq *models.Review) (int64 ,error) {
	// var review models.CreateReviewDTO
	result:=rw.DB.Create(reviewq)
	if result.Error != nil {
		return -1,result.Error
	}
	return reviewq.ID,nil
}

func (rw *ReviewRepositoryImpl)DeleteReview(reviewId int64) error{
	resultErr:=rw.DB.Delete(&models.Review{},reviewId)
	if resultErr != nil {
		return resultErr.Error
	}
	return nil
}

func (rw *ReviewRepositoryImpl) GetAllReviews() ([]models.Review,error){
	var reviews []models.Review
	result:=rw.DB.Find(&reviews)
	if result.Error!= nil {
		return nil,result.Error
	}
	return reviews,nil
}

func (rw *ReviewRepositoryImpl) UpdateReview(updated *models.Review,reviewId string)(*models.Review ,error){
	var review *models.Review
	result:=rw.DB.First(&review,"id=?",reviewId).Model(&review).Updates(updated)
	if result.Error!= nil {
		return nil,result.Error
	}
	return review,nil
}

func (rw *ReviewRepositoryImpl) UpdateReviewq(updated *models.Review, reviewId string) (*models.Review, error) {
	var review models.Review

	// Find existing record
	if err := rw.DB.First(&review, "id = ?", reviewId).Error; err != nil {
		return nil, err
	}

	// Update fields
	if err := rw.DB.Model(&review).Updates(updated).Error; err != nil {
		return nil, err
	}                  

	return &review, nil
}
