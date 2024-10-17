package handler

import (
	pb "api_gateway/genproto/reviews"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReview godoc
// @Summary      Create a new review
// @Description  Create a new review with specified details
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        body body reviews.CreateReviewReq true "Create Review Request"
// @Success      200 {object} reviews.CreateReviewRes
// @Failure      400 {object} string
// @Router       /review/create [post]
func (h *Handler) CreateReview(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req pb.CreateReviewReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.UserId = userID
	res, err := h.ReviewService.CreateReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)

}

// GetAllReviews godoc
// @Summary      Get all reviews
// @Description  Get a paginated list of all reviews
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        limit query int true "Limit"
// @Param        page query int true "Page"
// @Success      200 {object} reviews.GetAllReviewRes
// @Failure      400 {object} string
// @Router       /review/getallreview [get]
func (h *Handler) GetAllReviews(c *gin.Context) {
	var req pb.GetAllReviewReq
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ReviewService.GetAllReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetReviewById godoc
// @Summary      Get a review by ID
// @Description  Get the details of a specific review by its ID
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "Review ID"
// @Success      200 {object} reviews.GetByIdReviewRes
// @Failure      400 {object} string
// @Router       /review/getbyid/{id} [get]
func (h *Handler) GetReviewById(c *gin.Context) {
	var req pb.GetByIdReviewReq
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	res, err := h.ReviewService.GetByIdReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateReview godoc
// @Summary      Update an existing review
// @Description  Update the details of a review by its ID
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "Review ID"
// @Param        body body reviews.CreateReviewReq true "Update Review Request"
// @Success      200 {object} reviews.CreateReviewRes
// @Failure      400 {object} string
// @Router       /review/update/{id} [put]
func (h *Handler) UpdateReview(c *gin.Context) {
	var req pb.CreateReviewReq
	req.UserId = c.Param("id") // Assuming the user ID is passed as a path param, adjust as necessary
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.ReviewService.CreateReview(c, &req) // Assuming CreateReview is used for both creation and updates
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteReview godoc
// @Summary      Delete a review
// @Description  Delete a review by its ID
// @Tags         reviews
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Param        id path string true "Review ID"
// @Success      200 {object} reviews.DeleteReviewRes
// @Failure      400 {object} string
// @Router       /review/delete/{id} [delete]
func (h *Handler) DeleteReview(c *gin.Context) {
	var req pb.DeleteReviewReq
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	res, err := h.ReviewService.DeleteReview(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
