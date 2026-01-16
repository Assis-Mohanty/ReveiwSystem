package controllers

import (
	"reviewservice/models"
	"reviewservice/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ReviewController struct {
	ReviewService services.ReviewService
}

func NewReviewController(_reviewService services.ReviewService) *ReviewController{
	return &ReviewController{
		ReviewService: _reviewService,
	}
}

func (rc *ReviewController) GetById(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid id"})
	}

	review, err := rc.ReviewService.GetById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(review)
}

func(rc *ReviewController) CreateReview(c *fiber.Ctx) error{
	var review models.Review
	if err:=c.BodyParser(&review);err!=nil{
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid idqqq"})	 
	}
	bookingId ,err :=rc.ReviewService.CreateReview(&review)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"invalid id"})
	}
	return c.JSON(bookingId)
}

func (rc *ReviewController) GetAllReviews(c *fiber.Ctx)error{
	// var review models.Review
	reviews ,err :=rc.ReviewService.GetAllReviews()
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"invalid id"})
	}
	return c.JSON(reviews)	
}

func (rc *ReviewController) DeleteReview(c *fiber.Ctx)error{
	idParams:=c.Params("id")
	idKey,err:=strconv.ParseInt(idParams,10,64)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error":"invalid id"})

	}
	deletedErr:=rc.ReviewService.DeleteReview(idKey)
	if deletedErr != nil {
		return deletedErr
	}
	return nil
}

func (rc *ReviewController) UpdateReview(c *fiber.Ctx) error{
	var updatedBody *models.Review
	if err:=c.BodyParser(&updatedBody);err!=nil{
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid idqqq"})	 
	}
	reviewId:=c.Params("id")
	review,err:=rc.ReviewService.UpdateReview(updatedBody,reviewId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid id"})
	}
	return c.JSON(review)	
}