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
// @Param        body body pb.CreateReviewReq true "Create Review Request"
// @Success      200 {object} pb.CreateReviewRes
// @Failure      400 {object} string
// @Router       /review/createreview [post]
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

func (h *Handler) GetAllReviews(c *gin.Context) {}

func (h *Handler) GetReviewById(c *gin.Context) {}

func (h *Handler) UpdateReview(c *gin.Context) {}

func (h *Handler) DeleteReview(c *gin.Context) {}