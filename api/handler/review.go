package handler

import (
	pb "api_gateway/genproto/reviews"
	"api_gateway/api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateReview(c *gin.Context) {
	claims, err := token.ExtractClaims(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "cannot get user id"})
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
