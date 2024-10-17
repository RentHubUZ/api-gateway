package handler

import (
	request "api_gateway/genproto/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRequest godoc
// @Summary      Create a new request
// @Description  Add a new user request
// @Tags         request
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        body body request.CreateRequestRequest true "Create Request Request"
// @Success      200 {object} request.CreateRequestResponse
// @Failure      400 {object} string
// @Router       /request/create [post]
func (h *Handler) CreateRequest(c *gin.Context) {
	var req request.CreateRequestRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.UserId == "" || req.UserId == "string" {
		UserId, err := getUserID(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.UserId = UserId
	}

	res, err := h.RequestService.CreateRequest(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// GetRequestById godoc
// @Summary      Get a request by ID
// @Description  Retrieve a specific request by its ID
// @Tags         request
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id path string true "Request ID"
// @Success      200 {object} request.GetRequestResponse
// @Failure      400 {object} string
// @Router       /request/getbyid/{id} [get]
func (h *Handler) GetRequestById(c *gin.Context) {
	var req request.GetRequestRequest
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	res, err := h.RequestService.GetRequest(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteRequest godoc
// @Summary      Delete a request
// @Description  Remove a request by its ID
// @Tags         request
// @Accept       json
// @Produce      json
// @Security ApiKeyAuth
// @Param        id path string true "Request ID"
// @Success      200 {object} request.Void
// @Failure      400 {object} string
// @Router       /request/delete/{id} [delete]
func (h *Handler) DeleteRequest(c *gin.Context) {
	var req request.Request
	req.Id = c.Param("id")
	if req.Id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	_, err := h.RequestService.DeleteRequest(c, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Request deleted successfully"})
}
