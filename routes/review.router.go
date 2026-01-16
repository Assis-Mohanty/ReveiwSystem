package routes

import (
	"reviewservice/controllers"
	"github.com/gofiber/fiber/v2"
)

type ReviewRouter struct {
	ReviewController *controllers.ReviewController
}

func NewReviewRouter(reviewController *controllers.ReviewController) *ReviewRouter{
	return &ReviewRouter{
		ReviewController: reviewController,
	}
}

func (rr *ReviewRouter) Register(r fiber.Router) {
	r.Get("/:id",rr.ReviewController.GetById)
	r.Get("/",rr.ReviewController.GetAllReviews)
	r.Post("/",rr.ReviewController.CreateReview)
	r.Delete("/:id",rr.ReviewController.DeleteReview)
	r.Patch("/:id",rr.ReviewController.UpdateReview)
}