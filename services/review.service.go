package services

import (
	db "reviewservice/db/repository"
	"reviewservice/models"
)

type ReviewService interface {
	GetById(int64) (*models.Review,error)
	CreateReview(review *models.Review)(int64,error)
	GetAllReviews()([]models.Review,error)
	DeleteReview(reviewId int64)error
	UpdateReview(review *models.Review,reviewId string)(*models.Review,error)
}

type ReviewServiceImpl struct{
	ReviewRepository db.ReviewRepository
}

func NewReviewService(_reviewRepository db.ReviewRepository) ReviewService{
	return &ReviewServiceImpl{
		ReviewRepository: _reviewRepository,
	}
}

func (r *ReviewServiceImpl) GetById(id int64) (*models.Review,error){
	return r.ReviewRepository.GetById(id)
}

func (r *ReviewServiceImpl) CreateReview(review *models.Review) (int64, error) {
	review.IsSynced = false
	return r.ReviewRepository.CreateReview(review)
}
func (r *ReviewServiceImpl) GetAllReviews()([]models.Review,error){
	review,err:=r.ReviewRepository.GetAllReviews()
	if err != nil {
		return nil,err
	}
	return review,nil
}
func (r *ReviewServiceImpl) DeleteReview(reviewId int64)error{
	err:=r.ReviewRepository.DeleteReview(reviewId)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewServiceImpl) UpdateReview(review *models.Review,reviewId string)(*models.Review,error){
	result,err:=r.ReviewRepository.UpdateReview(review,reviewId)
	if err != nil {
		return nil,err
	}
	return result,err
}